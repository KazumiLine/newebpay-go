package newebpay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TradeInfo struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Result  struct {
		MerchantID      string      `json:"MerchantID"`
		Amt             int         `json:"Amt"`
		TradeNo         string      `json:"TradeNo"`
		MerchantOrderNo string      `json:"MerchantOrderNo"`
		TradeStatus     TradeStatus `json:"TradeStatus"`
		PaymentType     string      `json:"PaymentType"`
		CreateTime      string      `json:"CreateTime"`
		PayTime         string      `json:"PayTime"`
		CheckCode       string      `json:"CheckCode"`
		FundTime        string      `json:"FundTime"`
		ShopMerchantID  string      `json:"ShopMerchantID"`
	} `json:"Result"`
}

func (store *Store) QueryTradeInfo(orderNo string, amt int) (r *TradeInfo, err error) {
	var apiUrl = TradeInfoUrl
	if store.SimulationMode {
		apiUrl = TestTradeInfoUrl
	}
	resp, err := http.PostForm(apiUrl, url.Values{
		"MerchantID":      {store.MerchantID},
		"Version":         {"1.3"},
		"RespondType":     {"JSON"},
		"TimeStamp":       {fmt.Sprintf("%d", time.Now().Unix())},
		"MerchantOrderNo": {orderNo},
		"Amt":             {fmt.Sprintf("%d", amt)},
		"CheckValue": {GenerateCheckValue(url.Values{
			"Amt":             {fmt.Sprintf("%d", amt)},
			"MerchantID":      {store.MerchantID},
			"MerchantOrderNo": {orderNo},
		}.Encode(), store.HashKey, store.HashIV)},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bData, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
