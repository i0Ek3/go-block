package main

import (
	"fmt"
	"time"
	"strconv" //转换字符串
	"bytes"
	"crypto/sha256" //加密
)

type Block struct { //创建区块结构体
	Timer 	int64 	//时间戳
	Data 	[]byte 	//数据
	preHash []byte 	//前一个区块的哈希值
	Hash 	[]byte 	//该区块的哈希值
}

func (block *Block) calHash() { //计算当前区块的哈希值
	time := []byte(strconv.FormatInt(block.Timer, 10)) //将时间转换成特定格式，方便加密
	meta := bytes.Join([][]byte{time, block.Data, block.preHash}, []byte{}) //拼接时间，数据和前一个区块的哈希值, 用来生成新的哈希值
	hash := sha256.Sum256(meta)
	block.Hash = hash[:]
}

func Genesis() *Block {
	genesis := NewBlock("This is genesis block.", []byte{})
	return genesis
}

func NewBlock(data string, preHash []byte) *Block {
	block := Block{}
	block.Timer = time.Now().Unix()
	block.Data = []byte(data)
	block.preHash = preHash
	block.calHash()
	return &block
}

func main() {
	genesis := Genesis()
	fmt.Printf("The genesis block's hash is: %x", string(genesis.Hash))
}
