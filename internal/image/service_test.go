package image

import (
    "errors"
    "github.com/golang/mock/gomock"
    "testing"
)

func Test_service_GetImageUrl(t *testing.T) {
    type args struct {
        x int
        y int
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {name: "found image", args: args{123, 123}, want: "http://url.com", wantErr: false},
        {name: "not found image", args: args{123, 123}, want: "",  wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)

            // Assert that Bar() is invoked.
            defer ctrl.Finish()

            m := NewMockProvider(ctrl)
            if !tt.wantErr {
                m.EXPECT().GetImage(gomock.Eq(tt.args.x),gomock.Eq(tt.args.y)).Return(NewImageModel("http://url.com"), nil)
            } else {
                m.EXPECT().GetImage(gomock.Eq(tt.args.x),gomock.Eq(tt.args.y)).Return(nil, errors.New("invalid unsplash response"))
            }

            s := NewImagesService(m)
            got, err := s.GetImageUrl(tt.args.x, tt.args.y)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetImageUrl() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("GetImageUrl() got = %v, want %v", got, tt.want)
            }
        })
    }
}
