package controller

import (
	"encoding/json"
	"net/http"
	"project/model"

	"github.com/labstack/echo"
)

var data = []model.Article{
	{1, true, "Article 1", "Article 1 pro content"},
	{2, true, "Article 2", "Article 2 pro content"},
}

func ReadArcticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func Hello(e echo.Context) error {
	user := model.User{Name: "altaha", Email: "ada@deh"}
	return e.JSON(http.StatusOK, user)
}
