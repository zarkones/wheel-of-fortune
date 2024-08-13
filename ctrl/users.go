package ctrl

import (
	"api/contracts"
	"api/storage"
	"encoding/json"
	"errors"
	"net/http"
)

// Register registers a user.
func Register(w http.ResponseWriter, r *http.Request) {
	var ctx contracts.RegisterReqCtx

	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := storage.InsertUser(&storage.User{
		Name:    ctx.Name,
		Surname: ctx.Surname,
		Email:   ctx.Email,
	}); err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
