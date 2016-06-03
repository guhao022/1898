package dll

import (
	"1898/dal"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
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
func NewEvent(w http.ResponseWriter, r *http.Request) {

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

	price, err := strconv.Atoi(r.FormValue("price"))

	if err != nil {
		Errors(w, ErrMissParam("price", ErrCode_EventMissParamPrice))

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

	image := r.FormValue("image")

	if image == "" {
		Errors(w, ErrMissParam("image", ErrCode_EventMissParamImage))

		return
	}

	total, err := strconv.Atoi(r.FormValue("total"))

	if err != nil {
		Errors(w, ErrMissParam("total", ErrCode_EventMissParamTotal))

		return
	}

	start, err := strconv.ParseInt(r.FormValue("start"), 10, 64)
	if err != nil {
		Errors(w, ErrMissParam("start", ErrCode_EventMissParamStart))

		return
	}

	var u = new(dal.User)

	u.Id = ObjectIdHex(uid)

	err = u.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	var event = &dal.Event{}

	event.Uid = u.Id
	event.Username = u.Nickname
	event.Title = title
	event.Image = image
	event.Detail = detail
	event.Addr = addr
	event.Price = price
	event.Total = total
	event.SignUp = nil
	event.Start = start
	event.Created = time.Now()
	err = event.AddEvent()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "create event success", "ok")

}

//@name 修改活动
func EditEvent(w http.ResponseWriter, r *http.Request) {

	eid := r.FormValue("eid")

	if eid == "" {
		Errors(w, ErrMissParam("eid", ErrCode_MissParamEid))

		return
	}

	if !IsObjectId(eid) {
		Errors(w, ErrForbidden("eid must be ObjectId format", ErrCode_EidNotObjectId))
		return
	}

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

	price, err := strconv.Atoi(r.FormValue("price"))

	if err != nil {
		Errors(w, ErrMissParam("price", ErrCode_EventMissParamPrice))

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

	total, err := strconv.Atoi(r.FormValue("total"))

	if err != nil {
		Errors(w, ErrMissParam("total", ErrCode_EventMissParamTotal))

		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err = event.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	// 检测是否是创建人
	if uid != event.Uid.Hex() {
		Errors(w, ErrForbidden("not create user", ErrCode_EventNotCreateUser))

		return
	}

	event.Id = ObjectIdHex(eid)
	event.Title = title
	event.Detail = detail
	event.Addr = addr
	event.Price = price
	event.Total = total

	err = event.UpdateById()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "update event success", "ok")

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

	event.Id = ObjectIdHex(eid)

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

	u := new(dal.User)
	u.Id = ObjectIdHex(uid)
	err := u.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	if u.Nickname == "" {
		Errors(w, ErrForbidden("nickname is required", ErrCode_NickNameErr))
		return
	}

	var event = &dal.Event{}

	event.Id = ObjectIdHex(eid)

	err = event.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	// 是否已过报名期
	now := time.Now().Unix()

	if event.Start - 2*60*60 <= now {
		Errors(w, ErrForbidden("end of registration time", ErrCode_TimeOver))
		return
	}

	// 检查是否为组织者
	if ObjectIdHex(uid) == event.Uid {
		Errors(w, ErrInternalServer("the event organizer", ErrCode_EventOrganizer))
		return
	}

	//检测活动人数是否已满
	total := len(event.SignUp)

	if total >= event.Total {
		Errors(w, ErrForbidden("enrollment is full", ErrCode_EnrollmentFull))
		return
	}
	if total != 0 {
		// 检测是否已经参加活动
		if v, ok := event.SignUp[uid]; ok {
			Errors(w, ErrForbidden("user "+v+"already join", ErrCode_UserAlreadySignUp))
			return
		}
	}

	var sign = make(map[string]string, 0)

	sign[uid] = u.Nickname
	event.Id = ObjectIdHex(eid)
	event.SignUp = sign

	err = event.UpdateById()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "sign up event", "ok")

}

// @name 活动列表
func EventList(w http.ResponseWriter, r *http.Request) {
	var event dal.Event

	query := bson.M{}

	v, err := event.FindAll(0, 0, query, "created")

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
	if _, ok := event.SignUp[uid]; !ok {
		Errors(w, ErrForbidden("user not join this event", ErrCode_UserNotSignUp))
		return
	}

	delete(event.SignUp, uid)

	err = event.UpdateById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "event cancel success", "ok")
}

// 我发表的活动
func MyPublishEvent(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	var event = new(dal.Event)

	query := bson.M{"uid": ObjectIdHex(uid)}

	v, err := event.FindAll(0, 0, query, "created")

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "event list", v)
}

// 我参加的活动
func MyJoinEvent(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	var event *dal.Event

	query := bson.M{}

	v, err := event.FindAll(0, 0, query, "created")

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	var es = make([]*dal.Event, 0)

	for _, e := range v {
		if _, ok := e.SignUp[uid]; ok {
			es = append(es, e)
		}
	}

	Push(w, "event list", es)

}
