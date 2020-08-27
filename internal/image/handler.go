package image

import (
    "encoding/json"
    "fmt"
    "github.com/garnachod/random_images/internal"
    "log"
    "net/http"
    "strconv"
)

type Handler interface{
    GetImage(w http.ResponseWriter, r *http.Request)
}

type handler struct {
    service Service
}

func NewImageHandler(service Service) Handler {
    return &handler{
        service,
    }
}

func (h handler) GetImage(w http.ResponseWriter, r *http.Request) {
    xStr := r.FormValue("x")
    yStr := r.FormValue("y")
    x, err := strconv.Atoi(xStr)
    if err != nil {
        errString := fmt.Sprintf("%s|%s", internal.UrlParams.Error(), "x param is not a valid integer")
        http.Error(w, errString, http.StatusBadRequest)
        return
    }
    y, err := strconv.Atoi(yStr)
    if err != nil {
        errString := fmt.Sprintf("%s|%s", internal.UrlParams.Error(), "y param is not a valid integer")
        http.Error(w, errString, http.StatusBadRequest)
        return
    }
    url, err := h.service.GetImageUrl(x, y)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    imageResponse := struct {
        URL string `json:"url"`
    } {url}

    response, err := json.Marshal(imageResponse)
    if err != nil {
        errString := fmt.Sprintf("%s|%s", internal.JSONSerialization.Error(), err.Error())
        log.Printf("WARNING|image|get|serialization|%s", errString)
        http.Error(w, errString, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if _, err = w.Write(response); err != nil{
        log.Printf("ERROR|image|get|writeResponse|%s", err.Error())
    }
}
