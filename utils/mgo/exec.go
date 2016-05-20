package mgo
import (
	"sync"
)

type Exec struct {
	Database   	string                  // 数据库
	Username	string					// 用户名
	Password	string					// 密码
	Collection 	string                  // 集合

	Query      	map[string]interface{}	// 查询语句
	Sort       	[]string               	// 排序
	Skip       	int                    	// 跳过前n个文档
	Limit 	   	int 					// 返回文档个数
	Select     	interface{}            	// 返回指定字段，如{"name":1}

	Change     map[string]interface{} 	// 文档更新内容

	lock  		sync.RWMutex
}


// 写入
func (e *Exec) Insert(docs ...interface{}) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	s, db := DB(e.Database, e.Username, e.Password)
	defer s.Close()

	c := db.C(e.Collection)

	return c.Insert(docs...)
}

// 查询单个
func (e *Exec) Find(v interface{}) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	s, db := DB(e.Database, e.Username, e.Password)
	defer s.Close()

	c := db.C(e.Collection)

	q := c.Find(e.Query)
	if e.Select != nil {
		q.Select(e.Select)
	}

	return q.One(v)
}

// 查询所有
func (e *Exec) FindAll(v interface{}) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	s, db := DB(e.Database, e.Username, e.Password)
	defer s.Close()

	c := db.C(e.Collection)

	q := c.Find(e.Query)

	if len(e.Sort) > 0 {
		q.Sort(e.Sort...)
	}
	if e.Skip > 0 {
		q.Skip(e.Skip)
	}
	if e.Limit > 0 {
		q.Limit(e.Limit)
	}
	if e.Select != nil {
		q.Select(e.Select)
	}

	return q.All(v)
}

// 修改
func (e *Exec) Update() error {
	e.lock.Lock()
	defer e.lock.Unlock()
	s, db := DB(e.Database, e.Username, e.Password)
	defer s.Close()

	c := db.C(e.Collection)

	_, err := c.UpdateAll(e.Query, e.Change)

	return err
}


// 删除
func (e *Exec) Remove() error {
	e.lock.Lock()
	defer e.lock.Unlock()
	s, db := DB(e.Database, e.Username, e.Password)
	defer s.Close()

	c := db.C(e.Collection)

	_, err := c.RemoveAll(e.Query)

	return err
}

