package asktao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/asktao"
	rechargeGoldParam "github.com/flipped-aurora/gin-vue-admin/server/model/asktao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type AccountApi struct{}

func (b *AccountApi) RechargeGold(c *gin.Context) {
	var req rechargeGoldParam.RechargeGoldParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var account asktao.Account
	if err = global.GVA_DBList["asktao_dl_adb_all"].Where("id = ?", req.Account).First(&account).Error; err != nil {
		return
	}

	response.OkWithDetailed(account, "创建成功", c)
}
