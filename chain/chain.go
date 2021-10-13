package chain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash, PrevHash, Data string
}
type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AppendBlocks(data string) *Block {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := makeBlock(prevBlock.Hash, data)
	chain.Blocks = append(chain.Blocks, newBlock)
	return newBlock
}
func MakeBlockChain() *BlockChain {
	genesisBlock := makeBlock("000", "Genesis hash")
	chain := []*Block{genesisBlock}
	return &BlockChain{chain}
}
func makeBlock(prevHash, data string) *Block {
	return &Block{
		Hash:     makeHash(prevHash, data),
		PrevHash: prevHash,
		Data:     data,
	}
}
func makeHash(prevHash, data string) string {
	hash := sha256.Sum256([]byte(prevHash + data))
	return hex.EncodeToString(hash[:])
}
