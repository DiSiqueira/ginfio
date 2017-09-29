package main

import (
	"github.com/disiqueira/ginfio/pkg/config"
	"github.com/disiqueira/ginfio/pkg/handler"
)

const version = "0.0.1"

func main() {
	config.Read()
	handler.Start(version)
}
