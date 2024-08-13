package storage

import (
	"api/config"
	"api/core"
	"log"
	"testing"
)

func TestStorage(t *testing.T) {
	config.PRIZE_SCORE_MAX = 100
	config.PRIZE_SCORE_MIN = 10

	if err := config.VerifyEnv(); err != nil {
		log.Println(err)
		t.FailNow()
	}

	user := User{
		Name:    "Test",
		Surname: "QA",
		Email:   "test@test.com",
	}

	if err := InsertUser(&user); err != nil {
		log.Println(err)
		t.FailNow()
	}

	wheelInt, err := core.GenerateWheelNumber()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	win := core.IsWinningNummber(wheelInt)

	prize := 0
	if win {
		prize, err = core.GeneratePrize()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}

	spinID, err := InsertSpin(&Spin{
		Number:    wheelInt,
		Win:       win,
		Prize:     prize,
		SpinnedBy: user.Email,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	prevCredits := user.Credits

	user, err = GetUser(user.Email)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if prevCredits == user.Credits {
		t.Log("credits not subtracted")
		t.FailNow()
	}

	if prevCredits-1 != user.Credits {
		t.Log("credits miscalculated")
		t.FailNow()
	}

	spin, err := GetSpin(spinID)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if spin.Win {
		t.Log("You spinned", spin.Number, "and have won", spin.Prize)
	} else {
		t.Log("You spinned", spin.Number, "and lost :(")
	}
}
