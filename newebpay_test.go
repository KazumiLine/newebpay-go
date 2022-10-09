package newebpay

import (
	"fmt"
	"net/http"
	"testing"
)

var store = NewStore("StoreID", "Key", "IV", true)

func TestNewTrade(t *testing.T) {
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		html, _ := store.NewTradeRequest("", 100, "Product Info").
			SetNotifyURL("callback url").
			SetEmail("custom@gmail.com").
			SetOrderComment("Order Comment").
			UseWebATM().UseVACC().UseBARCODE().UseCVS().
			GenerateHTML()
		fmt.Println(html)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	})
}

func TestReadResponse(t *testing.T) {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		resp, err := store.ParseTradeResponse(r)
		fmt.Println(resp, err)
	})
}
