package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 定义区块结构
type Block struct {
	// 时间戳
	Timestamp int64

	// 区块数据
	Data []byte

	// 前区块哈希
	PrevBlockHash []byte

	// 当前区块哈希
	Hash []byte
}

// 设置当前区块哈希，将timestamp + data + prevblockhash拼接起来
func (b *Block) SetHash() {
	// Timestamp是int64类型，需要转换为[]byte
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// 用bytes.join拼接[]byte
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	// 计算hash
	hash := sha256.Sum256(headers)

	// 赋值给b.Hash
	b.Hash = hash[:]
}

// 初始化Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data: []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash: []byte{},
	}
	return block
}


// 定义区块链第一个区块：创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}