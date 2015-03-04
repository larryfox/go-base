# Rebase

Convert integers to a given base and back again. Accepts bases 2—62. Port of [`rebase-ruby`](https://github.com/larryfox/rebase-ruby).

~~~go
package main

import "github.com/larryfox/rebase-go"

func main() {
	Base62 := rebase.New(62)
	encoded, _ := Base62.Encode(10622433094)
	decoded, _ := Base62.Decode([]byte("Base62"))
	print(string(encoded), " : ", decoded) // => "Base62 : 10622433094"
}
~~~

### `New(int) *Rebase`
### `(*Rebase) Encode(int) ([]byte, error)`
### `(*Rebase) Decode([]byte) (int, error)`
### `IsBaseValid(int) bool`

### License
The MIT License. Copyright © 2014 Larry Fox
