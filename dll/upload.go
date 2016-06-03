package dll

import (
	"net/http"
	"1898/dal"
	"os"
	"1898/utils"
)

func Upload(w http.ResponseWriter, r *http.Request) {
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

	furl := os.Getenv("IMAGE_URL") + fp

	Push(w, "upload success", furl)
}
