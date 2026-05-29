package main

import "fmt"

// struct Personne
type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

// p Personne : copie mais pas de * car on ne modifie pas la personne, juste on utilise ses données pour construire une présentation
func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}
func (p Personne) Presentation() string {
	return fmt.Sprintf("Je m'appelle %s, j'ai %d ans. Email : %s", p.NomComplet(), p.Age, p.Email)
}

func main() {
	p := Personne{Prenom: "Hugo", Nom: "Duperthuy", Age: 24, Email: "hugo.duperthuy@outlook.fr"}
	fmt.Println(p.NomComplet())
	fmt.Println(p.Presentation())
}
