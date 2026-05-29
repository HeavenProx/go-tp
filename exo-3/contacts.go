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

// struct Adresse
type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

// retourne l'adresse formatée
func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// struct Employe (embedding Personne + Adresse)
type Employe struct {
	Personne // peut utiliser les champs et méthodes de Personne directement
	Adresse  // peut utiliser les champs et méthodes de Adresse directement
	Poste    string
	Salaire  float64
}

// FicheEmploye retourne toutes les infos de l'employé
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Fiche Employé :\n  Nom     : %s\n  Age     : %d\n  Email   : %s\n  Poste   : %s\n  Salaire : %.2f€\n  Adresse : %s",
		e.NomComplet(), e.Age, e.Email, e.Poste, e.Salaire, e.Adresse.Format(),
	)
}

// Sans * on modifierait une copie et l'original ne changerait pas
func (e *Employe) AugmenterSalaire(pourcent float64) {
	e.Salaire += e.Salaire * pourcent / 100
}

// struct Etudiant (embedding Personne) ---

type Etudiant struct {
	Personne // peut utiliser les champs et méthodes de Personne directement
	Promo    string
	Moyenne  float64
}

// Retourne la mention selon la moyenne avec un switch
// pas de fallthrough, on s'arrète direct à la bonne mention
func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très Bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez Bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Insuffisant"
	}
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"Fiche Étudiant :\n  Nom     : %s\n  Age     : %d\n  Promo   : %s\n  Moyenne : %.2f\n  Mention : %s",
		e.NomComplet(), e.Age, e.Promo, e.Moyenne, e.MentionObtenue(),
	)
}

func main() {
	// mon objet Personne
	p := Personne{Prenom: "Hugo", Nom: "Duperthuy", Age: 24, Email: "hugo.duperthuy@outlook.fr"}
	fmt.Println(p.NomComplet())
	fmt.Println(p.Presentation())

	// mon objet Adresse
	a := Adresse{Rue: "11 rue du bowling", Ville: "Lyon", CodePostal: "69008"}
	fmt.Println(a.Format())

	// mon object Employe
	employes := []Employe{
		{
			Personne: Personne{Prenom: "Alice", Nom: "Martin", Age: 32, Email: "alice@corp.fr"},
			Adresse:  Adresse{Rue: "3 av. des Fleurs", Ville: "Paris", CodePostal: "75008"},
			Poste:    "Développeuse",
			Salaire:  3500,
		},
		{
			Personne: Personne{Prenom: "Bob", Nom: "Lebrun", Age: 45, Email: "bob@corp.fr"},
			Adresse:  Adresse{Rue: "7 rue du Port", Ville: "Bordeaux", CodePostal: "33000"},
			Poste:    "Manager",
			Salaire:  5000,
		},
	}

	// augmenter le salaire de tous les employés de 10%
	for i := range employes {
		employes[i].AugmenterSalaire(10)
	}

	// afficher les infos des employés
	for _, e := range employes {
		fmt.Println(e.FicheEmploye())
	}

	// --- Étudiants dans une slice ---
	etudiants := []Etudiant{
		{
			Personne: Personne{Prenom: "Clara", Nom: "Petit", Age: 21, Email: "clara@univ.fr"},
			Promo:    "M2IW 2026",
			Moyenne:  17.5,
		},
		{
			Personne: Personne{Prenom: "Tom", Nom: "Renard", Age: 22, Email: "tom@univ.fr"},
			Promo:    "M2IW 2026",
			Moyenne:  11.8,
		},
	}

	for _, et := range etudiants {
		fmt.Println(et.FicheEtudiant())
	}
}
