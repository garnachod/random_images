package user

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/garnachod/random_images/internal"
    "time"
)

type Service interface {
    Login(username, password string) (string, error)
}

type service struct {
    provider Provider
    key      []byte
}

func NewUserService(prov Provider, key string) Service {
    return &service{
        prov,
        []byte(key),
    }
}

func (s service) Login(username, password string) (token string, err error) {
    user, err := s.provider.GetUser(username)
    if err != nil {
        return "", err
    }
    if err = s.isValidPassword(password, user.password); err != nil {
        return "", err
    }
    token, err = s.getToken(username)
    return
}

/**
    check password
    v0: plain
    v1: hashed
 */
func (s service) isValidPassword(passwordIn, passwordPersist string) error {
    if passwordIn == passwordPersist {
        return nil
    }
    return internal.InvalidPassword
}

func (s service) getToken(username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    /* Create a map to store our claims */
    claims := token.Claims.(jwt.MapClaims)

    /* Set token payload */
    claims["sub"] = username
    claims["type"] = "user"
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    /* Sign the token with our secret */
    return token.SignedString(s.key)
}