package dll

import (
	"net/http"
	"strconv"
	"1898/dal"
	"time"
	"1898/utils"
	"1898/utils/web"
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
		web.Errors(w, web.ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	title := r.FormValue("title")

	if title == "" {
		web.Errors(w, web.ErrMissParam("title", ErrCode_EventMissParamTitle))

		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))

	if err != nil {
		web.Errors(w, web.ErrMissParam("price", ErrCode_EventMissParamPrice))

		return
	}

	detail := r.FormValue("detail")

	if detail == "" {
		web.Errors(w, web.ErrMissParam("detail", ErrCode_EventMissParamDetail))

		return
	}

	if len(detail) < 10 {
		web.Errors(w, web.ErrForbidden("detail must be at least 10", ErrCode_EventDetailLenNotEnough))

		return
	}

	addr := r.FormValue("addr")

	if addr == "" {
		web.Errors(w, web.ErrMissParam("addr", ErrCode_EventMissParamAddr))

		return
	}

	total, err := strconv.Atoi(r.FormValue("total"))

	if err != nil {
		web.Errors(w, web.ErrMissParam("total", ErrCode_EventMissParamTotal))

		return
	}

	start := r.FormValue("start")
	if start == "" {
		web.Errors(w, web.ErrMissParam("start", ErrCode_EventMissParamStart))

		return
	}

	var event = &dal.Event{}

	event.Uid = ObjectIdHex(uid)
	event.Title = title
	event.Detail = detail
	event.Addr = addr
	event.Price = price
	event.Total = total
	event.Start = start
	event.Created = time.Now()

	err = event.AddEvent()

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "create event success", "ok")

}

//@name 修改活动
func EditEvent(w http.ResponseWriter, r *http.Request) {

	eid := r.FormValue("eid")

	if eid == "" {
		web.Errors(w, web.ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		web.Errors(w, web.ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	uid := r.FormValue("uid")

	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	title := r.FormValue("title")

	if title == "" {
		web.Errors(w, web.ErrMissParam("title", ErrCode_EventMissParamTitle))

		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))

	if err != nil {
		web.Errors(w, web.ErrMissParam("price", ErrCode_EventMissParamPrice))

		return
	}

	detail := r.FormValue("detail")

	if detail == "" {
		web.Errors(w, web.ErrMissParam("detail", ErrCode_EventMissParamDetail))

		return
	}

	if len(detail) < 10 {
		web.Errors(w, web.ErrForbidden("detail must be at least 10", ErrCode_EventDetailLenNotEnough))

		return
	}

	addr := r.FormValue("addr")

	if addr == "" {
		web.Errors(w, web.ErrMissParam("addr", ErrCode_EventMissParamAddr))

		return
	}

	total, err := strconv.Atoi(r.FormValue("total"))

	if err != nil {
		web.Errors(w, web.ErrMissParam("total", ErrCode_EventMissParamTotal))

		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err = event.FindByID()

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	// 检测是否是创建人
	if uid != event.Uid.Hex() {
		web.Errors(w, web.ErrForbidden("not create user", ErrCode_EventNotCreateUser))

		return
	}

	event.Title = title
	event.Detail = detail
	event.Addr = addr
	event.Price = price
	event.Total = total

	err = event.UpdateById(eid)

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "update event success", "ok")

}

// @name 活动详情
func EventInfo(w http.ResponseWriter, r *http.Request) {
	eid := r.FormValue("eid")

	if eid == "" {
		web.Errors(w, web.ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		web.Errors(w, web.ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err := event.FindByID()

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "event information", event)

}

//@name 活动报名
func RegEvent(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	eid := r.FormValue("eid")

	if eid == "" {
		web.Errors(w, web.ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		web.Errors(w, web.ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err := event.FindByID()
	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	//检测活动人数是否已满
	total := event.SignUp.Len()
	if total >= event.Total {
		web.Errors(w, web.ErrForbidden("enrollment is full", ErrCode_EnrollmentFull))
		return
	}

	// 检测是否已经参加活动
	contain, _ := utils.ListContains(event.SignUp, uid)

	if contain {
		web.Errors(w, web.ErrForbidden("user already sign up", ErrCode_UserAlreadySignUp))
		return
	}

	event.SignUp.PushBack(uid)

	err = event.UpdateById(eid)

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "sign up event", "ok")

}

// @name 活动列表
func EventList(w http.ResponseWriter, r *http.Request) {
	var event dal.Event

	v, err := event.FindAll(0, 0, "created")

	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "event list", v)
}

// @name 删除活动
func DelEvent(w http.ResponseWriter, r *http.Request) {
	eid := r.FormValue("eid")

	if eid == "" {
		web.Errors(w, web.ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		web.Errors(w, web.ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	err := event.DelById(eid)
	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "event delete success", "ok")
}

// @name 取消参加活动
func CancelEvent(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		web.Errors(w, web.ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		web.Errors(w, web.ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	eid := r.FormValue("eid")

	if eid == "" {
		web.Errors(w, web.ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		web.Errors(w, web.ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err := event.FindByID()
	if err != nil {
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	// 检测是否已经参加活动
	contain, e := utils.ListContains(event.SignUp, uid)

	if !contain {
		web.Errors(w, web.ErrForbidden("user not sign up", ErrCode_UserNotSignUp))
		return
	}

	eval := event.SignUp.Remove(e)

	web.Push(w, "event cancel success", eval)
}


