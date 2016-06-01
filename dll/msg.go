package dll

import (
	"1898/dal"
	"net/http"
	"time"
)

//@name 发送消息
func PushMsg(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	toid := r.FormValue("toid")

	if toid == "" {
		Errors(w, ErrMissParam("toid", ErrCode_MissParamToId))

		return
	}

	if !IsObjectId(toid) {
		Errors(w, ErrForbidden("toid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	// 检测是否是朋友关系
	f := new(dal.Friends)
	f.UId = ObjectIdHex(uid)
	f.Fid = ObjectIdHex(toid)
	err := f.FindByUFID()

	if err != nil /* || f.Agree.IsZero()*/ {
		Errors(w, ErrForbidden("not frients", ErrCode_NotFriend))
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		Errors(w, ErrMissParam("msg", ErrCode_MissParamMsg))

		return
	}

	if len(msg) > 252 {
		Errors(w, ErrForbidden("title must be at least 250 characters", ErrCode_StringLenErr))
		return
	}

	m := new(dal.Message)

	m.SendUId = ObjectIdHex(uid)
	m.GetUId = ObjectIdHex(toid)
	m.Msg = msg
	m.Read = 0
	m.Created = time.Now()

	err = m.AddMsg()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", m)
}

//@name 拉取消息
func PullMsg(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	m := new(dal.Message)
	ms, err := m.FindGetByUid(uid)

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	nms := make([]*dal.Message, 0)

	for _, v := range ms {
		u := new(dal.User)
		u.Id = v.SendUId
		u.FindByID()
		v.Nickname = u.Nickname

		nms = append(nms, v)
	}

	Push(w, "get news success", nms)
}

//@name 阅读消息
func ReadMsg(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		Errors(w, ErrMissParam("id", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(id) {
		Errors(w, ErrForbidden("id must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	m := new(dal.Message)

	err := m.FindById(id)
	if err != nil {
		Errors(w, ErrInternalServer("message not found", ErrCode_InternalServer))
		return
	}

	m.Read = 1
	err = m.UpdateById(id)
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "read message", "ok")

}
