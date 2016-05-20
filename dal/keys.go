package dal

import (
	"1898/utils"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func (k *Keys) mgo() *Mgo {
	return NewMgo("keys")
}

// 查找key
func (k *Keys) FindByKey() error {
	c := k.mgo()
	c.Query = bson.M{"key": k.Key}
	return c.Find(&k)
}

//查重
func (k *Keys) CheckKey() bool {
	err := k.FindByKey()
	if err != nil {
		return false
	}

	return true
}

// 添加key
func (k *Keys) AddKey() error {
	c := k.mgo()
	return c.Insert(k)
}

// 生成key
func (k *Keys) CreateKey() string {
	key := utils.RandomCreateBytes(6)

	k.Key = strings.ToUpper(string(key))
	// 查重
	if k.CheckKey() {
		k.CreateKey()
	}

	return k.Key
}

// 修改key
func (k *Keys) UpdateByKey(key string) error {
	c := k.mgo()
	c.Query = bson.M{"key": key}
	c.Change = bson.M{"$set": k}

	return c.Update()
}
