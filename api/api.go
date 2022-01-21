package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/top-ten-words-api/service"
)

type Server struct {
	mux                *mux.Router
	topTenWordsService service.ITopTenWords
}

func InitServer(topTenWordsService service.ITopTenWords) *Server {
	mux := mux.NewRouter()

	return &Server{
		mux:                mux,
		topTenWordsService: topTenWordsService,
	}
}

func (s *Server) Start() error {
	s.SetupRoutes()
	server := http.Server{
		Addr:         ":3000",
		Handler:      s.mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	serverError := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("server is starting at port 3000")
		serverError <- server.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("server error: %v", err.Error())
	case sig := <-shutdown:
		log.Println("starting graceful shutdown of port 3000")
		defer log.Printf("graceful shutdown complete: %s\n", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not stop the server gracefully: %v", err.Error())
		}
	}

	return nil
}

func (s *Server) SetupRoutes() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello world!",
		})
	}).Methods("GET")

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer s.topTenWordsService.Reset()
		var payload map[string]string
		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(s.topTenWordsService.TopTenWords(payload["text"]))
	}).Methods("POST")
}
