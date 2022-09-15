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

// 建立新訂單
//  - tradeNo  訂單編號
//  - amount   金額
//  - itemDesc 商品資訊
func (store *Store) NewTradeRequest(tradeNo string, amount int, itemDesc string) *TradeRequest {
	if tradeNo == "" {
		tradeNo = fmt.Sprintf("%d", time.Now().UnixMicro())
	}
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

// 語系
//  - 英文版 = en
//  - 繁體中文版 = zh-tw
//  - 日文版 = jp
func (trade *TradeRequest) SetLangType(lang LangType) *TradeRequest {
	trade.Set("LangType", string(lang))
	return trade
}

// 交易限制秒數
//  1.限制交易的秒數，當秒數倒數至 0 時，交易當做失敗
//  2.僅可接受數字格式
//  3.秒數下限為 60 秒，當秒數介於 1~59 秒時，會以 60 秒計算
//  4.秒數上限為 900 秒，當超過 900 秒時，會以 900 秒計算
func (trade *TradeRequest) SetTradeLimit(limit int) *TradeRequest {
	trade.Set("TradeLimit", fmt.Sprintf("%d", limit))
	return trade
}

// 繳費有效期限
//  1.僅適用於非即時支付
//  2.格式為 date('Ymd') ，例：20140620
//  3.此參數若為空值，系統預設為 7 天。自取號時間起算至第 7 天 23:59:59例：2014-06-23 14:35:51 完成取號，則繳費有效期限為 2014-06-29 23:59:59
//  4.可接受最大值為 180 天
func (trade *TradeRequest) SetExpireDate(date string) *TradeRequest {
	trade.Set("ExpireDate", date)
	return trade
}

// 支付完成返回商店網址
//  1.若支付工具為玉山 Wallet、台灣 Pay 或本欄位為空值，於交易完成後，消費者將停留在藍新金流付款或取號結果頁面
//  2.只接受 80 與 443 Port
func (trade *TradeRequest) SetReturnURL(uri string) *TradeRequest {
	trade.Set("ReturnURL", uri)
	return trade
}

// 支付通知網址
//  1.以幕後方式回傳給商店相關支付結果資料
//  2. 只接受 80 與 443 Port
func (trade *TradeRequest) SetNotifyURL(uri string) *TradeRequest {
	trade.Set("NotifyURL", uri)
	return trade
}

// 商店取號網址
//  1.此參數若為空值，則會顯示取號結果在藍新金流頁面
func (trade *TradeRequest) SetCustomerURL(uri string) *TradeRequest {
	trade.Set("CustomerURL", uri)
	return trade
}

// 返回商店網址
//  1.在藍新支付頁或藍新交易結果頁面上所呈現之返回鈕，我方將依據此參數之設定值進行設定，引導商店消費者依以此參數 網址返回商店指定的頁面
//  2.此參數若為空值時，則無返回鈕
func (trade *TradeRequest) SetClientBackURL(uri string) *TradeRequest {
	trade.Set("ClientBackURL", uri)
	return trade
}

// 付款人電子信箱
func (trade *TradeRequest) SetEmail(email string) *TradeRequest {
	trade.Set("Email", email)
	return trade
}

// 設定付款人電子信箱不可修改
func (trade *TradeRequest) DisableEmailModify() *TradeRequest {
	trade.Set("EmailModify", "0")
	return trade
}

// 須要登入藍新金流會員
func (trade *TradeRequest) NeedLogin() *TradeRequest {
	trade.Set("LoginType", "1")
	return trade
}

// 商店備註
//  1.限制長度為 300 字
func (trade *TradeRequest) SetOrderComment(comment string) *TradeRequest {
	trade.Set("OrderComment", comment)
	return trade
}

// 信用卡一次付清啟用
func (trade *TradeRequest) UseCreditCard() *TradeRequest {
	trade.Set("CREDIT", "1")
	return trade
}

// Google Pay 啟用
func (trade *TradeRequest) UseGooglePay() *TradeRequest {
	trade.Set("ANDROIDPAY", "1")
	return trade
}

// Samsung Pay 啟用
func (trade *TradeRequest) UseSamsungPay() *TradeRequest {
	trade.Set("SAMSUNGPAY", "1")
	return trade
}

// LINE Pay
//  - imageUrl 此連結的圖檔將顯示於 LINE Pay 付款前的 產品圖片區，若無產品圖檔連結網址，會使 用藍新系統預設圖檔。
func (trade *TradeRequest) UseLinePay(imageUrl string) *TradeRequest {
	trade.Set("LINEPAY", "1")
	if len(imageUrl) > 0 {
		trade.Set("ImageUrl", imageUrl)
	}
	return trade
}

// 信用卡分期付款啟用
//  - Installment_3_Month   啟用分三期付款
//  - Installment_6_Month   啟用分六期付款
//  - Installment_12_Month  啟用分十二期付款
//  - Installment_18_Month  啟用分十八期付款
//  - Installment_24_Month  啟用分二十四期付款
//  - Installment_30_Month  啟用分三十期付款
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

// 信用卡紅利支付方式啟用
func (trade *TradeRequest) UseCreditRed() *TradeRequest {
	trade.Set("CreditRed", "1")
	return trade
}

// 信用卡銀聯卡啟用
func (trade *TradeRequest) UseUnionPay() *TradeRequest {
	trade.Set("UNIONPAY", "1")
	return trade
}

// WEBATM 啟用
func (trade *TradeRequest) UseWebATM() *TradeRequest {
	trade.Set("WEBATM", "1")
	trade.useATM = true
	return trade
}

// ATM 轉帳啟用
func (trade *TradeRequest) UseVACC() *TradeRequest {
	trade.Set("VACC", "1")
	trade.useATM = true
	return trade
}

// 金融機構
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

// 超商代碼繳費啟用
func (trade *TradeRequest) UseCVS() *TradeRequest {
	trade.Set("CVS", "1")
	return trade
}

// 超商條碼繳費啟用
func (trade *TradeRequest) UseBARCODE() *TradeRequest {
	trade.Set("BARCODE", "1")
	return trade
}

// 玉山Wallet啟用
func (trade *TradeRequest) UseESUNWallet() *TradeRequest {
	trade.Set("ESUNWALLET", "1")
	return trade
}

// 台灣Pay啟用
func (trade *TradeRequest) UseTaiwanPay() *TradeRequest {
	trade.Set("TAIWANPAY", "1")
	return trade
}

// 簡單付電子錢包啟用
func (trade *TradeRequest) UseEZPay() *TradeRequest {
	trade.Set("EZPAY", "1")
	return trade
}

// 簡單付微信支付
func (trade *TradeRequest) UseEZPWeChat() *TradeRequest {
	trade.Set("EZPWECHAT", "1")
	return trade
}

// 簡單付支付寶
func (trade *TradeRequest) UseEZPaliPay() *TradeRequest {
	trade.Set("EZPALIPAY", "1")
	return trade
}

// 物流啟用(LgsMethod, LgsType)
// LgsMethod:
//   - LgsMethod_NoNeedPay 啟用超商取貨不付款
//   - LgsMethod_NeedPay   啟用超商取貨付款
//   - LgsMethod_Both      啟用超商取貨不付款及超商取貨付款
// LgsType:
//   - LgsType_B2C 大宗寄倉(目前僅支援 7-ELEVEN)
//   - LgsType_C2C 店到店(目前支援 7-ELEVEN、全家)
func (trade *TradeRequest) UseLogistics(lgsMethod LgsMethod, lgsType LgsType) *TradeRequest {
	trade.Set("CVSCOM", string(lgsMethod))
	trade.Set("LgsType", string(lgsType))
	return trade
}

// 產生自動送出的網頁
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
