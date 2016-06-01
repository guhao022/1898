package mgo

import (
	"fmt"
	"github.com/num5/pool"
	"gopkg.in/mgo.v2"
	"os"
)

var mgoPool = pool.NewPool(new(Mgo))

type Mgo struct {
	Session *mgo.Session
}

func (*Mgo) New() pool.Entity {
	session, err := mgo.Dial(os.Getenv("MGO_REMOTE"))
	if err != nil {
		fmt.Printf("数据库连接失败: %s\n", err)
	}

	return &Mgo{Session: session}
}

// 判断连接有效性
func (m *Mgo) Usable() bool {
	if m.Session.Ping() != nil {
		return false
	}
	return true
}

func (m *Mgo) Close() {
	m.Session.Close()
}

func (m *Mgo) Clean() {}

func Session() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MGO_REMOTE"))
	if err != nil {
		fmt.Printf("数据库连接失败: %s\n", err)
	}
	return session.Clone()
}

func DB(database, username, password string) (*mgo.Session, *mgo.Database) {

	session := Session()

	db := session.DB(database)
	if db == nil {
		session.Close()
		fmt.Println("数据库连接失败!")
	}
	err := db.Login(username, password)

	if err != nil {
		fmt.Printf("mongodb登陆失败: %s\n", err)
	}

	return session, db
}
