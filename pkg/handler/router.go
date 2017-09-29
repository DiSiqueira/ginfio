package handler

import (
	"log"
	"net/http"

	"github.com/disiqueira/ginfio/pkg/config"
	"github.com/disiqueira/ginfio/pkg/web"
	"github.com/julienschmidt/httprouter"
)

type (
	Handler interface {
		ServeHTTP(http.ResponseWriter, *http.Request, httprouter.Params)
	}
)

func Start(conf config.Reader, version string) {
	jsonResponser := web.NewJSON()

	getInfoHandler := NewGetInfoHandler(jsonResponser, conf, version)

	router := httprouter.New()
	router.GET("/", getInfoHandler.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", router))
}
