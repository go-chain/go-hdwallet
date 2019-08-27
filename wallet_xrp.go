package hdwallet

import (
	"github.com/go-chain/go-xrp/crypto"
)


func init() {
	coins[XRP] = newXRP
}

type xrp struct {
	name   string
	symbol string
	key    *Key
}

func newXRP(key *Key) Wallet {
	return &xrp{
		name:   "ripple",
		symbol: "xrp",
		key:    key,
	}
}

func (c *xrp) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *xrp) GetName() string {
	return c.name
}

func (c *xrp) GetSymbol() string {
	return c.symbol
}

func (c *xrp) GetKey() *Key {
	return c.key
}

func (c *xrp) GetAddress() (string, error) {
	return GenerateAddress(c.key.Public.SerializeCompressed())
}

func GenerateAddress(b []byte) (string,error) {
	
	hash, err := crypto.NewAccountId(crypto.Sha256RipeMD160(b))
	if err != nil {
		return  "",err
	}
	return hash.String(),nil
}