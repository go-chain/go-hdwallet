package hdwallet

func init() {
	coins[DASH] = newDASH
}

type dash struct {
	*btc
}

func newDASH(key *Key) Wallet {
	if key.opt.Params == nil {
		key.opt.Params = &DASHParams
	}
	token := newBTC(key).(*btc)
	token.name = "Dash"
	token.symbol = "DASH"

	return &dash{btc: token}
}
