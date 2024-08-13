package core

import (
	"api/config"
	"testing"
)

func TestWheelRNG(t *testing.T) {
	for i := 0; i < 10000; i++ {
		number, err := GenerateWheelNumber()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if number < 0 {
			t.Log("random number lower than expected:", number)
			t.FailNow()
		}
		if number > 11 {
			t.Log("random number greater than expected:", number)
			t.FailNow()
		}
	}
}

func TestPrizeRNG(t *testing.T) {
	config.PRIZE_SCORE_MAX = 1000
	config.PRIZE_SCORE_MIN = 100

	for i := 0; i < 10000; i++ {
		number, err := GeneratePrize()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if number < config.PRIZE_SCORE_MIN {
			t.Log("random number lower than expected:", number)
			t.FailNow()
		}
		if number > config.PRIZE_SCORE_MAX {
			t.Log("random number greater than expected:", number)
			t.FailNow()
		}
	}
}
