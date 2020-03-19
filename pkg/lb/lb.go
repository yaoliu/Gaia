//__author__ = "YaoYao"
//Date: 2019/12/8
package lb

import (
	"github.com/yaoliu/Gaia/pkg/backend"
)

type LoadBalance interface {
	DoSelect(backends []backend.Backend) backend.Backend
}
