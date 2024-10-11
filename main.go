package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

type StockageForm struct {
	CheckValue bool
	Value      string
}

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

type Change struct {
	Paire bool
	Nb    int
}

var stockageFormNom = StockageForm{false, ""}
var stockageFormPrenom = StockageForm{false, ""}
var stockageFormDate = StockageForm{false, ""}
var stockageFormSexe = StockageForm{false, ""}

func main() {

	Page := Change{
		Paire: false,
		Nb:    -1,
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./template/Promo.html")
		if err != nil {
			fmt.Print("Toi tu as fais une connerie du type : %s", err.Error())
			return
		}
		B1Info := Classe{
			NameClasse: "B1 Info/Cyber",
			Filiere:    "Informatique/Cybersécurité",
			Niveau:     1,
			NbEtudiant: 5,
			Etudiant:   []Profil{Profil{"Adrien", "LECOMTE", 20, "M"}, Profil{"Alexandre", "PETITFRERE", 20, "M"}, Profil{"Azilis", "ROSELLO", 19, "F"}, Profil{"Jonathan", "PEREZ", 19, "M"}, Profil{"Léo", "VELAZQUEZ", 18, "M"}},
		}

		temp.Execute(w, B1Info)

	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./template/Change.html")
		if err != nil {
			fmt.Print("Toi tu as fais une connerie du type : %s", err.Error())
			return
		}
		Page.Nb++
		if Page.Nb%2 != 0 {
			Page.Paire = false
		} else if Page.Nb%2 == 0 {
			Page.Paire = true
		}

		temp.Execute(w, Page)
	})

	type User struct {
		Nom           string
		Prenom        string
		DateNaissance string
		Sexe          string
	}
	// Route pour afficher le formulaire
	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./template/Formulaire.html")
		if err != nil {
			fmt.Print("Toi tu as fais une connerie du type : %s", err.Error())
			return
		}
		fmt.Println("LoL")

		temp.Execute(w, nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("toi t'a fais de la merde")
			return
		}

		checkValueNom, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("nom"))
		checkValuePrenom, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("prenom"))
		if !checkValueNom {
			stockageFormNom = StockageForm{false, ""}
			http.Redirect(w, r, "/user/form", http.StatusSeeOther)
			return
		}
		if !checkValuePrenom {
			stockageFormPrenom = StockageForm{false, ""}
			http.Redirect(w, r, "/user/form", http.StatusSeeOther)
			return
		}

		stockageFormNom = StockageForm{true, r.FormValue("nom")}
		stockageFormPrenom = StockageForm{true, r.FormValue("prenom")}
		stockageFormDate = StockageForm{true, r.FormValue("dateNaissance")}
		stockageFormSexe = StockageForm{true, r.FormValue("sexe")}

		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})

	type Envoie struct {
		Nom    string
		Prenom string
		Age    string
		Sexe   string
	}

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./template/FormDisplay.html")
		if err != nil {
			fmt.Print("Toi tu as fais une connerie du type : %s", err.Error())
			return
		}
		Data := Envoie{stockageFormNom.Value, stockageFormPrenom.Value, stockageFormDate.Value, stockageFormSexe.Value}
		temp.Execute(w, Data)
	})

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
