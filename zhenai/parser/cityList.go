package parser

import (
	"regexp"

	"imooc.com/learngo/crawler/engine"
)

const CityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]+>([^<]+)</a>`

// ParserCityList ParserCityList
func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(CityListRegex)
	// matches := compile.FindAll(contents, -1)
	matches := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		// result.Items = append(result.Items, "City "+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
