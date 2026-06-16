package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/404bad/students-api/internal/config"
)


func main(){
	
	// load configuration
	cfg := config.MustLoadConfig()

	//database connection

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to student Api"))
	})

	//setup http server
	server := &http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,	
	}

	slog.Info("Server starting at http://" + cfg.HTTPServer.Address)
	

	// we create a channel to listen for shutdown signal, and we will use this channel to gracefully shutdown the server when we receive a shutdown signal. This is important because it allows us to finish processing any ongoing requests before shutting down the server, which can help prevent data loss and ensure a better user experience.

	done := make(chan os.Signal, 1)
	
	// we will listen for shutdown signal in a separate go routine, and when we receive a shutdown signal, we will close the done channel, which will signal the main go routine to shutdown the server gracefully.

	signal.Notify(done, os.Interrupt, syscall.SIGTERM,syscall.SIGINT) // we will listen for os.Interrupt and os.Kill signals, which are the most common signals used to shutdown a server. os.Interrupt is sent when the user presses Ctrl+C in the terminal, and os.Kill is sent when the process is killed by the operating system.

	// we have to imolement graceful shutdown of the server, because if we just call server.ListenAndServe() and the server is running, if we want to stop the server, we will have to kill the process, which is not a good way to stop the server. We should implement graceful shutdown, which will allow us to stop the server gracefully, by allowing it to finish processing any ongoing requests before shutting down.
	// go routine to listen for shutdown signal
	go func(){
		// listen for shutdown signal
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// wait for shutdown signal
	<-done
	// until we receive a shutdown signal, the main go routine will be blocked on the done channel, and when we receive a shutdown signal, we will close the done channel, which will signal the main go routine to shutdown the server gracefully.
	
	slog.Info("Shutting down the server")

	// we will use the Shutdown method of the http.Server struct to shutdown the server gracefully. The Shutdown method takes a context as an argument, which we can use to set a timeout for the shutdown process. This is important because if we have any ongoing requests that are taking a long time to complete, we don't want to wait indefinitely for them to finish before shutting down the server.

	// we will use a context with a timeout of 5 seconds for the shutdown process, which means that if there are any ongoing requests that are taking longer than 5 seconds to complete, the server will forcefully shutdown after the timeout expires.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // we will defer the cancel function to ensure that the context is properly cleaned up after the shutdown process is complete.	

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown the server gracefully",slog.String("error", err.Error())) // we will log an error if the shutdown process fails, which can help us identify any issues that may arise during the shutdown process. We will use slog to log the error, and we will include the error message in the log entry for better visibility.
	}

	slog.Info("Server stopped successfully") // we will log a message indicating that the server has stopped successfully, which can help us confirm that the shutdown process was completed without any issues.

}