package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/middleware"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ShowParams ...
type ShowParams struct {
	ID uint `json:"id"`
}

// Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var teacher model.TeacherBTUser
	var param string
	var err error

	if params["userId"] == "" {
		j, ok := middleware.ParseJWT(w, r)
		if !ok {
			return
		} else {
			param = strconv.Itoa(int(j.UserID))
		}
	} else {
		param = params["userId"]
	}

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	err = db.DB.Where("id = ? OR username = ?", param, param).First(&teacher.User).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	err = db.DB.First(&teacher, model.Teacher{
		UserID: teacher.User.ID,
	}).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(teacher)

}
