package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/quenyu/pr-reviewer/internal/config"
)

type Server struct {
	cfg        config.Config
	router     chi.Router
	httpServer *http.Server
	errChan    chan error
}

func New(cfg config.Config) *Server {
	r := chi.NewRouter()

	s := &Server{
		cfg:     cfg,
		router:  r,
		errChan: make(chan error, 1),
	}

	s.registerRoutes()

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return s
}

func (s *Server) Start() error {
	if s.httpServer == nil {
		return errors.New("http server is not initialized")
	}

	go func() {
		log.Printf("HTTP server listening on %s", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errChan <- err
		}
	}()

	select {
	case err := <-s.errChan:
		return err
	default:
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) registerRoutes() {
	s.router.Get("/health", s.handleHealth)
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
