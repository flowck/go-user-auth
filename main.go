package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"user-auth/infra"
	"user-auth/modules/user"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) []byte {
	response, err := json.Marshal(MessageResponse{Message: message})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return response
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var signUpDto user.SignupUserDto
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewMessageResponse("Internal Server Error"))
		return
	}

	err = json.Unmarshal(body, &signUpDto)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewMessageResponse("Please provide a valid JSON"))
		return
	}

	err = user.Signup(ctx, signUpDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// "Unable to signup user"
		w.Write(NewMessageResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(NewMessageResponse("User create successfully"))
}

func GlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Service will start here")

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
		r.Post("/signup", signUpHandler)
	})

	fmt.Println("Server is running at port", infra.Configs.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(int(infra.Configs.Port))), r)

	if err != nil {
		panic(err)
	}
}
