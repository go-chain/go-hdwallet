package hdwallet

import (
	"github.com/go-chain/go-tron/address"
)


func init() {
	coins[TRON] = newTRX
}

type trx struct {
	name   string
	symbol string
	key    *Key
}

func newTRX(key *Key) Wallet {
	return &trx{
		name:   "tron",
		symbol: "trx",
		key:    key,
	}
}

func (c *trx) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *trx) GetName() string {
	return c.name
}

func (c *trx) GetSymbol() string {
	return c.symbol
}

func (c *trx) GetKey() *Key {
	return c.key
}

func (c *trx) GetAddress() (string, error) {
	return address.FromPublicKey(c.key.PublicECDSA).ToBase58(),nil
}
