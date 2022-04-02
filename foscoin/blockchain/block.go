package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jon890/foscoin/db"
	"github.com/jon890/foscoin/utils"
)

type Block struct {
	Hash         string `json:"hash"` // hash(data + prevHash)
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"dfficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

var ErrNotFound = errors.New("block not found")

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\n:Nonce%d\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Timestamp = int(time.Now().Unix())
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Hash:         "",
		PrevHash:     prevHash,
		Height:       height,
		Difficulty:   Blockchain().difficulty(),
		Nonce:        0,
		Transactions: []*Tx{makeCoinbaseTx("fos"), makeCoinbaseTx("fos")},
	}

	block.mine()
	block.persist()
	return block
}
