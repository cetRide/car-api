package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cetRide/car-api/models"
	"github.com/cetRide/car-api/routers"

	"github.com/gorilla/handlers"
	_ "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("PORT")
	dbUrI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	models.InitDB(dbUrI)

	router := routers.NewRouter()
	fmt.Println(port)
	err = http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		panic(err)
	}
}
