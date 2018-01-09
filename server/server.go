package server

import (
	"github.com/VinkDong/redirect/types"
	"net/http"
)

var Context types.Context

func initHandler() {
	http.HandleFunc("/redirect", redirect)
}

func Run() {
	config := Context.Config
	initHandler()
	if config.EnableSSL {
		err := http.ListenAndServeTLS(":443", config.Cert, config.Key, nil)
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServe(":8381", nil)
		if err != nil {
			panic(err)
		}
	}
}
