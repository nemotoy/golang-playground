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
	users, _ := u.UserAppSrv.GetAll()
	log.Printf("dump: %v\n", users)
	w.Write([]byte("got all users!\n"))
}

func (u *userHandler) Post(w http.ResponseWriter, r *http.Request) {
	user, _ := u.UserAppSrv.Store()
	log.Printf("dump: %v\n", user)
	w.Write([]byte("stored given user!\n"))
}

func (u *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
}
