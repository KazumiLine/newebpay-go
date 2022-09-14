package newebpay

import (
	"fmt"
	"testing"
)

func TestNewTrade(t *testing.T) {
	store := NewStore("merchaneId", "hashKey", "hashIV", ApiUrl)
	html, _ := store.NewTradeRequest("hihihi", 100, "哈囉你好").
		SetBankType(BankType_FCBK).
		SetClientBackURL("url").
		SetEmail("email").
		SetLangType(LangType_TW).
		UseCVS().
		UseCreditInst(Installment_24_Month, Installment_30_Month).
		GenerateHTML()
	fmt.Println(html)
}
