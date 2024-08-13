package config

import (
	"net"
	"os"
	"strconv"
)

// VerifyEnv assures that all environment variables have been set correctly.
func VerifyEnv() (err error) {
	if len(HOST) == 0 {
		HOST = "127.0.0.1"
	}
	if len(PORT_STATIC_SERVE) == 0 {
		PORT_STATIC_SERVE = "8080"
	}
	if len(PORT_API) == 0 {
		PORT_API = "8081"
	}

	if len(os.Getenv("PRIZE_SCORE_MAX")) == 0 {
		PRIZE_SCORE_MAX = 1000
	} else {
		PRIZE_SCORE_MAX, err = strconv.Atoi(os.Getenv("PRIZE_SCORE_MAX"))
		if err != nil {
			return err
		}
		if PRIZE_SCORE_MAX == 0 {
			return ErrInvalidMaxPrize
		}
	}

	if len(os.Getenv("PRIZE_SCORE_MIN")) == 0 {
		PRIZE_SCORE_MIN = 100
	} else {
		PRIZE_SCORE_MIN, err = strconv.Atoi(os.Getenv("PRIZE_SCORE_MIN"))
		if err != nil {
			return err
		}
		if PRIZE_SCORE_MIN == 0 {
			return ErrInvalidMinPrize
		}
	}

	if len(os.Getenv("STARTER_SPIN_CREDITS")) == 0 {
		STARTER_SPIN_CREDITS = 2
	} else {
		STARTER_SPIN_CREDITS, err = strconv.Atoi(os.Getenv("STARTER_SPIN_CREDITS"))
		if err != nil {
			return err
		}
	}

	if len(PATH_STATIC_FILES) == 0 {
		PATH_STATIC_FILES = "static"
	}

	if len(ALLOWED_ORIGIN) == 0 {
		ALLOWED_ORIGIN = "http://" + net.JoinHostPort(HOST, PORT_STATIC_SERVE)
	}

	return nil
}

var (
	HOST              = os.Getenv("HOST")
	PORT_STATIC_SERVE = os.Getenv("PORT_STATIC_SERVE")
	PORT_API          = os.Getenv("PORT_API")

	PATH_STATIC_FILES = os.Getenv("PATH_STATIC_FILES")

	ALLOWED_ORIGIN = os.Getenv("ALLOWED_ORIGIN")

	PRIZE_SCORE_MAX = 0
	PRIZE_SCORE_MIN = 0

	STARTER_SPIN_CREDITS = 0
)
