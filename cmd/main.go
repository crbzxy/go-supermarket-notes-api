package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "go-supermarket-notes-api/config"
    "go-supermarket-notes-api/routes"
    _ "go-supermarket-notes-api/docs" // Importa la documentación generada
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title Supermarket and Notes API
// @version 1.0
// @description This is a sample server for managing supermarket lists and notes.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.basic BasicAuth
func main() {
    // Verifica el directorio actual y el archivo .env
    dir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current directory: %v", err)
    }
    log.Println("Current directory:", dir)

    // Verifica que el archivo .env esté en el directorio correcto
    if _, err := os.Stat(".env"); os.IsNotExist(err) {
        log.Fatalf(".env file does not exist in the current directory")
    }

    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    } else {
        log.Println("Successfully loaded .env file")
        log.Println("MONGODB_URI:", os.Getenv("MONGODB_URI"))
        log.Println("DB_NAME:", os.Getenv("DB_NAME"))
        log.Println("PORT:", os.Getenv("PORT"))
    }

    router := mux.NewRouter()
    routes.RegisterRoutes(router)

    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    config.ConnectDB()

    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
