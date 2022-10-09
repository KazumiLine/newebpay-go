package newebpay

const TestTradeReqUrl string = "https://ccore.newebpay.com/MPG/mpg_gateway"
const TradeReqUrl string = "https://core.newebpay.com/MPG/mpg_gateway"
const TestTradeInfoUrl string = "https://ccore.newebpay.com/API/QueryTradeInfo"
const TradeInfoUrl string = "https://core.newebpay.com/API/QueryTradeInfo"

type LangType string

const (
	LangType_EN LangType = "en"
	LangType_TW LangType = "zh-tw"
	LangType_JP LangType = "jp"
)

func bool2str(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

type InstFlag string

const (
	Installment_3_Month  InstFlag = "3"  // 啟用分三期付款
	Installment_6_Month  InstFlag = "6"  // 啟用分六期付款
	Installment_12_Month InstFlag = "12" // 啟用分十二期付款
	Installment_18_Month InstFlag = "18" // 啟用分十八期付款
	Installment_24_Month InstFlag = "24" // 啟用分二十四期付款
	Installment_30_Month InstFlag = "30" // 啟用分三十期付款
)

type BankType string

const (
	BankType_TWBK BankType = "BOT"       // 臺灣銀行
	BankType_HNCB BankType = "HNCB"      // 華南銀行
	BankType_FCBK BankType = "FirstBank" // 第一銀行
)

type LgsMethod string

const (
	LgsMethod_NoNeedPay LgsMethod = "1"
	LgsMethod_NeedPay   LgsMethod = "2"
	LgsMethod_Both      LgsMethod = "3"
)

type LgsType string

const (
	LgsType_B2C LgsType = "B2C"
	LgsType_C2C LgsType = "C2C"
)

type TradeStatus string

const (
	TradeStatus_UNPAID    = "0"
	TradeStatus_PAID      = "1"
	TradeStatus_FAILED    = "2"
	TradeStatus_CANCELLED = "3"
)

func (p *TradeStatus) UnmarshalText(data []byte) error {
	switch string(data) {
	case "0":
		*p = "UNPAID"
	case "1":
		*p = "PAID"
	case "2":
		*p = "FAILED"
	case "3":
		*p = "CANCELLED"

	default:
		*p = "Unrecognized"
	}
	return nil
}
