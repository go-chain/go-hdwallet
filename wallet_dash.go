package hdwallet

func init() {
	coins[DASH] = newDASH
}

type dash struct {
	*btc
}

func newDASH(key *Key) Wallet {
	if key.opt.Params == nil {
		key.opt.Params = &BCHParams
	}
	token := newBTC(key).(*btc)
	token.name = "Dash"
	token.symbol = "DASH"
	token.key.opt.Params = &DASHParams

	return &dash{btc: token}
}
