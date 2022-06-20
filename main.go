package main

import (
	"fmt"
	"net/http"
	"user-auth/infra"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Service will start here")

	dbErr := infra.InitDatabase(&infra.Configs)

	if dbErr != nil {
		panic(dbErr)
	}
}
