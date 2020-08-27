package image

type Model struct {
    url string
}

func NewImageModel(url string) *Model {
    return &Model{url: url}
}

