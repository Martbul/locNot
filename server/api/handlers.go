// server/api/handlers.go
package api

import (
	"encoding/json"
	"net/http"
)

// HelloHandler is a sample endpoint handler
func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello, world!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ProtectedHandler is an example of an authenticated route
func (s *Server) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "This is a protected route!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
