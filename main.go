package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"user-auth/infra"
	"user-auth/modules/user"
)

func GlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	dbErr := infra.InitDatabase(infra.Configs)

	if dbErr != nil {
		panic(dbErr)
	}

	err := infra.RunMigrations(infra.DB, infra.Configs)

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(GlobalMiddleware)
	r.Use(middleware.Logger)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", user.SignUpHandler)
		r.Post("/signin", user.SignInHandler)
	})

	fmt.Println("Server is running at port", infra.Configs.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(int(infra.Configs.Port))), r)

	if err != nil {
		panic(err)
	}
}
