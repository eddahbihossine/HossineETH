package main

type Block struct{
	TimeStamp string
	PreviousBlockHash []byte
	Hash []byte
	data[]byte
	
}
type BlockChain struct{
	Blocks []*Block
}