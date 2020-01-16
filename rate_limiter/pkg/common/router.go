package common

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Start starts listener
func Start(listenAddress string) {
	r := GetRouter()
	addRoutes(r)

	go func() {
		err := http.ListenAndServe(listenAddress, r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Listening on port:", listenAddress)
	}()
}

// addRoutes adds routes
func addRoutes(router *mux.Router) {
	h := NewHandlers()
	rl := NewRateLimitingHandler()

	router.Handle("/ratelimiter", rl.RateLimitHandle(http.HandlerFunc(h.GetHandler))).Methods("GET")
}

// GetRouter gets router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
