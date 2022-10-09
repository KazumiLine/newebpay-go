# newebpay-go

## Description
NewebPay SDK for Golang, with some helpers package.

## Feature
- Configure
    - MerchantID
    - HashKey
    - HashIV
    - SimulationMode
- API
    - NewTradeRequest
    - ParseTradeResponse
    - QueryTradeInfo

## How to start
### Register
- Formal env
    - https://www.newebpay.com/website/Page/content/register
- Simulation env
    - https://cwww.newebpay.com/website/Page/content/register

### Install
```sh
go get github.com/KazumiLine/newebpay-go
```
```go
import (
    "github.com/KazumiLine/newebpay-go"
)
```
### Init client
```go
store := NewStore("StoreID", "Key", "IV", true)
```
### New Trade
```go
html, _ := store.NewTradeRequest("Order ID", amount, "Product Name").
    SetNotifyURL("callback url").
    SetEmail("custom@gmail.com").
    SetOrderComment("Order Comment").
    UseWebATM().UseVACC().UseBARCODE().UseCVS().
    // ...more options
    GenerateHTML()
```
### Order Callback
```go
func callback(w http.ResponseWriter, r *http.Request) {
    resp, err := store.ParseTradeResponse(r)
    fmt.Println(resp, err)
    // ...handle order resp
}
```
### Query Trade Info
```go
tradeInfo, err := store.QueryTradeInfo("Order ID", amount)
```

## More Info
[API doc](https://www.newebpay.com/website/Page/content/download_api)