package dal

import "gopkg.in/mgo.v2/bson"

func (n *News) mgo() *Mgo {
	return NewMgo("news")
}

// 添加新闻
func (n *News) AddNews() error {
	c := n.mgo()

	return c.Insert(n)
}

// 新闻列表
func (n *News) FindAll(skip, limit int, sort ...string) (v []*Event, err error) {
	c := n.mgo()
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

// 新闻信息
func (n *News) FindById() error {
	c := n.mgo()
	c.Query = bson.M{"_id": n.Id}

	return c.Find(&n)
}

// 修改新闻
func (n *News) UpdateNewById() error {
	c := n.mgo()
	c.Query = bson.M{"_id": n.Id}

	c.Change = bson.M{"$set": n}

	return c.Update()
}

// 删除新闻
func (n *News) DeleteNewById(id string) error {
	c := n.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}
	return c.Remove()
}

// 新闻总数
func (n *News) Count() (int, error) {
	c := n.mgo()
	return c.Count()
}
