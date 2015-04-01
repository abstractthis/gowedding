package controllers

import (
    "os"
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "unicode"
    "html/template"
    "strings"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
    dj, err := json.MarshalIndent(d, "", " ")
    if err != nil {
        http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
        Logger.Println(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(c)
    fmt.Fprintf(w, "%s", dj)
}

func CreateResponse(w http.ResponseWriter, locURL string) {
    headers := w.Header()
    headers.Set("Content-Length", "0")
    headers.Set("Content-Type", "text/plain; charset=utf-8")
    headers.Set("Location", locURL)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "")
}

func CodeResponse(w http.ResponseWriter, c int) {
    headers := w.Header()
    headers.Set("Content-Length", "0")
    headers.Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(c)
    fmt.Fprint(w, "")
}

func ErrorResponse(w http.ResponseWriter, msg string) {
    Logger.Println(msg)
    http.Error(w, msg, http.StatusInternalServerError)
}

func ErrorResponseWithPayload(w http.ResponseWriter, c int) {
    var tName string
    if c == http.StatusNotFound {
        tName = "templates/404.html"
    } else {
        tName = "templates/doh.html"
    }
    t, err := template.ParseFiles(tName)
    if err != nil {
        ErrorResponse(w, "Template Parse failure: " + tName)
        return
    }
    w.WriteHeader(c)
    t.Execute(w, nil)
}

func FirstLetterUpper(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

func CapIt(s string) string {
    // Deal with hyphenated names first
    hyphenChop := strings.Split(s, "-")
    if len(hyphenChop) > 1 {
        name := make([]string, len(hyphenChop))
        for i, partial := range hyphenChop {
            name[i] = FirstLetterUpper(partial)
        }
        return strings.Join(name, "-")
    } else {
        return FirstLetterUpper(s)
    }
}

func FullName(first string, last string) string {
    if first == "" && last == "" {
        return ""
    }
    first = CapIt(first)
    last = CapIt(last)
    return first + " " + last
}

/*
 * Used in the template for rendering Guest pagination values.
 */
func Sum(x, y int) int {
    return x + y
}
