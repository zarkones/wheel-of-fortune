package storage

import (
	"api/config"
	"sync"
)

var users = map[string]User{}
var userMux = &sync.Mutex{}

type User struct {
	Name    string
	Surname string
	Email   string
	Credits int
}

func InsertUser(user *User) (err error) {
	if user == nil {
		return ErrUserInvalid
	}

	userMux.Lock()
	defer userMux.Unlock()

	if _, ok := users[user.Email]; ok {
		return ErrUserExists
	}

	user.Credits = config.STARTER_SPIN_CREDITS

	users[user.Email] = *user

	return nil
}

func GetUser(email string) (user User, err error) {
	user, ok := users[email]
	if !ok {
		return user, ErrUserNotFound
	}
	return user, nil
}
