package main

// 定义区块链
type BlockChain struct {
	// 将区块放在数组中
	blocks []*Block
}


// 定义一个添加区块的方法
func (bc *BlockChain) AddBlock(data string)  {
	// 获取最后一个区块
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// 调用NewBlock函数新增区块
	newBlock := NewBlock(data, prevBlock.Hash)

	// 将新区块加入到区块链中，即加入数组中
	bc.blocks = append(bc.blocks, newBlock)
}


// 初始化区块链
func NewBlockchain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{NewGenesisBlock()},
	}
}