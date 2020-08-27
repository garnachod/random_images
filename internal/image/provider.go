package image

type Provider interface {
    GetImage(x, y int) (*Model, error)
}
