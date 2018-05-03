package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"fmt"
)
  //go run ethereum/account/importkey.go
	//Address addr = "0x6e5ab887860e199b91b92d81f418c95d9ffd32cb"
  
func main(){
	var key = "40dad29726f7e1b56359d2f1cc5a5365cb105b410e1108b3da65c1d97bfe6f8e"
	var privkey = new(big.Int)
  
	privkey.SetString(key,16)
	pk, err := crypto.ToECDSA(privkey.Bytes())
	if err != nil {
            fmt.Println(err)
	    return
	}
  
	publicKey := pk.PublicKey;
	addr := crypto.PubkeyToAddress(publicKey)
  
	fmt.Println("publicKey = ", publicKey)
	fmt.Println("address = ", addr.Hex())
}
