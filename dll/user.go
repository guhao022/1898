package dll

import (
	"net/http"
	"time"
	"strings"

	"1898/dal"
	"1898/utils"
	"strconv"
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
	u.Username = phone
	u.Phone = phone
	u.Password = password
	u.Token = token
	u.Root = 0
	u.Created = time.Now()
	u.Updated = time.Now()

	err := u.AddUser()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "register success", u)

}

// @name 创建管理用户
func CreateRoot(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	if username == "" {
		Errors(w, ErrMissParam("phone", ErrCode_UserMissParamUsername))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		Errors(w, ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))



	id := NewObjectId()

	u := new(dal.User)

	// 生成token
	token := u.NewToken()

	u.Id = id
	u.Username = username
	u.Phone = username
	u.Password = password
	u.Token = token
	u.Root = 1
	u.Created = time.Now()
	u.Updated = time.Now()

	err := u.AddUser()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "register success", u)
}

// @name 修改密码
func EditPassword(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	oldpassword := r.FormValue("oldpassword")

	if oldpassword == "" {
		Errors(w, ErrMissParam("oldpassword", ErrCode_UserMissParamOldPassword))

		return
	}

	password := r.FormValue("password")

	if password == "" {
		Errors(w, ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	u := new(dal.User)
	u.Id = ObjectIdHex(uid)

	err := u.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	if u.Password != utils.Md5(utils.Md5(oldpassword)) {
		Errors(w, ErrForbidden("old password error", ErrCode_OldPwdErr))
		return
	}

	u.Password = utils.Md5(utils.Md5(password))

	err = u.UpdateById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "update password success", "ok")

}

// @name 修改用户信息
func EditUser(w http.ResponseWriter, r *http.Request) {
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

	nickname := r.FormValue("nickname")

	if nickname == "" {
		Errors(w, ErrMissParam("nickname", ErrCode_MissParamNickname))

		return
	}

	email := r.FormValue("email")

	if email == "" {
		email = u.Email
	}

	job := r.FormValue("job")

	if job == "" {
		job = u.Job
	}

	about := r.FormValue("about")

	if about == "" {
		about = u.About
	}

	company := r.FormValue("company")

	if company == "" {
		company = u.Company
	}

	pro := r.FormValue("pro")

	if pro == "" {
		pro = u.Profession
	}

	city := r.FormValue("city")

	if city == "" {
		city = u.City
	}

	expert := r.FormValue("expert")

	if expert == "" {
		expert = u.Expert
	}

	hobby := r.FormValue("hobby")

	if hobby == "" {
		hobby = u.Hobby
	}

	honor := r.FormValue("honor")

	if honor == "" {
		honor = u.Honor
	}

	age, err :=  strconv.Atoi(r.FormValue("age"))

	if err != nil {
		age = u.Age
	}
	sex, err :=  strconv.Atoi(r.FormValue("sex"))

	if err != nil {

		sex = u.Sex
	}

	u.Nickname = nickname
	u.Email = email
	u.Job = job
	u.About = about
	u.Company = company
	u.Expert = expert
	u.Hobby = hobby
	u.Profession = pro
	u.City = city
	u.Honor = honor
	u.Age = age
	u.Sex = sex
	u.Updated = time.Now()

	err = u.UpdateById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "update user success", u)
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

// @name 管理员登录
func RootLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")

	if username == "" {
		Errors(w, ErrMissParam("username", ErrCode_UserMissParamUsername))

		return
	}

	pwd := r.FormValue("password")

	if pwd == "" {
		Errors(w, ErrMissParam("password", ErrCode_UserMissParamPassword))

		return
	}

	password := utils.Md5(utils.Md5(pwd))

	u := new(dal.User)
	u.Username = username
	u.Password = password

	err := u.RootLogin()

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

// 根据用户id获取用户信息
func GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")

	if phone == "" {
		Errors(w, ErrMissParam("phone", ErrCode_UserMissParamPhone))

		return
	}


	u := new(dal.User)
	u.Phone = phone

	err := u.FindByPhone()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "find user success", u)

}

// 头像上传
func Avatar(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	if uid == "" {
		Errors(w, ErrMissParam("uid", ErrCode_MissParamUid))

		return
	}

	if !IsObjectId(uid) {
		Errors(w, ErrForbidden("uid must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	fp, err := utils.Upload(r)

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_UploadErr))
		return
	}

	user := new(dal.User)

	user.Id = ObjectIdHex(uid)
	err = user.FindByID()

	if err != nil {
		Errors(w, ErrInternalServer("not found user", ErrCode_UserNotFound))
		return
	}

	user.Avatar = fp

	err = user.UpdateById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}
	Push(w, "upload success", fp)
}



