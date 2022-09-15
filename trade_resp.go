package newebpay

import (
	"encoding/json"
	"net/http"
)

type TradeResponse struct {
	Status       string
	MerchantID   string
	TradeInfoEnc string
	TradeInfo    struct {
		Status  string `json:"Status"`
		Message string `json:"Message"`
		Result  struct {
			MerchantID      string `json:"MerchantID"`
			Amt             int    `json:"Amt"`
			TradeNo         string `json:"TradeNo"`
			MerchantOrderNo string `json:"MerchantOrderNo"`
			RespondType     string `json:"RespondType"`
			IP              string `json:"IP"`
			EscrowBank      string `json:"EscrowBank"`
			PaymentType     string `json:"PaymentType"`
		} `json:"Result"`
	}
	TradeSha    string
	Version     string
	EncryptType int
}

func (store *Store) ParseTradeResponse(r *http.Request) (*TradeResponse, error) {
	resp := &TradeResponse{
		Status:       r.FormValue("Status"),
		MerchantID:   r.FormValue("MerchantID"),
		TradeInfoEnc: r.FormValue("TradeInfo"),
		TradeSha:     r.FormValue("TradeSha"),
		Version:      r.FormValue("Version"),
	}
	info, err := KeyDecrypt(resp.TradeInfoEnc, store.HashKey, store.HashIV)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(info), &resp.TradeInfo)
	return resp, err
}
