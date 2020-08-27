package memory

import (
    "github.com/garnachod/random_images/internal"
    "github.com/garnachod/random_images/internal/user"
)

type userProvider struct {

}

func NewMemoryUserProvider() user.Provider {
    return &userProvider{}
}

func (u userProvider) GetUser(username string) (*user.Model, error) {
    if username == "admin" {
        return user.NewUserModel("admin", "password"), nil
    }
    return nil, internal.NotFound
}
