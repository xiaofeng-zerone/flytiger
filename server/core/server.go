package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello handler.")
	w.Write([]byte("Hello flytiger!"))
}

func StartServer() error {

	mux1 := mux.NewRouter()
	mux1.HandleFunc("/hello", handler).
		Methods("GET")

	s := endless.NewServer(":6666", mux1)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
