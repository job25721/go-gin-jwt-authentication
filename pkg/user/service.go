package user

import (
	"errors"
	"net/http"

	"github.com/job25721/go-jwt/pkg/bcrypt"
	"github.com/job25721/go-jwt/pkg/jwt"
	"github.com/job25721/go-jwt/pkg/model"
	"github.com/job25721/go-jwt/pkg/store"
)

type IUserService interface {
	Register(req model.User) (string, int, error)
	Login(req model.User) (string, int, error)
}

type service struct {
	store store.Storer
}

func NewService(store store.Storer) IUserService {
	return &service{
		store: store,
	}
}

func (s *service) Register(req model.User) (string, int, error) {
	if u := s.store.FindUser(req.Username); u != nil {
		return "", http.StatusConflict, errors.New("user exist")
	}

	hashPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	s.store.AddUser(model.User{Username: req.Username, Password: hashPassword})
	return "register success", http.StatusOK, nil
}

func (s *service) Login(req model.User) (string, int, error) {
	user := s.store.FindUser(req.Username)
	if user == nil {
		return "", http.StatusNotFound, errors.New("user not found")
	}

	if ok := bcrypt.CheckPasswordHash(req.Password, user.Password); !ok {
		return "", http.StatusUnauthorized, errors.New("Unauthorized")
	}

	token, _ := jwt.GenerateJWT(user.Username)
	return token, http.StatusOK, nil
}
