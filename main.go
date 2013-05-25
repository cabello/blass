package main

import (
	"fmt"
    "bloom"
    "random"
)

func main() {
    var keys []string = make([]string, 100)

    var b bloom.Filter = bloom.New(250, 0.001)

    rs := random.NewRS("abcdefghijklmnopqrstuvxz")
    for i := 0; i < 100; i++ {
        keys[i] = rs.NewRandomString(16)
        b.Add([]byte(keys[i]))
    }

    fmt.Printf("%+v\n", b);

    fmt.Printf("%v\n", b.Contains([]byte("guilherme")))
    fmt.Printf("%v\n", b.Contains([]byte(keys[0])))
    fmt.Printf("%v\n", b.Contains([]byte(keys[50])))
    fmt.Printf("%v\n", b.Contains([]byte(keys[99])))
    fmt.Printf("%v\n", b.Contains([]byte("adalberto")))
    fmt.Printf("%v\n", b.Contains([]byte(keys[23])))
    fmt.Printf("%v\n", b.Contains([]byte(keys[44])))
    fmt.Printf("%v\n", b.Contains([]byte(keys[11])))
}
