package web

import (
	"encoding/json"
	"net/http"
)

type (
	JSON struct {
	}
)

func NewJSON() *JSON {
	return &JSON{}
}

func (j *JSON) Write(w http.ResponseWriter, status int, val interface{}) error {
	if err := json.NewEncoder(w).Encode(val); err != nil {
		return err
	}
	return nil
}
