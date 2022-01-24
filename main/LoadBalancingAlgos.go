package main

type loadBalancingAlgo interface {
	getNextBackend() *Backend
}

type roundRobin struct {
	index    int
	backends []Backend
}

func (rr *roundRobin) roundRobinInit(backends *[]Backend) {
	rr.index = 0
	rr.backends = *backends
}

func (rr *roundRobin) getNextBackend() *Backend {
	rr.index = (rr.index + 1) % len(rr.backends)
	return &rr.backends[rr.index]
}
