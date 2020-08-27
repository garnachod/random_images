package user

type Provider interface {
    GetUser(username string) (*Model, error)
}
