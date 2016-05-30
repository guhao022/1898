package dal

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) mgo() *Mgo {
	return NewMgo("users")
}

// 根据用户名查询用户
func (u *User) FindByPhone() error {
	c := u.mgo()
	c.Query = bson.M{"phone": u.Phone}
	return c.Find(&u)
}

// 根据ID查询用户
func (u *User) FindByID() error {
	c := u.mgo()
	c.Query = bson.M{"_id": u.Id}
	return c.Find(&u)
}

// 用户登录
func (u *User) Login() error {
	c := u.mgo()
	c.Query = bson.M{"phone": u.Phone, "password": u.Password}
	return c.Find(&u)
}
// 管理用户登录
func (u *User) RootLogin() error {
	c := u.mgo()
	c.Query = bson.M{"username": u.Username, "password": u.Password, "root": 1}
	return c.Find(&u)
}

// 用户查重
func (u *User) CheckPhone() bool {
	err := u.FindByPhone()
	if err != nil {
		return false
	}

	return true
}

// 添加用户
func (u *User) AddUser() error {
	c := u.mgo()
	if u.CheckPhone() {
		return fmt.Errorf("phone number cannot be repeated!")
	}
	return c.Insert(u)
}

// 查询所有用户
// @param sort 排序
// @param sel select 返回指定字段，如{"name":1} 只返回name字段 {"name":0} 不返回name字段
// @param limit 返回文档个数
// @param skip 跳过文档个数
func (u *User) FindAll(skip, limit int, sel interface{}, sort ...string) (v []*User, err error) {
	c := u.mgo()
	if len(sort) > 0 {
		c.Sort = sort
	}
	if skip > 0 {
		c.Skip = skip
	}
	if limit > 0 {
		c.Limit = limit
	}

	err = c.FindAll(&v)
	return v, err
}

// 根据用户名修改用户
func (u *User) UpdateByName(username string) error {
	c := u.mgo()
	c.Query = bson.M{"username": username}
	c.Change = bson.M{"$set": u}

	return c.Update()
}

// 根据ID更改用户信息
func (u *User) UpdateById() error {
	c := u.mgo()
	c.Query = bson.M{"_id": u.Id}
	c.Change = bson.M{"$set": u}

	return c.Update()
}

// 根据名称删除用户
func (u *User) DelByName(username string) error {
	c := u.mgo()
	c.Query = bson.M{"username": username}

	return c.Remove()
}

// 根据ID删除用户
func (u *User) DelById(id string) error {
	c := u.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return c.Remove()
}
