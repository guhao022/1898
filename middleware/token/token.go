package token

import (
	"1898/dal"
	"1898/dll"
	"gopkg.in/mgo.v2/bson"

	"net/http"
)

type Token struct {}

func New() *Token {
	return &Token{}
}

func (t *Token) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){

	if r.RequestURI == "/login" || r.RequestURI == "/register" || r.RequestURI == "/regkey" {
		next(w, r)
		return
	}

	uid := r.FormValue("uid")
	if uid == "" {
		dll.Errors(w, dll.ErrMissParam("uid", dll.ErrCode_MissParamUid))

		return
	}

	if !bson.IsObjectIdHex(uid) {
		dll.Errors(w, dll.ErrForbidden("uid must be ObjectId format", dll.ErrCode_UidNotObjectId))
		return
	}

	tokenstr := r.FormValue("token")
	if uid == "" {
		dll.Errors(w, dll.ErrMissParam("token", dll.ErrCode_MissParamToken))

		return
	}

	u := new(dal.User)
	expire := u.TokenExpire(uid, tokenstr)
	if expire {
		dll.Errors(w, dll.ErrInternalServer("this token already expired", dll.ErrCode_TokenExpired))
		return
	}

	next(w, r)
}

