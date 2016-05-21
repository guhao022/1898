package dll

import (
	"gopkg.in/mgo.v2/bson"
)

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

func IsObjectId(id string) bool {
	return bson.IsObjectIdHex(id)
}

func ObjectIdHex(id string) bson.ObjectId {
	return bson.ObjectIdHex(id)
}

