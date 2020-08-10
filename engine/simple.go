package engine

import (
	"log"

	"imooc.com/learngo/crawler/fetcher"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(Seeds ...Request) {
	var requests []Request
	for _, r := range Seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item is : %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url is : %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher : fetcher error , url is : %v, error is :%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), err
}
