package router

import (
	"calculatrice/controller"
	"net/http"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Home)
	// Routes de ton app

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
