package handler

import (
	"github.com/disiqueira/ginfio/pkg/config"
	"github.com/disiqueira/ginfio/pkg/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	GetInfoHandler struct {
		responser web.Responser
		version   string
		name      string
	}

	Ginfio struct {
		Version string `json:"version"`
		Owner   string `json:"owner"`
	}

	Info struct {
		Ginfio Ginfio `json:"ginfio"`
	}
)

func NewGetInfoHandler(responser web.Responser, conf config.Reader, version string) Handler {
	return &GetInfoHandler{
		responser: responser,
		name:      conf.Name(),
		version:   version,
	}
}

func (h *GetInfoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	info := &Info{
		Ginfio: Ginfio{
			Version: h.version,
			Owner:   h.name,
		},
	}

	err := h.responser.Write(w, http.StatusOK, info)
	if err != nil {
		InternalServerError(w)
	}
}
