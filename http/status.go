// Package http implements handlers used by rebasebot http server
package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Status(w http.ResponseWriter, r *http.Request) {
	event := strings.ToLower(r.Method)

	log.Printf("http.request.%s.received: %s\n", event, r.RequestURI)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")

	log.Printf("http.%s.response.sent: %d\n", event, http.StatusOK)
}
