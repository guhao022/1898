package dal

import (
	"1898/utils"
	"time"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

// 新建token
func (t *User) NewToken() *Token {
	token := string(utils.RandomCreateBytes(16))

	ut := new(Token)

	ut.Token = token
	ut.Created = time.Now()
	ut.Expired = time.Now().Add(2*time.Hour)

	return ut
}

// 生成token
func (t *User) CreateToken(uid bson.ObjectId) (*Token, error) {
	token := t.NewToken()

	t.Token = token

	c := t.mgo()
	c.Query = bson.M{"_id": uid}
	c.Change = bson.M{"$set": t}

	err := c.Update()

	return token, err
}

// 根据token值获取token信息
func (t *User) FindTokenByUId() (*Token, error) {
	c := t.mgo()

	c.Query = bson.M{"_id": t.Id}

	err := c.Find(&t)
	if err != nil {
		return nil, err
	}

	if t.Token == nil {
		return nil, errors.New("the user's token is null")
	}

	return t.Token, err
}

// 检测是否过期
func (t *User) TokenExpire(uid, tokenval string) bool {

	t.Id = bson.ObjectIdHex(uid)

	token, err := t.FindTokenByUId()
	if err != nil {
		return true
	}

	if tokenval != token.Token {
		return true
	}

	now := time.Now()

	if now.After(token.Expired) {
		return true
	}

	return false
}

