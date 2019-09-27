package hdwallet


import "math/big"

func EncodeBase58(b []byte) (s string) {

	const dict = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	x := new(big.Int).SetBytes(b)
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	for x.Cmp(zero) > 0 {
		x.QuoRem(x, m, r)
		s = string(dict[r.Int64()]) + s
	}

	for _, v := range b {
		if v != 0 {
			break
		}
		s = string(dict[0]) + s
	}

	return s
}
