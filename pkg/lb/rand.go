//__author__ = "YaoYao"
//Date: 2020/3/19
package lb

import (
	"github.com/yaoliu/Gaia/pkg/backend"
	"math/rand"
)

type Random struct { //随机
	size *uint64
}

func newRandom() LoadBalance {
	var size uint64
	size = 0
	return &Random{size: &size}
}

func (lb Random) DoSelect(backends []backend.Backend) backend.Backend {
	n := len(backends)
	if n <= 0 {
		return nil
	}
	target := backends[rand.Intn(n)]
	return target
}
