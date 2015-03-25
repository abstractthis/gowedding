package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    // "syscall"

    "github.com/gorilla/handlers"
    "github.com/abstractthis/gowedding/config"
    "github.com/abstractthis/gowedding/routes"
    "github.com/abstractthis/gowedding/models"
    "github.com/abstractthis/gowedding/emailer"
)

 var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)
 var emailSender *emailer.Emailer

 func main() {
    // Launch the emailer
    emailSender = emailer.New()
    // emailSender.Start()

    // Spin up goroutine to listen and deal with Ctrl-C
    // Actually listens for SIGINT, SIGKILL and SIGTERM
    sigChan := make(chan os.Signal, 2)
    stopSig := make(chan bool)
    go func() {
        signal.Notify(sigChan, os.Interrupt)
        select {
        case <-sigChan:
            Logger.Printf("Program interrupt received! Cleanup...")
            emailSender.Stop()
            Logger.Println("cleanup complete.")
            os.Exit(1)
        case <-stopSig:
            return
        }
    }()

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

    // Shutdown emailer on normal exit
    emailSender.Stop()
    // Stop the goroutine listening for interrupt signals
    stopSig <- true
 }