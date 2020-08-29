package main

import (
    "fmt"
    "github.com/garnachod/random_images/internal/image"
    "github.com/garnachod/random_images/internal/provider/memory"
    "github.com/garnachod/random_images/internal/provider/unsplash"
    "github.com/garnachod/random_images/internal/user"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/docgen"
    "github.com/go-chi/jwtauth"
    "log"
    "net/http"
    "os"
)

func main() {
    jwtPass := "password"
    tokenAuth := jwtauth.New("HS256", []byte(jwtPass), nil)

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    userProvider := memory.NewMemoryUserProvider()
    userService := user.NewUserService(userProvider, jwtPass)
    userHandler := user.NewUserHandler(userService)


    imageProvider := unsplash.NewUnsplashImageProvider(os.Getenv("unsplash"))
    imageService := image.NewImagesService(imageProvider)
    imageHandler := image.NewImageHandler(imageService)


    r.Route("/v0", func(r chi.Router) {
        r.Route("/images", func(r chi.Router) {
            r.Use(jwtauth.Verifier(tokenAuth))
            r.Use(jwtauth.Authenticator)
            r.Get("/random", imageHandler.GetImage)
        })
        r.Get("/login", userHandler.Login)
    })

    docu := os.Getenv("docu")
    if docu == "true"{
        os.Setenv("GOPATH", "/src/")
        f, _ := os.Create("/src/autogen/api.md")
        fmt.Fprint(f, docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
            ProjectPath: "github.com/garnachod/random_images",
            Intro:       "Welcome to the random_images api generated docs.",
        }))
        return
    }

    err := http.ListenAndServe(":3000", r)
    if err != nil {
        log.Print(err.Error())
        return
    }
}