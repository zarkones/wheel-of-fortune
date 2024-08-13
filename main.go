package main

import (
	"api/config"
	"api/middleware"
	"log"
	"net"
	"net/http"
)

// maybeFatal would panic if an error is not nil.
// The function is located in the main file,
// as other segments of the code should return an error,
// rather than crashing the program.
func maybeFatal(err error) {
	if err == nil {
		return
	}
	log.Panicln(err)
}

func main() {
	maybeFatal(config.VerifyEnv())

	mux := http.NewServeMux()

	initRoutes(mux)

	apiServer := http.Server{
		Addr:    net.JoinHostPort(config.HOST, config.PORT_API),
		Handler: middleware.Use(mux, middleware.CORS),
	}

	staticServer := http.Server{
		Addr:    net.JoinHostPort(config.HOST, config.PORT_STATIC_SERVE),
		Handler: http.FileServer(http.Dir(config.PATH_STATIC_FILES)),
	}

	log.Println("Serving API at:", apiServer.Addr)
	log.Println("Serving static files at:", staticServer.Addr)

	go func() {
		maybeFatal(apiServer.ListenAndServe())
	}()

	staticServer.ListenAndServe()
}
