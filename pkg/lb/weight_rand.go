//__author__ = "YaoYao"
//Date: 2020/3/19
package lb

import (
	"github.com/yaoliu/Gaia/pkg/backend"
	"math/rand"
)

type WeightRandom struct { //加权随机
	size *uint64
}

func newWeightRandom() LoadBalance {
	var size uint64
	size = 0
	return &WeightRandom{size: &size}
}

func (lb WeightRandom) DoSelect(backends []backend.Backend) backend.Backend {
	n := len(backends)
	if n <= 0 {
		return nil
	}
	bs := make([]backend.Backend, 0)
	var weight = 0
	for _, b := range backends {
		for i := 0; i < b.Weight(); i++ {
			bs = append(bs, b)
		}
		weight += b.Weight()
	}
	target := bs[rand.Intn(weight)]
	return target
}
