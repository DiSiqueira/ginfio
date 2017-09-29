package web

import "net/http"

type (
	Responser interface {
		Write(http.ResponseWriter, int, interface{}) error
	}
)
