package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"tren3/internal/user"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("create user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server is listening port 1234")
	log.Fatalln(server.Serve(listener))
}
