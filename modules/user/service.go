package user

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	token string
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

// func SignIn(user SigninUserDto) (AuthUser, error) {}
