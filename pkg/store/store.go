package store

import (
	"github.com/job25721/go-jwt/pkg/model"
)

var Users []model.User

type Storer interface {
	FindUser(username string) *model.User
	AddUser(user model.User)
}

type store struct{}

func NewStore() Storer {
	return &store{}
}

func (r *store) FindUser(username string) *model.User {
	var user *model.User
	//do a linear scan
	for _, u := range Users {
		if u.Username == username {
			user = &u
			break
		}
	}
	return user
}

func (r *store) AddUser(user model.User) {
	Users = append(Users, user)
}
