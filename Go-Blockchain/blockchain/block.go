package Blockchain

import (
	"bytes"
	"crypto/md5"
)

type Block struct {
	hash     string
	data     string
	prevHash string
}


func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.data), []byte(b.prevHash)}, []byte{})

	computedHash := md5.Sum(concatenatedData)

	b.hash = string(computedHash[:])
}

func CreateBlock(data string, prevHash string) *Block {
    block := &Block{"", data, prevHash}
    block.ComputeHash()
    return block
}

func Genesis() *Block {
    return CreateBlock("Genesis", "")
}