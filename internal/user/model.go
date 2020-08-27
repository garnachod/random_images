package user

type Model struct {
    username string
    password string
}

func NewUserModel(username, password string) *Model {
    return &Model{username: username, password: password}
}