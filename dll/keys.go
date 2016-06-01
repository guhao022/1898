package dll

import (
	"1898/dal"
	"net/http"
	"time"
)

// @name 生成注册码
func GetKey(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))
		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	var k = &dal.Keys{}

	key := k.CreateKey()

	k.Key = key
	k.Created = time.Now()
	k.UId = ObjectIdHex(uid)

	err := k.AddKey()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get key", map[string]string{"key": string(key)})

}
