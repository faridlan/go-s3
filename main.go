package main

import (
	"net/http"

	"github.com/faridlan/go-s3/controller"
)

func main() {

	sm := http.NewServeMux()
	sm.HandleFunc("/", controller.UploadS3Controller)

	server := http.Server{
		Addr:    ":8080",
		Handler: sm,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
