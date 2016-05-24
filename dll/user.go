package dll

import (
	"net/http"
	"time"
	"strings"

	"1898/dal"
	"1898/utils"
)

//@name 检测邀请码
func CheckRegKey(w http.ResponseWriter, r *http.Request) {
	key := strings.ToUpper(r.FormValue("key"))

	if key == "" {
		Errors(w, ErrMissParam("key", ErrCode_UserMissParamKey))

		return
	}

	if key != "999999" {

		// 检测注册码是否使用
		var k = dal.Keys{}
		k.Key = key
		err := k.FindByKey()
		if err != nil {
			Errors(w, ErrForbidden("no registration key found", ErrCode_UserKeyNotFound))
			return
		}

		if k.UsedId != "" {
			Errors(w, ErrForbidden("the key has been used", ErrCode_UserKeyUsed))
			return
		}
	}

	Push(w, "the invitation code is right", "ok")
}

// @name 用户注册
// @method POST
// @uri /user/register
// @param phone 用户名
// @param password 密码
// @param key 注册码
// @success errcode:0
// @success id:ok
// @success msg:register success
// @success data:ok
// @fail errcode:403
// @fail id:miss_param
// @fail title:Miss Param
// @fail Detail:miss query param phone
func Register(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")

	if phone == "" {
		Errors(w, ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}

	if !utils.MatchPhone(phone) {
		Errors(w, ErrForbidden("must be the correct phone number", ErrCode_UserPhoneNotMatch))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		Errors(w, ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))

	key := strings.ToUpper(r.FormValue("key"))

	if key == "" {
		Errors(w, ErrMissParam("key", ErrCode_UserMissParamKey))

		return
	}

	id := NewObjectId()

	if key != "999999" {

		// 检测注册码是否使用
		var k = dal.Keys{}
		k.Key = key
		err := k.FindByKey()
		if err != nil {
			Errors(w, ErrForbidden("no registration key found", ErrCode_UserKeyNotFound))
			return
		}

		if k.UsedId != "" {
			Errors(w, ErrForbidden("the key has been used", ErrCode_UserKeyUsed))
			return
		}

		// 使用注册码
		k.UsedId = id.Hex()
		k.Used = time.Now()

		err = k.UpdateByKey(key)
		if err != nil {
			Errors(w, ErrForbidden(err.Error(), ErrCode_UpdateKeyErr))
			return
		}

		f := new(dal.Friends)
		f.UId = id
		f.Fid = k.UId
		f.Agree = time.Now()
		f.Created = time.Now()
		if err = f.Add(); err != nil {
			Errors(w, ErrForbidden(err.Error(), ErrCode_AddFriendErr))
			return
		} else {
			f_id := NewObjectId()
			f.Id = f_id
			f.UId = k.UId
			f.Fid = id
			f.Agree = time.Now()
			f.Created = time.Now()
			err = f.Add()
			if err != nil {
				Errors(w, ErrForbidden(err.Error(), ErrCode_AddFriendErr))
				return
			}
		}

	}

	u := new(dal.User)

	// 生成token
	token := u.NewToken()

	u.Id = id
	u.Phone = phone
	u.Password = password
	u.Token = token
	u.Created = time.Now()
	u.Updated = time.Now()

	err := u.AddUser()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "register success", u)

}

// @name 用户登录
// @method POST
// @uri /user/login
// @param phone 用户名
// @param password 密码
// @param key 注册码
// @success status:200
// @success id:ok
// @success msg:login success
// @success data:ok
// @fail status:403
// @fail id:miss_param
// @fail title:Miss Param
// @fail Detail:miss query param phone
func Login(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")

	if phone == "" {
		Errors(w, ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}

	if !utils.MatchPhone(phone) {
		Errors(w, ErrForbidden("must be the correct phone number", ErrCode_UserPhoneNotMatch))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		Errors(w, ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))

	u := new(dal.User)
	u.Phone = phone
	u.Password = password

	err := u.Login()

	if err != nil {
		Errors(w, ErrForbidden("login failed", ErrCode_InternalServer))

		return
	}

	Push(w, "login success", u)
}

// 根据用户id获取用户信息
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	u := new(dal.User)
	u.Id = ObjectIdHex(uid)

	err := u.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "find user success", u)

}


