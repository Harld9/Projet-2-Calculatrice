package controller

import (
	"calculatrice/logic"
	"html/template"
	"net/http"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ResSom":     "",
		"EntreeSom1": "",
		"EntreeSom2": "",
		"ErreurSom1": "",
		"ErreurSom2": "",
		"ResSou":     "",
		"EntreeSou1": "",
		"EntreeSou2": "",
		"ErreurSou1": "",
		"ErreurSou2": "",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func AdditionWeb(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ResSom":     "",
		"EntreeSom1": "",
		"EntreeSom2": "",
		"ErreurSom1": "",
		"ErreurSom2": ""}
	if r.Method == http.MethodPost {
		nb1 := r.FormValue("nb1")
		nb2 := r.FormValue("nb2")
		nombre1, err1 := strconv.ParseFloat(nb1, 64)
		if err1 != nil {
			data["ErreurSom1"] = "Veuillez entrer un premier nombre valide"
		} else {
			data["EntreeSom1"] = nombre1
		}
		nombre2, err2 := strconv.ParseFloat(nb2, 64)
		if err2 != nil {
			data["ErreurSom2"] = "Veuillez entrer un deuxième nombre valide"
		} else {
			data["EntreeSom2"] = nombre2
		}
		if err1 == nil && err2 == nil {
			resultat := logic.Addition(nombre1, nombre2)
			if resultat == 0 {
				zero := "0"
				data["ResSom"] = zero
			} else if resultat != 0 {
				data["ResSom"] = resultat
			}
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func SoustractionWeb(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ResSou":     "",
		"EntreeSou1": "",
		"EntreeSou2": "",
		"ErreurSou1": "",
		"ErreurSou2": ""}
	if r.Method == http.MethodPost {
		nb1 := r.FormValue("nb1")
		nb2 := r.FormValue("nb2")
		nombre1, err1 := strconv.ParseFloat(nb1, 64)
		if err1 != nil {
			data["ErreurSou1"] = "Veuillez entrer un premier nombre valide"
		} else {
			data["EntreeSou1"] = nombre1
		}
		nombre2, err2 := strconv.ParseFloat(nb2, 64)
		if err2 != nil {
			data["ErreurSou2"] = "Veuillez entrer un deuxième nombre valide"
		} else {
			data["EntreeSou2"] = nombre2
		}

		if err1 == nil && err2 == nil {
			resultat := logic.Soustraction(nombre1, nombre2)
			if resultat == 0 {
				zero := "0"
				data["ResSou"] = zero
			} else if resultat != 0 {
				data["ResSou"] = resultat
			}
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func MultiplicationWeb(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ResMul":     "",
		"EntreeMul1": "",
		"EntreeMul2": "",
		"ErreurMul1": "",
		"ErreurMul2": ""}
	if r.Method == http.MethodPost {
		nb1 := r.FormValue("nb1")
		nb2 := r.FormValue("nb2")
		nombre1, err1 := strconv.ParseFloat(nb1, 64)
		if err1 != nil {
			data["ErreurMul1"] = "Veuillez entrer un premier nombre valide"
		} else {
			data["EntreeMul1"] = nombre1
		}
		nombre2, err2 := strconv.ParseFloat(nb2, 64)
		if err2 != nil {
			data["ErreurMul2"] = "Veuillez entrer un deuxième nombre valide"
		} else {
			data["EntreeMul2"] = nombre2
		}

		if err1 == nil && err2 == nil {
			resultat := logic.Multiplication(nombre1, nombre2)
			if resultat == 0 {
				zero := "0"
				data["ResMul"] = zero
			} else if resultat != 0 {
				data["ResMul"] = resultat
			}
		}

	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func DivisionWeb(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ResDiv":     "",
		"EntreeDiv1": "",
		"EntreeDiv2": "",
		"ErreurDiv1": "",
		"ErreurDiv2": ""}
	if r.Method == http.MethodPost {
		nb1 := r.FormValue("nb1")
		nb2 := r.FormValue("nb2")
		nombre1, err1 := strconv.ParseFloat(nb1, 64)
		if err1 != nil {
			data["ErreurDiv1"] = "Veuillez entrer un premier nombre valide"
		} else {
			data["EntreeDiv1"] = nombre1
		}
		nombre2, err2 := strconv.ParseFloat(nb2, 64)
		if err2 != nil || nombre2 == 0 {
			data["ErreurDiv2"] = "Veuillez entrer un deuxième nombre valide (Impossible de diviser par 0)"
		} else {
			data["EntreeDiv2"] = nombre1
		}

		if err1 == nil && err2 == nil {
			resultat := logic.Division(nombre1, nombre2)
			if resultat == 0 {
				zero := "0"
				data["ResDiv"] = zero
			} else if resultat != 0 {
				data["ResDiv"] = resultat
			}
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
