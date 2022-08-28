package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

func main() {
	fnSignature := []byte("exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160))")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(fnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xdb3e2198
}
