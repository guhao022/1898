package dll

import (
	"net/http"
	"strconv"
	"1898/dal"
	"time"
	"1898/utils"
)

// @name 创建活动
// @method POST
// @uri /event/new
// @param uid 创建用户id
// @param title 活动标题
// @param detail 活动内容
// @param addr 活动地址
// @param total 活动总共人数
// @success status:200
// @success id:ok
// @success msg:create event success
// @success data:ok
// @fail status:403
// @fail id:miss_param
// @fail title:Miss Param
// @fail Detail:miss query param title
func NewEvent (w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	title := r.FormValue("title")

	if title == "" {
		Errors(w, ErrMissParam("title", ErrCode_EventMissParamTitle))

		return
	}

	detail := r.FormValue("detail")

	if detail == "" {
		Errors(w, ErrMissParam("detail", ErrCode_EventMissParamDetail))

		return
	}

	if len(detail) < 10 {
		Errors(w, ErrForbidden("detail must be at least 10", ErrCode_EventDetailLenNotEnough))

		return
	}

	addr := r.FormValue("addr")

	if addr == "" {
		Errors(w, ErrMissParam("addr", ErrCode_EventMissParamAddr))

		return
	}

	total := r.FormValue("total")

	if total == "" {
		Errors(w, ErrMissParam("total", ErrCode_EventMissParamTotal))

		return
	}

	tn, err := strconv.Atoi(total);

	if err != nil {
		Errors(w, ErrForbidden("total must be int", ErrCode_EventTotalNotInt))

		return
	}

	var event = &dal.Event{}

	event.Uid = ObjectIdHex(uid)
	event.Title = title
	event.Detail = detail
	event.Addr = addr
	event.Total = tn
	event.Created = time.Now()

	err = event.AddEvent()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "create event success", "ok")

}

// @name 活动详情
func EventInfo(w http.ResponseWriter, r *http.Request) {
	eid := r.FormValue("eid")

	if eid == "" {
		Errors(w, ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		Errors(w, ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	err := event.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "event information", event)

}

//@name 活动报名
func RegEvent(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	eid := r.FormValue("eid")

	if eid == "" {
		Errors(w, ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		Errors(w, ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err := event.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	//检测活动人数是否已满
	total := event.SignUp.Len()
	if total >= event.Total {
		Errors(w, ErrForbidden("enrollment is full", ErrCode_EnrollmentFull))
		return
	}

	// 检测是否已经参加活动
	contain, _ := utils.ListContains(event.SignUp, uid)

	if contain {
		Errors(w, ErrForbidden("user already sign up", ErrCode_UserAlreadySignUp))
		return
	}

	event.SignUp.PushBack(uid)

	err = event.UpdateById(eid)

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "sign up event", "ok")

}

// @name 活动列表
func EventList(w http.ResponseWriter, r *http.Request) {
	var event dal.Event

	v, err := event.FindAll(0, 0, "created")

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "event list", v)
}

// @name 删除活动
func DelEvent(w http.ResponseWriter, r *http.Request) {
	eid := r.FormValue("eid")

	if eid == "" {
		Errors(w, ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		Errors(w, ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	err := event.DelById(eid)
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "event delete success", "ok")
}

// @name 取消参加活动
func CancelEvent(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	eid := r.FormValue("eid")

	if eid == "" {
		Errors(w, ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		Errors(w, ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err := event.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	// 检测是否已经参加活动
	contain, e := utils.ListContains(event.SignUp, uid)

	if !contain {
		Errors(w, ErrForbidden("user not sign up", ErrCode_UserNotSignUp))
		return
	}

	eval := event.SignUp.Remove(e)

	Push(w, "event cancel success", eval)
}


