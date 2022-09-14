package newebpay

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type TradeRequest struct {
	url.Values
	*Store
	useATM bool
}

func (store *Store) NewTradeRequest(tradeNo string, amount int, itemDesc string) *TradeRequest {
	trade := &TradeRequest{url.Values{}, store, false}
	trade.Set("MerchantID", store.MerchantID)
	trade.Set("RespondType", "JSON")
	trade.Set("TimeStamp", fmt.Sprintf("%d", time.Now().Unix()))
	trade.Set("Version", "2.0")
	trade.Set("MerchantOrderNo", tradeNo)
	trade.Set("Amt", fmt.Sprintf("%d", amount))
	trade.Set("ItemDesc", itemDesc)
	return trade
}

func (trade *TradeRequest) SetLangType(lang LangType) *TradeRequest {
	trade.Set("LangType", string(lang))
	return trade
}

func (trade *TradeRequest) SetTradeLimit(limit int) *TradeRequest {
	trade.Set("TradeLimit", fmt.Sprintf("%d", limit))
	return trade
}
func (trade *TradeRequest) SetExpireDate(date string) *TradeRequest {
	trade.Set("ExpireDate", date)
	return trade
}
func (trade *TradeRequest) SetReturnURL(uri string) *TradeRequest {
	trade.Set("ReturnURL", uri)
	return trade
}
func (trade *TradeRequest) SetNotifyURL(uri string) *TradeRequest {
	trade.Set("NotifyURL", uri)
	return trade
}
func (trade *TradeRequest) SetCustomerURL(uri string) *TradeRequest {
	trade.Set("CustomerURL", uri)
	return trade
}
func (trade *TradeRequest) SetClientBackURL(uri string) *TradeRequest {
	trade.Set("ClientBackURL", uri)
	return trade
}
func (trade *TradeRequest) SetEmail(email string) *TradeRequest {
	trade.Set("Email", email)
	return trade
}
func (trade *TradeRequest) DisableEmailModify() *TradeRequest {
	trade.Set("EmailModify", "0")
	return trade
}
func (trade *TradeRequest) NeedLogin() *TradeRequest {
	trade.Set("LoginType", "1")
	return trade
}
func (trade *TradeRequest) SetOrderComment(comment string) *TradeRequest {
	trade.Set("OrderComment", comment)
	return trade
}
func (trade *TradeRequest) UseCreditCard() *TradeRequest {
	trade.Set("CREDIT", "1")
	return trade
}
func (trade *TradeRequest) UseGooglePay() *TradeRequest {
	trade.Set("ANDROIDPAY", "1")
	return trade
}
func (trade *TradeRequest) UseSamsungPay() *TradeRequest {
	trade.Set("SAMSUNGPAY", "1")
	return trade
}
func (trade *TradeRequest) UseLinePay(imageUrl string) *TradeRequest {
	trade.Set("LINEPAY", "1")
	if len(imageUrl) > 0 {
		trade.Set("ImageUrl", imageUrl)
	}
	return trade
}
func (trade *TradeRequest) UseCreditInst(instFlgs ...InstFlag) *TradeRequest {
	flags := []string{}
	for _, instFlag := range instFlgs {
		flags = append(flags, string(instFlag))
	}
	if len(flags) > 0 {
		trade.Set("InstFlag", strings.Join(flags, ","))
	}
	return trade
}
func (trade *TradeRequest) UseCreditRed() *TradeRequest {
	trade.Set("CreditRed", "1")
	return trade
}
func (trade *TradeRequest) UseUnionPay() *TradeRequest {
	trade.Set("UNIONPAY", "1")
	return trade
}
func (trade *TradeRequest) UseWebATM() *TradeRequest {
	trade.Set("WEBATM", "1")
	trade.useATM = true
	return trade
}
func (trade *TradeRequest) UseVACC() *TradeRequest {
	trade.Set("VACC", "1")
	trade.useATM = true
	return trade
}
func (trade *TradeRequest) SetBankType(bankTypes ...BankType) *TradeRequest {
	if trade.useATM {
		types := []string{}
		for _, bktype := range bankTypes {
			types = append(types, string(bktype))
		}
		if len(types) > 0 {
			trade.Set("BankType", strings.Join(types, ","))
		}
	}
	return trade
}
func (trade *TradeRequest) UseCVS() *TradeRequest {
	trade.Set("CVS", "1")
	return trade
}
func (trade *TradeRequest) UseBARCODE() *TradeRequest {
	trade.Set("BARCODE", "1")
	return trade
}
func (trade *TradeRequest) UseESUNWallet() *TradeRequest {
	trade.Set("ESUNWALLET", "1")
	return trade
}

func (trade *TradeRequest) UseTaiwanPay() *TradeRequest {
	trade.Set("TAIWANPAY", "1")
	return trade
}

func (trade *TradeRequest) UseEZPay() *TradeRequest {
	trade.Set("EZPAY", "1")
	return trade
}

func (trade *TradeRequest) UseLogistics(lgsMethod LgsMethod, lgsType LgsType) *TradeRequest {
	trade.Set("CVSCOM", string(lgsMethod))
	trade.Set("LgsType", string(lgsType))
	return trade
}

func (trade *TradeRequest) UseEZPWeChat() *TradeRequest {
	trade.Set("EZPWECHAT", "1")
	return trade
}

func (trade *TradeRequest) UseEZPaliPay() *TradeRequest {
	trade.Set("EZPALIPAY", "1")
	return trade
}

func (trade *TradeRequest) GenerateHTML() (string, error) {
	info, sha := KeyEncrypt(trade.Encode(), trade.HashKey, trade.HashIV)
	params := map[string]string{
		"MerchantID":  trade.MerchantID,
		"TradeInfo":   info,
		"TradeSha":    sha,
		"Version":     "2.0",
		"EncryptType": "0",
	}
	return GenerateAutoSubmitHtmlForm(trade.ApiUrl, params)
}
