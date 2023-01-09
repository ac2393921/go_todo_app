package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ac2393921/go_todo_app/internal/config"
	"github.com/ac2393921/go_todo_app/internal/mux"
	"github.com/ac2393921/go_todo_app/internal/server"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mx := mux.NewMux()
	s := server.NewServer(l, mx)
	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
