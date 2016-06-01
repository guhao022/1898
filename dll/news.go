package dll

import (
	"1898/dal"
	"1898/utils"
	"net/http"
	"strconv"
	"time"
)

// @name 添加新闻
func AddNews(w http.ResponseWriter, r *http.Request) {
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
		Errors(w, ErrMissParam("title", ErrCode_MissParamNewTitle))

		return
	}
	if len(title) > 60 {
		Errors(w, ErrForbidden("title must be at least 30 characters", ErrCode_StringLenErr))
		return
	}

	filepath := r.FormValue("image")
	if filepath == "" {
		Errors(w, ErrMissParam("filepath", ErrCode_MissParamImagePath))

		return
	}

	content := r.FormValue("content")
	if content == "" {
		Errors(w, ErrMissParam("content", ErrCode_MissParamContent))

		return
	}

	u := new(dal.User)
	u.Id = ObjectIdHex(uid)
	err := u.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	n := new(dal.News)
	n.Title = title
	n.Uid = u.Id
	n.Username = u.Nickname
	n.Image = filepath
	n.Content = content
	n.Created = time.Now()
	n.Updated = time.Now()

	err = n.AddNews()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", n)
}

// @name 修改新闻
func EditNews(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		Errors(w, ErrMissParam("id", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(id) {
		Errors(w, ErrForbidden("id must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	title := r.FormValue("title")

	if title == "" {
		Errors(w, ErrMissParam("title", ErrCode_MissParamNewTitle))

		return
	}
	if len(title) > 60 {
		Errors(w, ErrForbidden("title must be at least 30 characters", ErrCode_StringLenErr))
		return
	}

	content := r.FormValue("content")
	if content == "" {
		Errors(w, ErrMissParam("content", ErrCode_MissParamContent))

		return
	}

	filepath := r.FormValue("image")
	if filepath == "" {
		Errors(w, ErrMissParam("filepath", ErrCode_MissParamImagePath))

		return
	}

	n := new(dal.News)
	n.Id = ObjectIdHex(id)
	err := n.FindById()
	if err != nil {
		Errors(w, ErrInternalServer("news not found", ErrCode_NewsNotFound))
		return
	}

	n.Title = title
	n.Content = content
	n.Image = filepath
	n.Updated = time.Now()

	err = n.UpdateNewById()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", n)
}

//@name 删除新闻
func DelNews(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		Errors(w, ErrMissParam("id", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(id) {
		Errors(w, ErrForbidden("id must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	n := new(dal.News)

	if err := n.DeleteNewById(id); err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "delete news success", "ok")
}

//@name 新闻列表
func NewsList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	n := new(dal.News)
	count := n.Count()

	p := utils.NewPaging(count, 3).SetPage(page).Calc()

	news, err := n.FindAll(p.Offset, p.Limit, "updated")

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, strconv.Itoa(p.TotalPages), news)
}

//@name 新闻信息
func FindNews(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		Errors(w, ErrMissParam("id", ErrCode_MissParamId))

		return
	}

	if !IsObjectId(id) {
		Errors(w, ErrForbidden("id must be ObjectId format", ErrCode_UidNotObjectId))
		return
	}

	n := new(dal.News)
	n.Id = ObjectIdHex(id)

	err := n.FindById()

	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	Push(w, "get news success", n)
}
