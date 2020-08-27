package main

import (
    "github.com/garnachod/random_images/internal/provider/memory"
    "github.com/garnachod/random_images/internal/provider/unsplash"
    "github.com/garnachod/random_images/internal/user"
    "github.com/garnachod/random_images/internal/image"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/jwtauth"
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

    http.ListenAndServe(":3000", r)
}