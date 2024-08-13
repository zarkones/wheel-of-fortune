package ctrl

import (
	"api/contracts"
	"api/core"
	"api/storage"
	"encoding/json"
	"errors"
	"net/http"
)

func GetPrize(w http.ResponseWriter, r *http.Request) {
	spinID := r.PathValue("spinID")

	spin, err := storage.GetSpin(spinID)
	if err != nil {
		if errors.Is(err, storage.ErrSpinNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&spin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Spin return a winning number on the wheel of fortune.
func Spin(w http.ResponseWriter, r *http.Request) {
	userEmail := r.PathValue("userEmail")

	if _, err := storage.GetUser(userEmail); err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wheelInt, err := core.GenerateWheelNumber()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	win := core.IsWinningNummber(wheelInt)

	prize := 0
	if win {
		prize, err = core.GeneratePrize()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	spinID, err := storage.InsertSpin(&storage.Spin{
		Number:    wheelInt,
		Win:       win,
		Prize:     prize,
		SpinnedBy: userEmail,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respCtx := contracts.SpinRespCtx{
		SpinID: spinID,
		Number: wheelInt,
		Win:    win,
	}

	if err := json.NewEncoder(w).Encode(&respCtx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
