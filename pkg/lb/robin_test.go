//__author__ = "YaoYao"
//Date: 2020/3/18
package lb

import (
	"fmt"
	"github.com/yaoliu/Gaia/pkg/backend"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	rr := newWeightRandom()
	var backends = []backend.Backend{
		backend.NewHTTPBackend("192.168.1.1", 1),
		backend.NewHTTPBackend("192.168.1.2", 2),
		backend.NewHTTPBackend("192.168.1.3", 7),
	}
	for i := 0; i < 10; i++ {
		target := rr.DoSelect(backends)
		fmt.Println(target)
	}
}
