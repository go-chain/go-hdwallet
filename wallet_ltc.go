package hdwallet

func init() {
	coins[LTC] = newLTC
}

type ltc struct {
	*btc
}

func newLTC(key *Key) Wallet {

	if key.opt.Params == nil {
		key.opt.Params = &LTCParams
	}

	token := newBTC(key).(*btc)
	token.name = "Litecoin"
	token.symbol = "LTC"

	return &ltc{btc: token}
}
