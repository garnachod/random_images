package user

import (
    "encoding/json"
    "fmt"
    "github.com/garnachod/random_images/internal"
    "log"
    "net/http"
)

type Handler interface{
    Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
    service Service
}

func NewUserHandler(service Service) Handler {
    return &handler{
        service,
    }
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
    username, password, ok := r.BasicAuth()
    if username == "" || password == "" || !ok {
        http.Error(w, internal.NoCredentialsProvided.Error(), http.StatusUnauthorized)
        return
    }
    login, err := h.service.Login(username, password)
    if err != nil {
        log.Printf("WARNING|user|login|service|%s|%s", username, err.Error())
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    loginResponse := struct {
        JWT string `json:"jwt"`
    } {login}

    response, err := json.Marshal(loginResponse)
    if err != nil {
        errString := fmt.Sprintf("%s|%s", internal.JSONSerialization.Error(), err.Error())
        log.Printf("WARNING|user|login|serialization|%s", errString)
        http.Error(w, errString, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if _, err = w.Write(response); err != nil{
        log.Printf("ERROR|user|login|writeResponse|%s", err.Error())
    }
}
