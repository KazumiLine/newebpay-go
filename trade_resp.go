package newebpay

type TradeResponse struct {
	Status       string `json:"Status"`
	MerchantID   string `json:"MerchantID"`
	TradeInfoEnc string `json:"TradeInfo"`
	TradeInfo    struct {
		Status  string
		Message string
	} `json:"-"`
	TradeSha    string `json:"TradeSha"`
	Version     string `json:"Version"`
	EncryptType int    `json:"EncryptType"`
}
