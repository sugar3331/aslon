package core

import (
	"aslon/gee"
	"time"
)
import "github.com/fvbock/endless"

func initServer(address string, router *gee.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
