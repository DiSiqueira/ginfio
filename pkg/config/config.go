package config

import (
	"github.com/spf13/viper"
)

type (
	Reader interface {
		Name() string
	}

	reader struct {
		getter Getter
	}

	Getter interface {
		Get(string) interface{}
	}
)

func New() (Reader, error) {
	v := viper.New()
	v.SetConfigName("ginfio")
	v.AddConfigPath("/etc/ginfio/")
	v.AddConfigPath("$HOME/.ginfio")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return NewReader(v), nil
}

func NewReader(getter Getter) Reader {
	return &reader{
		getter: getter,
	}
}

func (r *reader) Name() string {
	name, ok := r.getter.Get("name").(string)
	if !ok {
		return ""
	}
	return name
}
