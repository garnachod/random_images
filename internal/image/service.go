package image

type Service interface {
    GetImageUrl(x, y int) (string, error)
}

type service struct {
    provider Provider
}

func NewImagesService(prov Provider) Service {
    return &service{
        prov,
    }
}

func (s service) GetImageUrl(x, y int) (string, error) {
    image, err := s.provider.GetImage(x, y)
    if err != nil {
        return "", err
    }
    return image.url, nil
}