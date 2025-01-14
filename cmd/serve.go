/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	bookHttp "github.com/alfaysal/go-pet-project/book/delivery"
	bookRepository "github.com/alfaysal/go-pet-project/book/repository"
	bookUsecase "github.com/alfaysal/go-pet-project/book/usecase"
	"github.com/alfaysal/go-pet-project/internal/config"
	"github.com/alfaysal/go-pet-project/internal/conn"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A command for serving the application.",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		if err := conn.ConnectDefaultDB(); err != nil {
			fmt.Println(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		r := chi.NewRouter()
		r.Use(middleware.Logger)

		cfg := config.NewApplication()

		bookRepository := bookRepository.New(conn.GetDefaultDB())

		bookUsecase := bookUsecase.New(bookRepository)

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
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
