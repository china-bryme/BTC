package main

import (
	"../lib/bolt"
	"fmt"
)

const dbName = "test.db"
const bucketName1 = "bucketName1"

func main() {

	//1. 创建数据库
	//777
	//111, 111, 111
	//110,
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	//如果没有，自动创建
	db, err := bolt.Open(dbName, 0600, nil)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//2. 操作数据库
	//- Update: 修改
	//- View: 只读

	//func (db *DB) Update(fn func(*Tx) error) error {
	//这是一个事务，用于完成本次操作
	db.Update(func(tx *bolt.Tx) error {

		//1. 检查bucket是否存在
		b := tx.Bucket([]byte(bucketName1))

		// - 不存在：先创建bucket
		if b == nil {
			b, err = tx.CreateBucket([]byte(bucketName1))
			if err != nil {
				panic(err)
			}
		}

		//找到了bucket
		//2. 写入数据
		b.Put([]byte("111"), []byte("hello"))
		b.Put([]byte("222"), []byte("world"))

		//3. 读取数据
		v1 := b.Get([]byte("111"))
		v2 := b.Get([]byte("222"))
		v3 := b.Get([]byte("333"))

		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)
		fmt.Printf("v3 : %s\n", v3)

		return nil
	})
}
