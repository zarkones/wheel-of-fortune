package storage

import (
	"api/contracts"
	"sync"

	"github.com/google/uuid"
)

var spins = map[string]Spin{}
var spinMux = &sync.Mutex{}

type Spin struct {
	ID        string
	SpinnedBy string // Email address of the user.
	Win       bool
	Number    int
	Prize     int
}

func InsertSpin(spin *Spin) (spinID string, err error) {
	if spin == nil {
		return "", ErrSpinInvalid
	}

	spinMux.Lock()
	defer spinMux.Unlock()

	spin.ID = uuid.NewString()

	user, ok := users[spin.SpinnedBy]
	if !ok {
		return "", ErrUserNotFound
	}
	if user.Credits <= 0 {
		return "", contracts.ErrNotEnoughCredits
	}
	user.Credits -= 1
	users[spin.SpinnedBy] = user

	spins[spin.ID] = *spin

	return spin.ID, nil
}

func GetSpin(id string) (spin Spin, err error) {
	spin, ok := spins[id]
	if !ok {
		return spin, ErrSpinNotFound
	}
	return spin, nil
}
