package image

import (
    "context"
    "fmt"
    "github.com/garnachod/random_images/internal"
    "github.com/garnachod/random_images/internal/utils"
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

// GetImage returns a json with a url of a image
// needs 2 http query params x = with of the image and y = height of the image
// response {"url": "<image_url>"}
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

    ctx := context.WithValue(context.Background(), utils.ContextKeyCaller, "image|getImage")
    utils.SerializeJSON(ctx, w, imageResponse)
}
