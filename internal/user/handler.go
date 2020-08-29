package user

import (
    "context"
    "github.com/garnachod/random_images/internal"
    "github.com/garnachod/random_images/internal/utils"
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

    ctx := context.WithValue(context.Background(), utils.ContextKeyCaller, "user|login")
    utils.SerializeJSON(ctx, w, loginResponse)
}
