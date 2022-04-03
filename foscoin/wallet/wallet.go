package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/jon890/foscoin/utils"
)

func Start() {
	// 1. Generate Keypair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	// fmt.Println("Private Key", privateKey.D)
	// fmt.Println("Public Key, x, y", privateKey.X, privateKey.Y)

	// 2. Hash message
	message := "I am hungry"
	hashedMessage := utils.Hash(message)
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 3. Sign message
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)
	//fmt.Printf("R:%d\nS:%d\n", r, s)

	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	fmt.Println(ok)
}
