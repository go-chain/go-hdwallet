package hdwallet

func init() {
	coins[BCH] = newBCH
}

type bch struct {
	*btc
}

func newBCH(key *Key) Wallet {

	if key.opt.Params == nil {
		key.opt.Params = &BCHParams
	}
	token := newBTC(key).(*btc)
	token.name = "Bitcoin Cash"
	token.symbol = "BCH"

	return &bch{btc: token}
}
