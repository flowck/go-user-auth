package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	err := infra.RunMigrations(infra.DB, &infra.Configs)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at port", infra.Configs.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(int(infra.Configs.Port))), nil)

	if err != nil {
		panic(err)
	}
}
