package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/xiaofeng-zerone/flytiger/server/xinit"
)

func StartServer() error {

	Router := xinit.RouterInit()

	s := endless.NewServer(":6666", Router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
