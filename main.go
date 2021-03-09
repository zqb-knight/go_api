package main

import (
	"fmt"
	"go_api/router"
	"net/http"
)

func main() {
	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        r,
		ReadTimeout:    600,
		WriteTimeout:   600,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

}
