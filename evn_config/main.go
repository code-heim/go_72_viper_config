package main

import (
	"go_blogs/config"
	"go_blogs/controllers"
	"go_blogs/models"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	config.LoadConfig()
	addr := viper.GetString("SERVER_ADDRESS")

	models.ConnectDatabase()
	models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/blogs", controllers.BlogsIndex)
	mux.HandleFunc("/blogs/", controllers.BlogShow)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
