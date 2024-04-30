package main

import(
	"time"
	"bytes"
	"strconv"
	"crypto/sha256"
)

func (Block *block) setHash(){
	timestamp := []byte(Block.Timestamp)
	headers := bytes.Join([][]byte{Block.PreviousBlockHash,Block.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)
	Block.Hash = hash[:]
}

// create a new eth block
func createBlock(Data string,PreviousBlockHash []byte) *Block{
	Block := &Block{time.Now().String(),PreviousBlockHash,[]byte{},[]byte(Data)}
	Block.setHash()
	return Block
}

func NeWHossineBlock