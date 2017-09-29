package handler

import (
	"log"
	"net/http"

	"github.com/disiqueira/ginfio/pkg/web"
	"github.com/julienschmidt/httprouter"
)

type (
	Handler interface {
		ServeHTTP(http.ResponseWriter, *http.Request, httprouter.Params)
	}
)

func Start(version string) {
	jsonResponser := web.NewJSON()

	router := httprouter.New()
	router.GET("/", NewGetInfoHandler(jsonResponser, version).ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", router))
}
