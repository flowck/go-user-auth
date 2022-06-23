package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"user-auth/internal/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var signUpDto SignupUserDto
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.NewMessageResponse("Internal Server Error"))
		panic(err)
	}

	err = json.Unmarshal(body, &signUpDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.NewMessageResponse("Please provide a valid JSON"))
		panic(err)
	}

	err = Signup(ctx, signUpDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// "Unable to signup user"
		w.Write(utils.NewMessageResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(utils.NewMessageResponse("User create successfully"))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInDto SigninUserDto

	body, isReadAllOk := ioutil.ReadAll(r.Body)

	if isReadAllOk != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.NewMessageResponse("Internal Server Error"))
		return
	}

	isUnMarshalOk := json.Unmarshal(body, &signInDto)

	if isUnMarshalOk != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.NewMessageResponse("Please provide a valid JSON"))
		return
	}

	auth, isSignInOk := SignIn(r.Context(), signInDto)

	if isSignInOk != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(utils.NewMessageResponse(isSignInOk.Error()))
		return
	}

	response, isMarshal := json.Marshal(auth)

	if isMarshal != nil {
		log.Println(isMarshal)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.NewMessageResponse("Internal Server Error"))
		return
	}

	w.Write(response)
}
