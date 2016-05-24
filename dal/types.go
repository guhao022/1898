package dal

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// 用户
type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Phone    string        `bson:"phone" json:"phone"`
	Password string        `bson:"password" json:"password"`
	Salt     string        `bson:"salt" json:"salt"`
	Nickname string        `bson:"nickname" json:"nickname"`
	Email    string        `bson:"email" json:"email"`
	Job      string        `bson:"job" json:"job"`
	About    string        `bson:"about" json:"about"`
	Token 	 *Token        `bson:"token" json:"token"`
	Created  time.Time     `bson:"created" json:"created"`
	Updated  time.Time     `bson:"updated" json:"updated"`
	Deleted  time.Time     `bson:"deleted" json:"deleted"`
}

// 用户即时token
type Token struct {
	Token   string    `bson:"token" json:"token"`
	Created time.Time     `bson:"created" json:"created"`
	Expired  time.Time    `bson:"expired" json:"expired"`
}

// 好友
type Friends struct {
	Id    	bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UId    	bson.ObjectId `bson:"uid,omitempty" json:"uid"`
	Fid   	bson.ObjectId `bson:"fid,omitempty" json:"fid"`
	Agree 	time.Time   `bson:"agree" json:"agree"`
	Created time.Time    `bson:"created" json:"created"`
	Deleted  time.Time     `bson:"deleted" json:"deleted"`
}


// 注册码
type Keys struct {
	Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UId     bson.ObjectId `bson:"uid" json:"uid"`         // 生成注册码的用户
	UsedId  string        `bson:"usedid" json:"usedid"`   // 使用注册码的用户
	Key     string        `bson:"key" json:"key"`         // 注册码
	Created time.Time     `bson:"created" json:"created"` // 生成时间
	Used    time.Time     `bson:"used" json:"used"`       // 使用时间
}

// 新闻
type News struct {
	Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreateUser	*User `bson:"create_user" json:"create_user"`
	Title   string        `bson:"title" json:"title"`
	Detail  string        `bson:"detail" json:"detail"`
	Created time.Time     `bson:"created" json:"created"`
	Updated time.Time     `bson:"updated" json:"updated"`
	Deleted time.Time     `bson:"deleted" json:"deleted"`
}

// 活动
type Event struct {
	Id      bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	CreateUser    *User   `bson:"create_user" json:"create_user"`
	Title   string          `bson:"title" json:"title"`   // 活动标题
	Detail  string          `bson:"detail" json:"detail"` // 活动详情
	Addr    string          `bson:"addr" json:"addr"`     // 活动地址
	Price   int             `bson:"price" json:"price"`     // 活动价格
	Total   int             `bson:"total" json:"total"`     // 允许总参加人数
	SignUp map[string]string `bson:"signup" json:"signup"` // 已报名的用户
	Start 	string       	`bson:"start" json:"start"`	//开始时间
	Created time.Time       `bson:"created" json:"created"`
	Updated time.Time       `bson:"updated" json:"updated"`
	Deleted time.Time       `bson:"deleted" json:"deleted"`
}
