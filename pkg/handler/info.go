package handler

import (
	"github.com/disiqueira/ginfio/pkg/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	GetInfoHandler struct {
		responser web.Responser
		version   string
	}

	Info struct {
		Version string `json:"version"`
	}
)

func NewGetInfoHandler(responser web.Responser, version string) Handler {
	return &GetInfoHandler{
		responser: responser,
		version:   version,
	}
}

func (h *GetInfoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	info := &Info{
		Version: h.version,
	}

	err := h.responser.Write(w, http.StatusOK, info)
	if err != nil {
		InternalServerError(w)
	}
}
