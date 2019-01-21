package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//1. 简单版（区块字段少）
//
//1. 定义结构（区块头的字段比正常的少）
type Block struct {
	//1. 版本号
	Version uint64

	//2. 前区块哈希
	PrevBlockHash []byte

	//2. 当前区块哈希, 这是为了方便加入的字段，正常区块中没有这个字段
	Hash []byte

	//3. 梅克尔根
	MerkelRoot []byte //先忽略不管

	//4. 时间戳, 从1970.1.1至今描述，一个数字
	TimeStamp uint64

	//5. 难度值，一个数字，可以推导出难度哈希值
	Bits uint64

	//6. 随机数Nonce，挖矿要求得值
	Nonce uint64

	//7. 数据
	Data []byte
}

//2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		Hash:          nil,
		MerkelRoot:    nil,
		TimeStamp:     uint64(time.Now().Unix()),
		Bits:          0, //随便写一个数
		Nonce:         0, //随便写一个数

		Data: []byte(data),
	}

	//设置哈希值
	block.setHash()

	return &block
}

//3. 生成哈希
// - 将所有的数据拼接起来，做sha256处理
func (b *Block) setHash() {
	var blockInfo []byte

	blockInfo = append(blockInfo, uint2Bytes(b.Version)...)
	blockInfo = append(blockInfo, b.PrevBlockHash...)
	blockInfo = append(blockInfo, b.Hash...)

	blockInfo = append(blockInfo, b.MerkelRoot...)
	blockInfo = append(blockInfo, uint2Bytes(b.TimeStamp)...)
	blockInfo = append(blockInfo, uint2Bytes(b.Bits)...)
	blockInfo = append(blockInfo, uint2Bytes(b.Nonce)...)

	blockInfo = append(blockInfo, b.Data...)

	//把新的字段添加进来

	hash := sha256.Sum256(blockInfo)

	b.Hash = hash[:]
}

//将数字转成字节流
func uint2Bytes(num uint64) []byte {

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)

	if err != nil {
		panic(err)
	}

	//binary.Read(bytes.NewReader(buffer.Bytes()), binary.BigEndian, ))

	return buffer.Bytes()
}

