package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func Start() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/api/ipfs/covid/{id}", handler).Methods("GET", "PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// todo: fix func()
func New(
	allowedOrigins string,
) (http.HandlerFunc, func()) {
	// Split Allowed Origin
	origins := strings.Split(strings.TrimSpace(allowedOrigins), ",")
	for i, o := range origins {
		origins[i] = strings.ToLower(strings.TrimSpace(o))
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/ipfs/covid/{id}", handler).Methods("GET", "PUT")
	//log.Fatal(http.ListenAndServe(":8080", r))

	handler := handlers.CompressHandler(router)
	handler = handlers.CORS(
		handlers.AllowedOrigins(origins),
	)(handler)
	return handler.ServeHTTP, nil
}
