package user

import (
	"net/http"
	"tren3/internal/handlers"
	"tren3/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} //подсказка

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("this is list of users"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("this is user by uuid"))

}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("this is create user"))

}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("this is update user"))

}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update user"))

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is delete user"))

}
