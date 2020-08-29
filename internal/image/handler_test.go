package image

import (
    "errors"
    "github.com/golang/mock/gomock"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_handler_GetImage(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    m.EXPECT().GetImage(gomock.Eq(123),gomock.Eq(123)).Return(NewImageModel("http://url.com"), nil)
    s := NewImagesService(m)
    h := NewImageHandler(s)

    req := httptest.NewRequest("GET", "localhost:3000/v0/images/random?x=123&y=123", nil)

    w := httptest.NewRecorder()

    h.GetImage(w, req)
    if w.Code != 200 {
        t.Errorf("GetImage unexpected error")
        return
    }
}

func Test_handler_GetImage_x_not_provided(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    s := NewImagesService(m)
    h := NewImageHandler(s)

    req := httptest.NewRequest("GET", "localhost:3000/v0/images/random?y=123", nil)

    w := httptest.NewRecorder()

    h.GetImage(w, req)
    if w.Code != http.StatusBadRequest {
        t.Errorf("GetImage expect error x not provided")
        return
    }
}

func Test_handler_GetImage_y_not_provided(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    s := NewImagesService(m)
    h := NewImageHandler(s)

    req := httptest.NewRequest("GET", "localhost:3000/v0/images/random?x=123", nil)

    w := httptest.NewRecorder()

    h.GetImage(w, req)
    if w.Code != http.StatusBadRequest {
        t.Errorf("GetImage expect error y not provided")
        return
    }
}

func Test_handler_GetImage_internal_server(t *testing.T) {
    ctrl := gomock.NewController(t)

    // Assert that Bar() is invoked.
    defer ctrl.Finish()

    m := NewMockProvider(ctrl)
    m.EXPECT().GetImage(gomock.Eq(123),gomock.Eq(123)).Return(nil, errors.New("invalid unsplash response"))
    s := NewImagesService(m)
    h := NewImageHandler(s)

    req := httptest.NewRequest("GET", "localhost:3000/v0/images/random?x=123&y=123", nil)

    w := httptest.NewRecorder()

    h.GetImage(w, req)
    if w.Code != http.StatusInternalServerError {
        t.Errorf("GetImage error from provider")
        return
    }
}