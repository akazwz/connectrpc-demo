package main

import (
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"opendrive/auth"
	"opendrive/gen/auth/v1/authv1connect"
	"opendrive/middlewares"
)

func main() {
	fmt.Println("starting server on :8080")
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	interceptors := connect.WithInterceptors(middlewares.NewAuthInterceptor())
	authPath, authHandler := authv1connect.NewPublicAuthServiceHandler(auth.NewAuthService(), interceptors)
	r.Mount(authPath, authHandler)
	log.Fatal(http.ListenAndServe(":8080", h2c.NewHandler(r, &http2.Server{})))
}
