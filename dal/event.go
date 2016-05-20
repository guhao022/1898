package dal

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

func (e *Event) mgo() *Mgo {
	return NewMgo("events")
}

// 查询所有活动
// skip 跳过前n个文档
// limit 返回文档个数
// sort 用法如("firstname", "-lastname")，优先按firstname正序排列，其次按lastname倒序排列
func (e *Event) FindAll(skip, limit int, sort ...string) (v []*Event, err error) {
	c := e.mgo()
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

// 根据ID查询活动
func (e *Event) FindByID() error {
	c := e.mgo()
	c.Query = bson.M{"_id": e.Id}
	return c.Find(&e)
}

// 根据用户名查询用户
func (e *Event) FindByTitle() error {
	c := e.mgo()
	c.Query = bson.M{"title": e.Title}
	return c.Find(&e)
}

// 检测活动名称
func (e *Event) CheckTitle() bool {
	err := e.FindByTitle()
	if err != nil {
		return false
	}

	return true
}

// 添加活动
func (e *Event) AddEvent() error {
	c := e.mgo()
	if e.CheckTitle() {
		return fmt.Errorf("event title cannot be repeated!")
	}
	return c.Insert(e)
}

// 根据ID更改活动信息
func (e *Event) UpdateById(id string) error {
	c := e.mgo()
	c.Query = bson.M{"_id": bson.IsObjectIdHex(id)}
	c.Change = bson.M{"$set": e}

	return c.Update()
}

// 根据ID删除活动
func (e *Event) DeleteById(id string) error {
	c := e.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return c.Remove()
}

// 软删除
func (e *Event) DelById(id string) error {
	e.Deleted = time.Now()

	return e.UpdateById(id)
}



