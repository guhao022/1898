// friend controller
package dll

import (
	"1898/dal"
	"net/http"
	"time"
)

// @name 添加好友
// @method POST
// @uri /friend/add
// @param uid 用户id dassdasda
// @param phone 用户电话号码
// @success errcode 0
// @success id ok
// @success msg add success
// @success data ok
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

	phone := r.FormValue("phone")

	if uid == "" {
		Errors(w, ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}

	u := new(dal.User)

	u.Phone = phone
	err := u.FindByPhone()

	if err != nil {
		Errors(w, ErrForbidden("user not found", ErrCode_UserNotFound))
		return
	}

	f := new(dal.Friends)
	f.UId = ObjectIdHex(uid)
	f.Fid = u.Id
	f.Created = time.Now()

	// 检测是否是本人
	if u.Id.Hex() == uid {
		Errors(w, ErrForbidden("你是分身吗？", ErrCode_FriendRepeat))
		return
	}

	// 检查好友是否重复
	err = f.FindByUFID()
	if err == nil {
		Errors(w, ErrForbidden("repeat", ErrCode_FriendRepeat))
		return
	}

	err = f.Add()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	fu := new(dal.Friends)
	fu.UId = u.Id
	fu.Fid = ObjectIdHex(uid)
	fu.Created = time.Now()

	// 检查好友是否重复
	err = fu.FindByUFID()
	if err == nil {
		Errors(w, ErrForbidden("repeat", ErrCode_FriendRepeat))
		return
	}

	err = fu.Add()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "add success", "ok")

}

// 同意加为好友
/*func FriendAgree(w http.ResponseWriter, r *http.Request) {

}*/

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
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	fid := r.FormValue("fid")

	if fid == "" {
		Errors(w, ErrMissParam("fid", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(fid) {
		Errors(w, ErrForbidden("fid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	f := new(dal.Friends)
	f.UId = ObjectIdHex(uid)
	f.Fid = ObjectIdHex(fid)

	err := f.FindByUFID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	err = f.DeleteById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	fu := new(dal.Friends)
	fu.UId = ObjectIdHex(fid)
	fu.Fid = ObjectIdHex(uid)

	err = fu.FindByUFID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}
	err = fu.DeleteById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "soft delete success", "ok")

}
