package dll

import (
	"net/http"
	"time"
	"strings"

	"1898/dal"
	"1898/utils"
	"1898/utils/web"
)

//@name 检测邀请码
func CheckRegKey(w http.ResponseWriter, r *http.Request) {
	key := strings.ToUpper(r.FormValue("key"))

	if key == "" {
		web.Errors(w, web.ErrMissParam("key", ErrCode_UserMissParamKey))

		return
	}

	if key != "999999" {

		// 检测注册码是否使用
		var k = dal.Keys{}
		k.Key = key
		err := k.FindByKey()
		if err != nil {
			web.Errors(w, web.ErrForbidden("no registration key found", ErrCode_UserKeyNotFound))
			return
		}

		if k.UsedId != "" {
			web.Errors(w, web.ErrForbidden("the key has been used", ErrCode_UserKeyUsed))
			return
		}
	}

	web.Push(w, "the invitation code is right", "ok")
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
		web.Errors(w, web.ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}

	if !utils.MatchPhone(phone) {
		web.Errors(w, web.ErrForbidden("must be the correct phone number", ErrCode_UserPhoneNotMatch))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		web.Errors(w, web.ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))

	key := strings.ToUpper(r.FormValue("key"))

	if key == "" {
		web.Errors(w, web.ErrMissParam("key", ErrCode_UserMissParamKey))

		return
	}

	id := NewObjectId()

	if key != "999999" {

		// 检测注册码是否使用
		var k = dal.Keys{}
		k.Key = key
		err := k.FindByKey()
		if err != nil {
			web.Errors(w, web.ErrForbidden("no registration key found", ErrCode_UserKeyNotFound))
			return
		}

		if k.UsedId != "" {
			web.Errors(w, web.ErrForbidden("the key has been used", ErrCode_UserKeyUsed))
			return
		}

		// 使用注册码
		k.UsedId = id.Hex()
		k.Used = time.Now()

		err = k.UpdateByKey(key)
		if err != nil {
			web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
			return
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
		web.Errors(w, web.ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	web.Push(w, "register success", u)

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
		web.Errors(w, web.ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}

	if !utils.MatchPhone(phone) {
		web.Errors(w, web.ErrForbidden("must be the correct phone number", ErrCode_UserPhoneNotMatch))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		web.Errors(w, web.ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))

	u := &dal.User{}
	u.Phone = phone
	u.Password = password

	err := u.Login()

	if err != nil {
		web.Errors(w, web.ErrForbidden("login failed", ErrCode_InternalServer))

		return
	}

	web.Push(w, "login success", u)
}


