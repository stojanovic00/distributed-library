package main

import (
	"central_library/app"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Loads config and instantiates all dependencies
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	//Server startup
	address := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.Port)
	server := &http.Server{
		Handler:           app.Router,
		Addr:              address,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 100 * time.Millisecond,
		MaxHeaderBytes:    2048,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//GRACEFUL SHUTDOWN
	quitChan := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	log.Println("Shutdown Server ...")

	//TODO revert to 2
	//timeoutTime := 2
	timeoutTime := 0.1
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutTime)*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds.", timeoutTime)
	}
	log.Println("Server exiting")
}
