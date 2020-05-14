package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/floydeconomy/arisaedo-go/api/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

// todo: fix func()
func New(
	allowedOrigins string,
) (http.HandlerFunc, func()) {
	// split allowed origin
	origins := strings.Split(strings.TrimSpace(allowedOrigins), ",")
	for i, o := range origins {
		origins[i] = strings.ToLower(strings.TrimSpace(o))
	}

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/ipfs/x/{id}", utils.ErrorHandler(handler)).Methods("GET", "PUT")

	// CORS
	handler := handlers.CompressHandler(router)
	handler = handlers.CORS(
		handlers.AllowedOrigins(origins),
	)(handler)

	// return
	return handler.ServeHTTP, nil
}
