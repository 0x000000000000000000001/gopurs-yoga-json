package main

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	b := big.NewInt(10)
	out, _ := json.Marshal(b)
	fmt.Println(string(out))
}
