package hdwallet


func init() {
	coins[WICC] = newWicc
}

type wicc struct {
	*btc
}

func newWicc(key *Key) Wallet {

	if key.opt.Params == nil {
		key.opt.Params = &WICCParams
	}
	token := newBTC(key).(*btc)
	token.name = "WaykiChain"
	token.symbol = "WICC"

	return &wicc{btc: token}
}