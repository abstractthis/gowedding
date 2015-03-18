package routes

import (
    "log"
    "os"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/abstractthis/gowedding/controllers"
    //"github.com/justinas/nosurf"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func CreateWeddingRouter(serveStatic bool) http.Handler {
    Logger.Println("Creating Wedding Router...")
    router := mux.NewRouter()
    router.StrictSlash(true)
    router.HandleFunc("/rsvp/{id:[0-9]+}/{first:[a-z]+}/{last:[a-z]+}/{nonce:[a-z0-9]{40}}/{stamp:[0-9]+}/", controllers.Respondez).Methods("GET")
    router.HandleFunc("/rsvp/", controllers.Respondez).Methods("POST")
    router.HandleFunc("/rsvp/reply/", controllers.RSVP_Reply).Methods("POST")

    // Setup static handlers if need be
    if serveStatic {
        router.HandleFunc("/wedding", func(w http.ResponseWriter, r *http.Request) {
            http.ServeFile(w, r, "static/html/index.html")
        })
        fs := http.FileServer(http.Dir("./static/"))
        router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    }
    return router

    // Setup Cross-site Request Forgery protection
    // csrfProtectedHandler := nosurf.New(router)
    // csrfProtectedHandler.ExemptRegexp("/rsvp/[a-z]+/[a-z]+/[a-z0-9]{40}")
    // return csrfProtectedHandler
}