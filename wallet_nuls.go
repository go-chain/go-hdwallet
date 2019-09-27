package hdwallet

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"errors"
)

func init() {
	coins[NULS] = newNULS
}

type nuls struct {
	name   string
	symbol string
	key    *Key
}

func newNULS(key *Key) Wallet {
	return &nuls{
		name:   "nulsio",
		symbol: "nuls",
		key:    key,
	}
}

func (c *nuls) GetType() uint32 {
	return c.key.opt.CoinType
}

func (c *nuls) GetName() string {
	return c.name
}

func (c *nuls) GetSymbol() string {
	return c.symbol
}

func (c *nuls) GetKey() *Key {
	return c.key
}

func (c *nuls) GetAddress() (string, error) {
	return GetAddressByPub(c.key.Public.SerializeCompressed())
}

const (
	addType = 1
	prefix  = "NULS"
)

var (
	constant = []rune{'a', 'b', 'c', 'd', 'e'}
)

//sha256之后hash160
func Sha256hash160(pub []byte) []byte {
	h := sha256.New()
	h.Write(pub)
	hasher := ripemd160.New()
	hasher.Write(h.Sum(nil))
	hashBytes := hasher.Sum(nil)
	return hashBytes
}

func GetAddressByPub(pub []byte) (string, error) {
	pubPart := Sha256hash160(pub)
	if len(pubPart) != 20 {
		return "", errors.New("pubPart len not 20")
	}
	chainPart := ShortToBytes(1)
	resultPart1 := make([]byte, 23)
	for index, v := range chainPart {
		resultPart1[index] = v
	}
	resultPart1[2] = addType
	for index, v := range pubPart {
		resultPart1[index+3] = v
	}
	xor := GetXor(resultPart1)
	resultPart2 := make([]byte, 24)
	for index, v := range resultPart1 {
		resultPart2[index] = v
	}
	resultPart2[23] = xor
	resultBytes := EncodeBase58(resultPart2)
	return prefix + string(constant[len(prefix)-1]) + string(resultBytes), nil
}

//异或方法
func GetXor(body []byte) byte {
	var xor byte
	xor = 0x00
	for i := 0; i < len(body); i++ {
		xor = byte(xor) ^ body[i]
	}
	return xor
}

//chainid转换
func ShortToBytes(val int) []byte {
	bytes := make([]byte, 2)
	bytes[1] = (byte)(0xFF & (val >> 8))
	bytes[0] = (byte)(0xFF & (val >> 0))
	return bytes
}