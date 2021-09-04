package presentation

import (
	"log"
	"net/http"

	"github.com/nemotoy/golang-playground/ddd/server/application"
)

type userHandler struct {
	UserAppSrv *application.UserApplicationService
}

func NewUserHandler(userAppSrv *application.UserApplicationService) *userHandler {
	return &userHandler{userAppSrv}
}

func (u *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("dump: %v\n", u.UserAppSrv.GetAll())
	w.Write([]byte("hi!\n"))
}

func (u *userHandler) Post(w http.ResponseWriter, r *http.Request) {
}

func (u *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
}
