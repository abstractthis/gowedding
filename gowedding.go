package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/handlers"
    "github.com/abstractthis/gowedding/config"
    "github.com/abstractthis/gowedding/routes"
    "github.com/abstractthis/gowedding/models"
)

 var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

 func main() {
    // Initialize the database and its globals
    err := models.Initialize()
    if err != nil {
        Logger.Printf("Model Initialization failed: %v\n", err)
        os.Exit(1)
    }

    // Fire up the api server
    Logger.Printf("Launching gowedding api server at http://%s\n", config.Conf.ApiURL)
    http.ListenAndServe(config.Conf.ApiURL,
        handlers.CombinedLoggingHandler(os.Stdout, routes.CreateWeddingRouter(config.Conf.IsDev)))
 }