package utils

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/garnachod/random_images/internal"
    "log"
    "net/http"
)

var (
    ContextKeyCaller = "caller"
)

func SerializeJSON(ctx context.Context, w http.ResponseWriter, jsonSerializable interface{}) {
    callerCtx, ok := ctx.Value(ContextKeyCaller).(string)
    if !ok {
        callerCtx = "utils|SerializeJSON"
    }

    response, err := json.Marshal(jsonSerializable)
    if err != nil {
        errString := fmt.Sprintf("%s|%s", internal.JSONSerialization.Error(), err.Error())

        log.Printf("WARNING|%s|serialization|%s", callerCtx,  errString)
        http.Error(w, errString, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if _, err = w.Write(response); err != nil{
        log.Printf("ERROR|%s|writeResponse|%s", callerCtx, err.Error())
    }
}
