package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "sync"
	"github.com/brijesh025/students-api/internal/config"
	students "github.com/brijesh025/students-api/internal/http/handlers/student"
	"github.com/joho/godotenv"
)
func main(){
	//1. load config
	e := godotenv.Load();
	if(e!=nil){
		log.Println(".env file was not loaded")
	}
	cfg := config.MustLoad();
	//2. database setup
	//3. setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /students/api", students.Welcome())
	//4. setup server
	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server started and Ready to work!! Check", slog.String("Adress",cfg.HTTPServer.Address) )
	// var wg sync.WaitGroup
	// wg.Add(1);
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM);
	var err error
	go func(/*wg *sync.WaitGroup*/){
		// defer wg.Done()
		err = server.ListenAndServe()
		if (err != nil) {
			log.Fatal("Failed to start server")
		}
	}(/*&wg*/)
	// wg.Wait()
	<-done
	slog.Info("\nshutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err!=nil{
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
	
}