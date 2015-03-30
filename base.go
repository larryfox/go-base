package base

var (
	base62table = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	base36table = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
)

func New(base int) (*Base, error) {
	if !IsBaseValid(base) {
		return nil, InvalidBaseError(base)
	}
	r := new(Base)
	if base > 36 {
		r.table = base62table[0:base]
	} else {
		r.table = base36table[0:base]
	}
	r.base = base
	r.index = tableIndex(r.table)
	return r, nil
}

type Base struct {
	table []byte
	base  int
	index map[byte]int
}

func (r *Base) Encode(n int) (b []byte, err error) {
	if n < 0 {
		return nil, NegativeIntegerError
	}
	if n == 0 {
		return r.table[n:1], nil
	}
	for n > 0 {
		b = append([]byte{r.table[n%r.base]}, b...)
		n = n / r.base
	}
	return b, nil
}

func (r *Base) Decode(bytes []byte) (n int, err error) {
	if len(bytes) == 0 {
		return 0, EmptyStringError
	}
	exp := 0
	for len(bytes) > 0 {
		i := len(bytes) - 1
		n += pow(r.base, exp) * r.index[bytes[i]]
		exp += 1
		bytes = bytes[0:i]
	}
	return n, nil
}

func tableIndex(bytes []byte) map[byte]int {
	index := make(map[byte]int, len(bytes))
	for i, b := range bytes {
		index[b] = i
	}
	return index
}
