package main

import "fmt"

func main() {
	var poids float64 = 90.0
	var taille float64 = 1.81

	const (
		Nom         = "Hugo"
		IMCMaigreur = 18.5
		IMCNormal   = 25.0
		IMCSurpoids = 30.0
	)

	var imc float64 = poids / (taille * taille)

	fmt.Printf("Mon nom est %s :)\n", Nom)
	fmt.Printf("L'IMC est de : %.2f\n", imc)

	var categorie string
	switch {
	case imc < IMCMaigreur:
		categorie = "Maigreur"
	case imc < IMCNormal:
		categorie = "Normal"
	case imc < IMCSurpoids:
		categorie = "Surpoids"
	default:
		categorie = "Obésité"
	}

	fmt.Printf("C'est donc une catégorie : %s\n", categorie)
}
