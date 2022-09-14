package newebpay

type Store struct {
	MerchantID string
	HashKey    string
	HashIV     string
	ApiUrl     string
}

func NewStore(merchantID, hashKey, hashIV, apiUrl string) (store *Store) {
	store = new(Store)
	store.MerchantID = merchantID
	store.HashKey = hashKey
	store.HashIV = hashIV
	store.ApiUrl = apiUrl
	return store
}
