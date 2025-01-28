/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	bookHttp "github.com/alfaysal/go-pet-project/app/book/delivery"
	bookRepository "github.com/alfaysal/go-pet-project/app/book/repository"
	bookUsecase "github.com/alfaysal/go-pet-project/app/book/usecase"
	"github.com/alfaysal/go-pet-project/internal/config"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/alfaysal/go-pet-project/internal/conn"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A command for serving the application.",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		if err := conn.ConnectDefaultDB(); err != nil {
			fmt.Println(err)
		}

		err := conn.ConnectDefaultRedis()
		if err != nil {
			fmt.Println("can not connect to database")
		}

		conn.InitHttpClient()

	},
	Run: func(cmd *cobra.Command, args []string) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		r := chi.NewRouter()
		r.Use(middleware.Logger)

		cfg := config.NewApplication()

		bookRepository := bookRepository.New(conn.GetDefaultDB())

		// cache sample
		cacheInstance := conn.DefaultCache()
		//
		//cacheInstance.Set("name", "faysal", time.Millisecond*1000)
		//
		//fmt.Println(cacheInstance.Get("name"))
		bookUsecase := bookUsecase.New(bookRepository, cacheInstance)

		apiRouter := chi.NewRouter()
		r.Mount("/api", apiRouter)
		bookHttp.New(apiRouter, bookUsecase)

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
			Handler: r,
		}

		go func(sr *http.Server) {
			if err := sr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Something went wrong")
			}
		}(srv)

		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Shutting down successfully")
		}

		fmt.Println("Server gracefully stopped")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func fetchFromHttp(urls []string, result chan<- string, client *http.Client) {

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Set headers (optional)
		req.Header.Set("Accept", "application/json")
		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
		}

		defer resp.Body.Close()
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		result <- string(body)
	}
}
