package asktao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (e *AccountRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	accountRouter := Router.Group("account").Use(middleware.OperationRecord())
	{
		accountRouter.POST("addGold", accountApi.RechargeGold) // 充值
	}
}
