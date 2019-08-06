package hdwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/eoscanada/eos-go/ecc"
)


func init() {
	coins[EOS] = newEOS
}

type eos struct {
	name   string
	symbol string
	key    *Key
}

func newEOS(key *Key) Wallet {
	return &eos{
		name:   "eos",
		symbol: "eos",
		key:    key,
	}
}

func (c *eos) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *eos) GetName() string {
	return c.name
}

func (c *eos) GetSymbol() string {
	return c.symbol
}

func (c *eos) GetKey() *Key {
	return c.key
}

func (c *eos) GetAddress() (string, error) {

	wif, err := btcutil.NewWIF(c.GetKey().Private,&chaincfg.MainNetParams,true)
	if err != nil {
		panic(err)
	}

	acc, err := ecc.NewPrivateKey(wif.String())
	if err != nil {
		return "",err
	}

	return acc.PublicKey().String(),nil
}

