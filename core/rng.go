package core

import (
	"api/config"
	"crypto/rand"
	"math/big"
)

// GenerateWheelNumber generates a cryptographically secure
// random number used to determine the winning slot on the wheel.
func GenerateWheelNumber() (number int, err error) {
	secureInteger, err := rand.Int(rand.Reader, big.NewInt(12))
	return int(secureInteger.Int64()), err
}

// IsWinningNummber determines if a number is a winning number.
func IsWinningNummber(number int) bool {
	switch number {
	case 0, 3, 7, 10:
		return true

	default:
		return false
	}
}

// GeneratePrize would generate a random prize in terms of the score.
func GeneratePrize() (prizeScore int, err error) {
	secureInteger, err := rand.Int(rand.Reader, big.NewInt(int64(config.PRIZE_SCORE_MAX-config.PRIZE_SCORE_MIN)))
	return int(secureInteger.Int64()) + config.PRIZE_SCORE_MIN, err
}
