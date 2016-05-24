package dll

import (
	"net/http"
	"1898/dal"
	"time"
)

// 添加好友
func AddFriend(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	fid := r.FormValue("fid")

	if uid == "" {
		Errors(w, ErrMissParam("fid", ErrCode_MissParamFid))

		return
	}

	if !IsObjectId(fid) {
		Errors(w, ErrForbidden("fid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	f := new(dal.Friends)
	f.UId = ObjectIdHex(uid)
	f.Fid = ObjectIdHex(fid)
	f.Created = time.Now()

	// 检查好友是否重复
	err := f.FindByUFID()
	if err == nil {
		Errors(w, ErrForbidden("repeat", ErrCode_FriendRepeat))
		return
	}

	err = f.Add()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "register success", "ok")

}

// 删除好友
//func DelFriend(w http.ResponseWriter, r *http.Request)

