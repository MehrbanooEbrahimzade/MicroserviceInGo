package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/cmd"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	factory := cmd.NewFactory()

	s := http.Server{
		Addr:    factory.Port,
		Handler: factory.Handler,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println(err)
			}
			fmt.Println("Server closed !")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	s.Shutdown(shutdownCtx)
}
