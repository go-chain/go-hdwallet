package hdwallet

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/btcsuite/btcutil"
)

func init() {
	coins[ZEN] = newZEN
}

const (
	btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var ZEN_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x20,0x89}, Suffix: nil}


type zen struct {
	name   string
	symbol string
	key    *Key
}

func newZEN(key *Key) Wallet {
	return &zen{
		name:   "ZenCash",
		symbol: "zen",
		key:    key,
	}
}

func (c *zen) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *zen) GetName() string {
	return c.name
}

func (c *zen) GetSymbol() string {
	return c.symbol
}

func (c *zen) GetKey() *Key {
	return c.key
}

func (c *zen) GetAddress() (string,error) {
	//fmt.Println(hex.EncodeToString(c.key.Public.SerializeCompressed()))
	//fmt.Println(hex.EncodeToString(btcutil.Hash160(c.key.Public.SerializeCompressed())))
	return addressEncoder.AddressEncode(btcutil.Hash160(c.key.Public.SerializeCompressed()), ZEN_mainnetAddressP2PKH),nil
}
