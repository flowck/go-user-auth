package user

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"user-auth/infra"
	"user-auth/models"
)

type SignupUserDto struct {
	email    string `json:"email"`
	password string `json:"password"`
}

type SigninUserDto struct {
	email    string `json:"email"`
	password string `json:"password"`
}

type AuthUser struct {
	token string `json:"token"`
}

func Signup(ctx context.Context, user SignupUserDto) error {

	// User exists?
	exists, _ := models.Users(qm.Where("email = ? ", user.email)).One(ctx, infra.DB)

	if exists != nil {
		return errors.New("User already exists.")
	}

	// Hash user's password
	password, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
		return err
	}

	newUser := models.User{
		Email:    user.email,
		Password: string(password),
	}

	err = newUser.Insert(ctx, infra.DB, boil.Infer())

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func SignIn(ctx context.Context, user SigninUserDto) (AuthUser, error) {

	fmt.Println("User email", user.email)

	if user.email == "" || user.password == "" {
		return AuthUser{}, errors.New("Please provide valid details.")
	}

	// find user
	foundUser, err := models.Users(qm.Where("email = ?", user.email)).One(ctx, infra.DB)

	if err != nil {
		log.Println(err)
		return AuthUser{}, errors.New("Unable to perform this operation")
	}

	fmt.Println("User", foundUser)

	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.password))

	if err != nil {
		return AuthUser{}, errors.New("Please provide the valid password for this account.")
	}

	// Generate JWT token
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(0, 0)),
		Issuer:    "USER_AUHT",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var token string
	token, err = jwtToken.SignedString(infra.Configs.JwtSigningKey)

	return AuthUser{token: token}, nil
}
