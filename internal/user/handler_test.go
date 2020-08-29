package user

import (
    "github.com/garnachod/random_images/internal"
    "github.com/golang/mock/gomock"
    "net/http/httptest"
    "testing"
)

func Test_handler_Login_admin(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    m.EXPECT().GetUser(gomock.Eq("admin")).Return(NewUserModel("admin", "password"), nil)

    jwtPass := "password"
    s := NewUserService(m, jwtPass)

    handler := NewUserHandler(s)
    req := httptest.NewRequest("GET", "localhost:3000/v0/login", nil)
    req.SetBasicAuth("admin", "password")

    w := httptest.NewRecorder()

    handler.Login(w, req)
    if w.Code != 200 {
        t.Errorf("login error admin")
        return
    }
}

func Test_handler_Login_no_user(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)

    jwtPass := "password"
    s := NewUserService(m, jwtPass)

    handler := NewUserHandler(s)
    req := httptest.NewRequest("GET", "localhost:3000/v0/login", nil)
    req.SetBasicAuth("", "")

    w := httptest.NewRecorder()

    handler.Login(w, req)
    if w.Code != 401 {
        t.Errorf("login error admin")
        return
    }
}

func Test_handler_Login_no_valid_user(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    m.EXPECT().GetUser(gomock.Eq("admin")).Return(nil, internal.NotFound)

    jwtPass := "password"
    s := NewUserService(m, jwtPass)

    handler := NewUserHandler(s)
    req := httptest.NewRequest("GET", "localhost:3000/v0/login", nil)
    req.SetBasicAuth("admin", "password")

    w := httptest.NewRecorder()

    handler.Login(w, req)
    if w.Code != 401 {
        t.Errorf("login error no_valid_user")
        return
    }
}