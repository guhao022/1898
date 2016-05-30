package dal

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (f *Friends) mgo() *Mgo {
	return NewMgo("friends")
}

// 添加好友
func (f *Friends) Add() error {
	c := f.mgo()
	return c.Insert(f)
}

// 获取好友列表
func (f *Friends) FindByUid(uid string) ([]*Friends, error) {
	c := f.mgo()

	c.Query = bson.M{"uid": bson.ObjectIdHex(uid), "agree":bson.M{"$gt": new(time.Time)}, "deleted": new(time.Time)}

	var fs []*Friends

	err := c.FindAll(&fs)

	//println(err.Error())

	return fs, err
}

// 根据uid和fid查询好友
func (f *Friends) FindByUFID() error {
	c := f.mgo()
	c.Query = bson.M{"uid": f.UId, "fid": f.Fid}

	return c.Find(&f)
}

// 根据ID获取好友信息
func (f *Friends) FindByID() error {
	c := f.mgo()
	c.Query = bson.M{"_id": f.Id}

	return c.Find(&f)
}

// 修改
func (f *Friends) UpdateById(id string) error {
	c := f.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	c.Change = bson.M{"$set": f}

	return c.Update()
}

// 软删除
func (f *Friends) DelByUFid() error {
	err := f.FindByUFID()
	if err != nil {
		return err
	}
	f.Deleted = time.Now()

	return f.UpdateById(f.Id.Hex())
}

// 删除
func (f *Friends) DeleteById() error {
	c := f.mgo()
	c.Query = bson.M{"_id": f.Id}

	return c.Remove()
}

