package handler

import (
	"encoding/json"
	users "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/proto/user"
	"net/http"
	"time"
)

type newUser struct {
	userName  string
	email     string
	mobileNo  string
	birthDate time.Time
	password  string
}

func (h *Handler) getAllUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	res := jsonResponse{}
	encoder := json.NewEncoder(rw)
	result, err := h.userSvc.GetAllUsers(&users.ReadAllReq{})
	if err != nil {
		res.Err = err.Error()
		rw.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(res)
		return
	}
	res.Data = result
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(res)
}

func (h *Handler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	res := jsonResponse{}
	encoder := json.NewEncoder(rw)
	decoder := json.NewDecoder(r.Body)
	user := &users.User{}
	decoder.Decode(&user)
	result, err := h.userSvc.CreateUser(&users.CreateUserReq{User: user})

	if err != nil {
		res.Err = err.Error()
		rw.WriteHeader(http.StatusBadRequest)
		encoder.Encode(res)
		return
	}
	res.Data = result
	rw.WriteHeader(http.StatusCreated)
	encoder.Encode(res)
}
