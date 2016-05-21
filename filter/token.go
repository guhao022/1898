package filter

import (
	"1898/dal"
	"1898/utils/web"
	"net/http"
	"1898/dll"
	"gopkg.in/mgo.v2/bson"
)

// 检测token是否已过期
func CheckToken(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")
	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", dll.ErrCode_MissParamUid))

		return
	}

	if !bson.IsObjectIdHex(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", dll.ErrCode_UidNotObjectId))
		return
	}

	token := r.FormValue("token")
	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", dll.ErrCode_MissParamToken))

		return
	}

	u := new(dal.User)
	expire := u.TokenExpire(uid, token)
	if expire {
		web.Errors(w, web.ErrInternalServer("this token already expired", dll.ErrCode_TokenExpired))
		return
	}

}

