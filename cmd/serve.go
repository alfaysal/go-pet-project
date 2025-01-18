/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/alfaysal/go-pet-project/internal/conn"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis"
	"log"
	"net/http"
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
		r := chi.NewRouter()
		r.Use(middleware.Logger)

		r.Route("/books", func(r chi.Router) {
			r.Post("/", createBook)
			r.Get("/", listBooks)
			r.Get("/{id}", getBook)
			r.Put("/{id}", updateBook)
			//r.Delete("/{id}", deleteBook)
		})

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		pong, err := client.Ping().Result()

		if err != nil {
			fmt.Println("Failed to connect to redis")
		}

		fmt.Println(pong)

		err = client.Set("name", "Al faysal", time.Millisecond*1000).Err()

		if err != nil {
			fmt.Println("Failed to save to redis")
		}

		fmt.Println(client.Get("name").Val())
		log.Fatal(http.ListenAndServe(":8099", r))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	w.Header().Set("Content-Type", "application/json")

	DB := conn.GetDefaultDB()
	if err := DB.Find(&books).Error; err != nil {
		fmt.Println(err)
	}

	err := json.NewEncoder(w).Encode(books)

	if err != nil {
		fmt.Println(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	DB := conn.GetDefaultDB()

	if err := DB.First(&book, id).Error; err != nil {
		fmt.Println(err)
	}

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		fmt.Println(err)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	w.Header().Set("Content-Type", "application/json")

	DB := conn.GetDefaultDB()

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println(err)
	}

	if err := DB.Create(&book).Error; err != nil {
		fmt.Println(err)
	}

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		fmt.Println(err)
	}

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	DB := conn.GetDefaultDB()

	if err := DB.First(&book, id).Error; err != nil {
		fmt.Println(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println(err)
	}

	if err := DB.Save(&book).Error; err != nil {
		fmt.Println(err)
	}

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		fmt.Println(err)
	}
}
