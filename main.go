package main

import (
	"fmt"
	"time"
)

func main() {
	//block := NewBlock("hello world!", []byte{})
	bc := NewBlockChain()

	bc.AddBlock("hello 航头!")
	bc.AddBlock("再见 航头!")

	for i, block := range bc.Blocks {
		fmt.Printf(" ========= 区块高度 : %d =======\n", i)

		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevBlockHash)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)

		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)

		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %s\n", timeFormat)
		//fmt.Printf("时间戳: %d\n", block.TimeStamp)

		fmt.Printf("难度值: %d\n", block.Bits)
		fmt.Printf("随机数: %d\n", block.Nonce)

		pow := NewProofOfWork(*block)
		fmt.Printf("IsValid : %v\n", pow.IsValid())

		fmt.Printf("区块数据: %s\n", block.Data)
	}
}
