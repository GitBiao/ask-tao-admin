package asktao

type Account struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	GoldCoin   string `json:"goldCoin"`
	SilverCoin string `json:"silverCoin"`
	Privilege  string `json:"privilege"`
	Checksum   string `json:"checksum"`
}

func (Account) TableName() string {
	return "account"
}
