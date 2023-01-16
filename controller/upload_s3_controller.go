package controller

import (
	"net/http"

	"github.com/faridlan/go-s3/helper"
	"github.com/faridlan/go-s3/model"
	"github.com/faridlan/go-s3/service"
)

func UploadS3Controller(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(err)
	}

	f, _, err := r.FormFile("testImage")
	if err != nil {
		panic(err)
	}

	// Filename := fileHeader.Filename

	defer f.Close()

	u := service.UploadS3(f)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   u,
	}

	helper.WriteToResponse(w, webResponse)
}
