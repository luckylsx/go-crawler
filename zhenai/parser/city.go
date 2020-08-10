package parser

import (
	"regexp"

	"imooc.com/learngo/crawler/engine"
)

var (
	profileRegex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlRegx  = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte) engine.ParseResult {
	// compile := regexp.MustCompile(cityRegex)
	// matches := compile.FindAll(contents, -1)
	matches := profileRegex.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		name := match[2]
		// result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(name))
			},
		})
	}
	matches = cityUrlRegx.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
