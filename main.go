package main

import (
	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/scheduler"
	"imooc.com/learngo/crawler/zhenai/parser"
)

func main() {
	var e = engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		// Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}
	/*e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})*/

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})
}
