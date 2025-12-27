package router

import (
	"calculatrice/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/addition", controller.AdditionWeb)
	mux.HandleFunc("/soustraction", controller.SoustractionWeb)
	mux.HandleFunc("/multiplication", controller.MultiplicationWeb)
	mux.HandleFunc("/division", controller.DivisionWeb)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
