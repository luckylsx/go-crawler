package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const ResultSize = 311
	urls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	cities := []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(result.Requests) != ResultSize {
		t.Errorf("result should have %d request but had %d", ResultSize, len(result.Requests))
	}

	if len(result.Items) != ResultSize {
		t.Errorf("result should have %d request but had %d", ResultSize, len(result.Items))
	}

	for i := 0; i < len(urls); i++ {
		if urls[i] != result.Requests[i].Url {
			t.Errorf("the url is not matched,  %d url should be %s, but got %s", i, urls[i], result.Requests[i].Url)
		}
	}

	for i := 0; i < len(cities); i++ {
		if cities[i] != result.Items[i] {
			t.Errorf("the city name is not matched, %d city should be %s, but got %s", i, cities[i], result.Items[i])
		}
	}
}
