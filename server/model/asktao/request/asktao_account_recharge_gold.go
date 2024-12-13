package request

type RechargeGoldParam struct {
	Account    string `json:"account"`
	GoldCoin   int    `json:"goldCoin"`
	SilverCoin int    `json:"silverCoin"`
}
