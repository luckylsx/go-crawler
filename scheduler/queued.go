package scheduler

import "imooc.com/learngo/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workChan <- w
}

func (s *QueueScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}
