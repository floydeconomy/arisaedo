package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
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
	router.HandleFunc("/api/ipfs/x/{id}", handler).Methods("GET", "PUT")

	// CORS
	handler := handlers.CompressHandler(router)
	handler = handlers.CORS(
		handlers.AllowedOrigins(origins),
	)(handler)

	// return
	return handler.ServeHTTP, nil
}

