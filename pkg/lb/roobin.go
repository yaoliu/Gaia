//__author__ = "YaoYao"
//Date: 2020/3/18
package lb

import (
	"github.com/yaoliu/Gaia/pkg/backend"
	"sync/atomic"
)

type RoundRobin struct { //轮询
	size *uint64
}

func newRoundRobin() LoadBalance {
	var size uint64
	size = 0
	return &RoundRobin{size: &size}
}

func (lb RoundRobin) DoSelect(backends []backend.Backend) backend.Backend {
	n := len(backends)
	if n <= 0 {
		return nil
	}
	target := backends[int(atomic.AddUint64(lb.size, 1)%10)]
	return target
}
