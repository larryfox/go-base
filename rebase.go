package rebase

var (
	alphabet62 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	alphabet36 = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
)

func New(base int) *Rebase {
	if !IsBaseValid(base) {
		panic(InvalidBaseError(base))
	}
	r := new(Rebase)
	if base > 36 {
		r.alphabet = alphabet62[0:base]
	} else {
		r.alphabet = alphabet36[0:base]
	}
	r.base = base
	r.index = alphabetIndex(r.alphabet)
	return r
}

type Rebase struct {
	alphabet []byte
	base     int
	index    map[byte]int
}

func (r *Rebase) Encode(n int) (b []byte, err error) {
	if n < 0 {
		return nil, NegativeIntegerError
	}
	if n == 0 {
		return r.alphabet[n:1], nil
	}
	for n > 0 {
		b = append([]byte{r.alphabet[n%r.base]}, b...)
		n = n / r.base
	}
	return b, nil
}

func (r *Rebase) Decode(bytes []byte) (n int, err error) {
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

func alphabetIndex(bytes []byte) map[byte]int {
	index := make(map[byte]int, len(bytes))
	for i, b := range bytes {
		index[b] = i
	}
	return index
}
