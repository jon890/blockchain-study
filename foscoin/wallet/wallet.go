package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/jon890/foscoin/utils"
)

const (
	privateKey    = "30770201010420946539b844e42f19ffe8b78724d7fb1af462af1841b7f8591dce1ffffcacd758a00a06082a8648ce3d030107a1440342000409415138a382d8dd45bb0e0d3fe2249e48eee1951a570f952a0351c9fb452f701b2b3c106338ae26416d651f121499d7c4aea9e85ef93d4e30f57d394a4a4880"
	hashedMessage = "a68c33f6b0eed19a07a21ae8e576ae570dafa489df93b8e13a92d0ea2ca43a2c"
	signature     = "65dc77d184cf3fe9dfda92433f8cf3c8cfbef666a43258795db2dff91d952ac8da3c67041ff442efeb4bfa3096d48a222ef0adcf51d7a9dba00b9219cbb4cbf5"
)

func Start() {
	privateKeyAsByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoredKey, err := x509.ParseECPrivateKey(privateKeyAsByte)
	utils.HandleErr(err)

	fmt.Println(restoredKey)

	signBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := signBytes[:len(signBytes)/2]
	sBytes := signBytes[len(signBytes)/2:]

	fmt.Println(signBytes, rBytes, sBytes)

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
