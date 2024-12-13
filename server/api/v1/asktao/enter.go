package asktao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	AccountApi
}

var (
	asktaoAccountService = service.ServiceGroupApp.AsktaoServiceGroup.AsktaoAccountService
)
