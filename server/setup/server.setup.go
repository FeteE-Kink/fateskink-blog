package setup

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/app/database"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	shutDownTimeout = 5 * time.Second
)

func InitServer() *http.Server {
	handler := SetupRouter()
	serverAddr := strings.TrimSpace(os.Getenv("PORT"))
	if serverAddr == "" {
		serverAddr = "3007"
	}

	serverAddr = fmt.Sprintf(":%s", serverAddr)

	return &http.Server{
		Addr:    serverAddr,
		Handler: handler,
	}
}

func StartServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
}

func WaitForShutDown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	if err := database.CloseDb(); err != nil {
		log.Fatal("Database connection closure error:", err)
	}

	log.Println("Server exiting")
}
