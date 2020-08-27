package unsplash

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/garnachod/random_images/internal/image"
    "net/http"
)

type imageProvider struct {
    token string
}


func NewUnsplashImageProvider(token string) image.Provider {
    return &imageProvider{
        token: token,
    }
}

func (imgP imageProvider) GetImage(x, y int) (*image.Model, error) {
    req, err := http.NewRequest("GET", "https://api.unsplash.com/photos/random", nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Accept-Version", "v1")
    req.Header.Add("Authorization", "Client-ID " + imgP.token)
    q := req.URL.Query()
    q.Add("width", fmt.Sprint(x))
    q.Add("height", fmt.Sprint(y))
    req.URL.RawQuery = q.Encode()
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }

    target := map[string]interface{}{}
    err = json.NewDecoder(res.Body).Decode(&target)
    if err != nil {
        return nil, err
    }

    errorsApi, ok := target["errors"].([]interface{})
    if !ok {
        urls, ok := target["urls"].(map[string]interface{})
        if !ok {
            return nil, errors.New("invalid unsplash urls")
        }
        raw := urls["raw"].(string)
        return image.NewImageModel(raw), nil
    } else {
        if len(errorsApi) > 0 {
            return nil, fmt.Errorf("%v", errorsApi)
        }
    }

   return nil, errors.New("invalid unsplash response")

}