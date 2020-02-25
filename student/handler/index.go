package student

import (
	customHTTP "backend-qrcode/http"

	"backend-qrcode/db"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var student []Student

	if err := db.DB.Debug().Preload("User").Find(&student).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(student)
}