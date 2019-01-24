package main

import (
	"math/big"
	"fmt"
	"crypto/sha256"
	"bytes"
	"github.com/ethereum/go-ethereum/common/math"
)

//1. 定义一个ProofOfWork的结构
//- 区块
//- 目标值

type ProofOfWork struct {
	//- 区块
	block Block

	//- 目标值
	target big.Int
}

//2. 提供一个创建ProofOfWork的方法
func NewProofOfWork(block Block) (*ProofOfWork) {
	//正常来说，难度值是由一个数值推导出来的，最终展示位一个哈希值
	//我们这里为了简便，先直接写死一个哈希值，后续pow调用起来之后，我们再进行推导
	//难度值哈希：0001000000000000000000000000000000000000000000000000000000000000

	//bits 去推导这个targetStr

	//第一阶段：写成固定的难度值
	//targetStr := "0001000000000000000000000000000000000000000000000000000000000000"
	//bigIntTmp := big.Int{}
	//
	//bigIntTmp.SetString(targetStr, 16)
	//bigIntTmp.SetString(targetStr, 16)

	//第二阶段：给定难度值，推导出难度值
	//推导逻辑：
	// 0001000000000000000000000000000000000000000000000000000000000000 :  目标值
	// 0000000000000000000000000000000000000000000000000000000000000001 ： 初始值1
	// 0000000000000000000000000000000000000000000000000000000000000010 ： 左移一位(二进制4位)
	//10000000000000000000000000000000000000000000000000000000000000000 ： 向左移动256位
	//00001000000000000000000000000000000000000000000000000000000000000 ： 再向右移动16位

	//1. 定义一个数值1
	bigIntTmp := big.NewInt(1)
	//2. 向左移动256位
	////left shift
	//bigIntTmp.Lsh(bigIntTmp, 256)
	////3. 向右移动16位，得到0001000000000000000000000000000000000000000000000000000000000000
	//bigIntTmp.Rsh(bigIntTmp, 16)

	//直接左移256 - 16位, 16就是难度值
	//magic number: 魔数
	bigIntTmp.Lsh(bigIntTmp, 256-bits)

	pow := ProofOfWork{
		block:  block,
		target: *bigIntTmp,
	}

	return &pow
}

func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	//var blockInfo []byte
	b := pow.block

	//blockInfo = append(blockInfo, uint2Bytes(b.Version)...)
	//blockInfo = append(blockInfo, b.PrevBlockHash...)
	//blockInfo = append(blockInfo, b.Hash...)
	//
	//blockInfo = append(blockInfo, b.MerkelRoot...)
	//blockInfo = append(blockInfo, uint2Bytes(b.TimeStamp)...)
	//blockInfo = append(blockInfo, uint2Bytes(b.Bits)...)
	////这里要注意，一定不要使用b.Nonce，这是我们要寻找的数字
	//blockInfo = append(blockInfo, uint2Bytes(nonce)...)
	//blockInfo = append(blockInfo, b.Data...)

	tmp := [][]byte{
		uint2Bytes(b.Version),
		b.PrevBlockHash,
		//b.Hash,
		b.MerkelRoot,
		uint2Bytes(b.TimeStamp),
		uint2Bytes(b.Bits),
		uint2Bytes(nonce),
		b.Data,
	}

	blockInfo := bytes.Join(tmp, []byte{})

	return blockInfo
}

//3. 给ProofOfWork提供一个计算的方法，用于找到Nonce
func (pow *ProofOfWork) Run() (uint64, []byte) {

	//2. 定义一个nonce变量，用于不断变化
	var nonce uint64
	var hash [32]byte

	//for ; ; {
	for nonce <= math.MaxInt64 {
		fmt.Printf("%x\r", hash)

		//3. 对拼接好的数据进行sha256运算
		// - 得到是一个哈希值，需要转换为big.Int
		//hash := sha256.Sum256(blockInfo + nonce)
		hash = sha256.Sum256(pow.prepareData(nonce))

		// - 需要一个中间变量，将[32]byte转换为big.Int
		//func (z *Int) SetBytes(buf []byte) *Int

		var bitIntTmp big.Int

		bitIntTmp.SetBytes(hash[:]) //当前block与nonce拼接之后得到的哈希值

		//4. 得到目标值: pow.target
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y

		//- 如果生成哈希 小于 目标值，满足条件，返回哈希值，nonce，直接退出
		if bitIntTmp.Cmp(&pow.target) == -1 {
			fmt.Printf("挖矿成功，hash: %x, nonce : %d\n", hash, nonce)
			break
		} else {
			//- 如果生成哈希 大于 目标值，不满足条件，nonce++, 继续遍历
			//fmt.Printf("当前哈希: %x, %d\n", hash, nonce)
			nonce++
		}
	} // for

	return nonce, hash[:]
}

//4. 在NewBlock最后，进行调用！

//提供一个校验方法，用于检测挖矿得到的随机数是否满足系统的条件
func (pow *ProofOfWork) IsValid() bool {

	//得到数据
	//做哈希运算
	//与系统的难度值比较

	//A : 挖矿成功， B：校验之
	block := pow.block

	//矿工校验时，会拿到区块数据，然后自己校验哈希
	data := pow.prepareData(block.Nonce)
	fmt.Printf("---- isValid, Nonce : %d\n", block.Nonce)

	hash := sha256.Sum256(data)
	fmt.Printf("---- isvliad, block.hash : %x\n", block.Hash)
	fmt.Printf("---- isvliad, Hash : %x\n", hash)

	//哈希值与big.Int的比较
	var bigIntTmp big.Int

	bigIntTmp.SetBytes(hash[:])
	fmt.Printf("pow.target : %x\n", pow.target.Bytes())

	//return bigIntTmp.Cmp(&pow.target) == -1
	res := bigIntTmp.Cmp(&pow.target)
	fmt.Printf("bigIntTmp : %x\n", bigIntTmp.Bytes())

	//res := pow.target.Cmp(&bigIntTmp)

	fmt.Printf("res : %d\n", res)

	if bigIntTmp.Cmp(&pow.target) == -1 {
		fmt.Printf("111111\n")
		return true
	}

	return false
}
