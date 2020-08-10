package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var rateLimit = time.Tick(10 * time.Microsecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimit
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	request, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4214.3 Safari/537.36")
	cookie1 := "sid=db53037f-5ef5-4313-9d56-68587aad92ec; FSSBBIl1UgzbN7NO=5sNICONn7kBFgu_LmDI4Nxad_RqUvst66ytJ3INz8cWUwXghrcZr9cT_8mc_kbBFVRVnDyDkmDmeuV1lde8qnJA; ec=nUwjtjqJ-1595343827633-8b63783abd068-1493495454; _efmdata=puSUckPNmffonDypEMBcywJ7HouKl16rV6wcuGo4gRb6ETf9dSC8KCNle7wXvcvUH4Om05BNacF7iPXS%2Fgh1s0sl0%2FqcCN%2BIDg57hGbjo70%3D; _exid=kklrIj0LbYubYc5WcN0qHETnot5lfQyTxy5NLK0tnfTMciucqgJ7EXBho1%2Fl1yU6vSHatwrp5dcq0AW4%2FEqEGA%3D%3D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1596277505,1596289401,1596812201,1596948447; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1596948449; FSSBBIl1UgzbN7NP=5Uned5TZBIdLqqqmu03BXRaUVgIMwvRHpRzJYU14vk8zDZESqMl8_TzQ04_RX4GlxeggYAJDepiD2.4Dq0o7Ob0ZvI3aHmASXlmQ3dVPNtmFoT8PK3ZNmgDcjusTj3Svz8fPWuT2W64qgWDRXJDJqYSd0a2grlgehKkWrkFz3ELsP8ly7l5s.378uoFVlQftKz0cm2P.fNTVDY3VzyNVuMTCL6KGV0aCgsgb7yIzH_scrlo2OtjLlj58_Bbxk5frMZ.0unPcUfAfAuDpo3qilgB"
	request.Header.Add("cookie", cookie1)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	// GBK 转换为utf8
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	// body, err := ioutil.ReadAll(utf8Reader)
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// determineEncoding 检测html 编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error : %s", err)
		return unicode.UTF8
	}
	// 自动检测编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
