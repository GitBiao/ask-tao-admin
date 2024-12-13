package asktao

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AccountRouter
}

var (
	accountApi = api.ApiGroupApp.AsktaoApiGroup.AccountApi
)
