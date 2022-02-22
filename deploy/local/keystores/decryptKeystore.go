package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	inPaths := [3]string{"./0-0x4dce5c8961e283786cb31ad7fc072347227d7ea2", "./1-0x9251e7932e2c941e0ee1f370a1c387754af9cfdb", "2-0x96932b7a373d8586c4a2d3c98517803ff2818cec"}
	outPaths := [3]string{"prikey_0.hex", "prikey_1.hex", "prikey_2.hex"}
	password := "123456"
	for i := 0; i < 3; i++ {
		keyjson, e := ioutil.ReadFile(inPaths[i])
		if e != nil {
			panic(e)
		}
		key, e := keystore.DecryptKey(keyjson, password)
		if e != nil {
			panic(e)
		}
		e = crypto.SaveECDSA(outPaths[i], key.PrivateKey)
		if e != nil {
			panic(e)
		}
		fmt.Println("Key saved to:", outPaths[i])
	}
}
