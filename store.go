package newebpay

type Store struct {
	MerchantID     string
	HashKey        string
	HashIV         string
	SimulationMode bool
}

func NewStore(merchantID, hashKey, hashIV string, sim bool) (store *Store) {
	store = new(Store)
	store.MerchantID = merchantID
	store.HashKey = hashKey
	store.HashIV = hashIV
	store.SimulationMode = sim
	return store
}
