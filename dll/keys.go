package dll

import (
	"net/http"
	"time"
	"1898/dal"
	"1898/utils/web"
)

// @name 生成注册码
func GetKey(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", ErrCode_MissParamUid))
		return
	}

	if !IsObjectId(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	var k = &dal.Keys{}

	key := k.CreateKey()

	k.Key = key
	k.Created = time.Now()
	k.UId = ObjectIdHex(uid)

	err := k.AddKey()
	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		panic(err)
		return
	}

	web.Push(w, "get key", map[string]string{"key": string(key)})

}
