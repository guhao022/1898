package dal

import "gopkg.in/mgo.v2/bson"

func (m *Message) mgo() *Mgo {
	return NewMgo("message")
}

// 发送消息
func (m *Message) AddMsg() error {
	c := m.mgo()
	return c.Insert(m)
}

// 根据消息ID获取消息
func (m *Message) FindById(id string) error {
	c := m.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return c.Find(&m)
}

// 根据用户id获取用户所有发送的消息
func (m *Message) FindSendByUid(uid string) ([]*Message,error) {
	var ms []*Message
	c := m.mgo()
	c.Query = bson.M{"sendid", bson.ObjectIdHex(uid)}

	err := c.Find(&ms)

	return ms, err
}

// 根据用户id获取用户所有收到的信息
func (m *Message) FindGetByUid(uid string) ([]*Message,error) {
	var ms []*Message
	c := m.mgo()
	c.Query = bson.M{"getid", bson.ObjectIdHex(uid)}

	err := c.Find(&ms)

	return ms, err
}