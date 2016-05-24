package dll

import (
	"net/http"
	"1898/dal"
	"time"
	"strconv"
	"1898/utils"
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

	title := r.FormValue("Title")

	if title == "" {
		Errors(w, ErrMissParam("title", ErrCode_MissParamNewTitle))

		return
	}
	if len(title) > 30 {
		Errors(w, ErrForbidden("title must be at least 30 characters", ErrCode_StringLenErr))
		return
	}

	content := r.FormValue("content")
	if content == "" {
		Errors(w, ErrMissParam("content", ErrCode_MissParamContent))

		return
	}

	u := new(dal.User)
	err := u.FindByID()
	if err != nil {
		Errors(w, ErrInternalServer(err.Error(), ErrCode_InternalServer))
		return
	}

	n := new(dal.News)
	n.Title = title
	n.CreateUser = u
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

	title := r.FormValue("Title")

	if title == "" {
		Errors(w, ErrMissParam("title", ErrCode_MissParamNewTitle))

		return
	}
	if len(title) > 30 {
		Errors(w, ErrForbidden("title must be at least 30 characters", ErrCode_StringLenErr))
		return
	}

	content := r.FormValue("content")
	if content == "" {
		Errors(w, ErrMissParam("content", ErrCode_MissParamContent))

		return
	}

	n := new(dal.News)
	err := n.FindById()
	if err != nil {
		Errors(w, ErrInternalServer("news not found", ErrCode_NewNotFound))
		return
	}

	n.Title = title
	n.Content = content
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
		page = 0
	}

	n := new(dal.News)
	count, err := n.Count()

	p := utils.NewPaging(count, 10).SetPage(page).Calc()

	n.FindAll(p.Offset, p.Limit, "updated")

	Push(w, "get news list success", n)
}

