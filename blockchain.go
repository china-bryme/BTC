package main

//4.引入区块链,使用数组来模拟区块链

type BlockChain struct {
	Blocks []*Block
}


//5.创建区块链
func NewBlockChain () *BlockChain  {

	genesisBlock := NewBlock(genesisInfo, []byte{})

	//创建区块链结构,一般会在创建的时候,添加一个区块,称之为:创世区块
	bc := BlockChain{
		Blocks: []*Block{genesisBlock},

	}

	return &bc

}

//6.添加区块
func (bc *BlockChain) AddBlock(data string) {
	//创建一个区块,前区块的世袭制从bc的最后一个区块元素获得即可
	lastBlock := bc.Blocks[(len(bc.Blocks) - 1)]

	//即将添加的区块的前哈希值,就是bc红的最后区块的hash字段的值
	prevHash := lastBlock.Hash
	newBlock := NewBlock(data, prevHash)

	//追加到区块链的Blocks数组中
	bc.Blocks = append(bc.Blocks,newBlock)

}