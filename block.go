package main


import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)
type Block struct {
    Index        int
    Timestamp    string
    Data         string
    PreviousHash string
    Hash         string
}

// CalculateHash generates the hash for the block using its properties
func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}


// CreateBlock creates a new block using the previous block's hash
func createblock(previousBlock Block, data string) Block {
    var newBlock Block

    newBlock.Index = previousBlock.Index + 1
    newBlock.Timestamp = time.Now().String()
    newBlock.Data = data
    newBlock.PreviousHash = previousBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock
}
// GenesisBlock creates the first block in the blockchain
func HossineEthBlock() Block {
	var HossineEthBlock Block

	HossineEthBlock.Index = 0
	HossineEthBlock.Timestamp = time.Now().String()
	HossineEthBlock.Data = "Hossine Eth Block"
	HossineEthBlock.PreviousHash = ""
	HossineEthBlock.Hash = calculateHash(HossineEthBlock)

	return HossineEthBlock
}
