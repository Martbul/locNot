package api

import (
	"os"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Server struct to hold router, logger, and DB connection
type Server struct {
	Router *mux.Router
	Logger hclog.Logger
}

// NewServer initializes the router, sets up middlewares, and routes
func NewServer(logger hclog.Logger) *Server {
	router := mux.NewRouter()
	server := &Server{
		Router: router,
		Logger: logger,
	}

	// Add CORS middleware
	corsAddress := os.Getenv("CORS_ADDRESS")
	corsMiddleware := gohandlers.CORS(gohandlers.AllowedOrigins([]string{corsAddress}))
	server.Router.Use(corsMiddleware)

	// Register routes and middlewares
	server.setupRoutes()
	return server
}

// setupRoutes registers all API routes
func (s *Server) setupRoutes() {
	//	s.Router.HandleFunc("/hello", s.HelloHandler).Methods("GET")
	//
	// Add other routes here
}
