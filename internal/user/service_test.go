package user

import (
    "github.com/garnachod/random_images/internal"
    "github.com/go-chi/jwtauth"
    "github.com/golang/mock/gomock"
    "testing"
)

func Test_service_Login(t *testing.T) {
    type args struct {
        username string
        password string
    }
    tests := []struct {
        name      string
        args      args
        wantErr   bool
    }{
        {name: "admin return correct hash", args: args{"admin", "password"}, wantErr: false},
        {name: "not found user", args: args{"admin", "password"}, wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)

            // Assert that Bar() is invoked.
            defer ctrl.Finish()

            m := NewMockProvider(ctrl)
            if !tt.wantErr {
                m.EXPECT().GetUser(gomock.Eq(tt.args.username)).Return(NewUserModel(tt.args.username, "password"), nil)
            } else {
                m.EXPECT().GetUser(gomock.Eq(tt.args.username)).Return(nil, internal.NotFound)
            }
            jwtPass := "password"
            s := NewUserService(m, jwtPass)

            gotToken, err := s.Login(tt.args.username, tt.args.password)
            if (err != nil) != tt.wantErr {
                t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr {
                tokenAuth := jwtauth.New("HS256", []byte(jwtPass), nil)
                _, err = tokenAuth.Decode(gotToken)
                if err != nil {
                    t.Errorf("Login() error, invalid token: %s", err)
                    return
                }
            }
        })
    }
}