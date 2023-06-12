package main

import (
	"fmt"
	"net/http"
	"portfolio/handlers"
)

func main() {
	fmt.Println("Démarage du serveur")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/project", handlers.Project)
	http.HandleFunc("/contact", handlers.Contact)
	http.ListenAndServe(":5500", nil)
}
