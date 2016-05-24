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

// 好友列表
func FriendsList(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	f := new(dal.Friends)

	fs, err := f.FindByUid(uid)
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get friends success", fs)

}

// 删除好友
func DelFriend(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("uid")

	if id == "" {
		Errors(w, ErrMissParam("id", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(id) {
		Errors(w, ErrForbidden("id must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	f := new(dal.Friends)
	err := f.DelByid(id)

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "soft delete success", "ok")

}

