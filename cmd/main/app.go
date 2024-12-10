package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"tren3/internal/user"
	"tren3/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register using handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port 1234")
	logger.Fatal(server.Serve(listener))
}
