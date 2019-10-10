package hdwallet

import "github.com/iost-official/go-iost/account"

func init() {
	coins[IOST] = newIOST
}

type iost struct {
	name   string
	symbol string
	key    *Key
}

func newIOST(key *Key) Wallet {
	return &iost{
		name:   "Internet of Services",
		symbol: "IOST",
		key:    key,
	}
}

func (c *iost) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *iost) GetName() string {
	return c.name
}

func (c *iost) GetSymbol() string {
	return c.symbol
}

func (c *iost) GetKey() *Key {
	return c.key
}

func (c *iost) GetAddress() (string,error)  {
	return account.EncodePubkey(c.key.Public.SerializeCompressed()),nil
}