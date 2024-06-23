package util

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

type server struct {
	server *http.Server
	port   string
}

func Server(port string) *server {
	return &server{}
}

func (s *server) Start(cb func(m *http.ServeMux)) {
	mux := http.NewServeMux()

	cb(mux)

	s.server = &http.Server{
		Addr:    s.port,
		Handler: mux,
	}

	fmt.Printf("Starting server on %v\n", s.port)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not start server: %s\n", err)
	}
}

func (s *server) Stop() {
	if s.server != nil {
		fmt.Println("Stopping server...")
		if err := s.server.Close(); err != nil {
			fmt.Printf("Error stopping server: %s\n", err)
		}
	}
}

func (s *server) MonitorInput(cb func(m *http.ServeMux)) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error entering raw mode: %v\n", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Press 'r' to restart the server...")
	fmt.Println("Press 'q' to quit the server...")
	byteSlice := make([]byte, 1)

	for {
		_, err := os.Stdin.Read(byteSlice)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}

		if byteSlice[0] == 'r' {
			fmt.Println("Restart command received")
			s.Stop()
			time.Sleep(1 * time.Second) // Brief pause to ensure the server stops
			go s.Start(cb)
		}

		if byteSlice[0] == 'q' {
			os.Exit(1)
		}
	}
}

func (s *server) Server(cb func(m *http.ServeMux)) {

	go s.Start(cb)

	// Monitor for keyboard interrupt (Ctrl+C) to gracefully shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.Stop()
		os.Exit(0)
	}()

	s.MonitorInput(cb)
}
