package hdwallet

func init() {
	coins[BCH] = newBCH
}

type bch struct {
	*btc
}

func newBCH(key *Key) Wallet {

	if key.opt.Params == nil {
		key.opt.Params = &LTCParams
	}
	token := newBTC(key).(*btc)
	token.name = "Bitcoin Cash"
	token.symbol = "BCH"
	token.key.opt.Params = &BCHParams

	return &bch{btc: token}
}
