package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	type Profil struct {
		FirstName string
		LastName  string
		Age       int
		Sexe      string
	}

	type Classe struct {
		NameClasse string
		Filiere    string
		Niveau     int
		NbEtudiant int
		Etudiant   []Profil
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./template/Promo.html")
		if err != nil {
			fmt.Print("Toi tu as fais une connerie du type : %s", err.Error())
			return
		}
		B1Info := Classe{
			NameClasse: "B1 Info",
			Filiere:    "Informatique",
			Niveau:     1,
			NbEtudiant: 5,
			Etudiant:   []Profil{Profil{"Adrien", "LECOMTE", 20, "M"}, Profil{"Alexandre", "PETITFRERE", 20, "M"}, Profil{"Azilis", "ROSELLO", 19, "F"}, Profil{"Jonathan", "PEREZ", 19, "M"}, Profil{"Léo", "VELAZQUEZ", 18, "M"}},
		}

		temp.Execute(w, B1Info)
	})

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
