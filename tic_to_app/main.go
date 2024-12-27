// main.go
package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Route handlers
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/play", handlePlay)

    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/home.html"))
    tmpl.Execute(w, nil)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/game.html"))
    tmpl.Execute(w, nil)
}

