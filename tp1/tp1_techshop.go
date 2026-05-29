package main

import "fmt"

// Struct Produit
type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

// Struct Catalogue (slice de Produits)
type Catalogue struct {
	Produits []Produit
}

func main() {
	catalogue := Catalogue{
		Produits: []Produit{
			{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 1199.00, Stock: 10, Categorie: "Smartphone", Actif: true},
			{ID: 2, Nom: "MacBook Pro", Marque: "Apple", Prix: 2499.00, Stock: 5, Categorie: "Ordinateur", Actif: true},
			{ID: 3, Nom: "Galaxy S24", Marque: "Samsung", Prix: 999.00, Stock: 8, Categorie: "Smartphone", Actif: true},
			{ID: 4, Nom: "Dell XPS 15", Marque: "Dell", Prix: 1799.00, Stock: 3, Categorie: "Ordinateur", Actif: true},
			{ID: 5, Nom: "AirPods Pro", Marque: "Apple", Prix: 279.00, Stock: 20, Categorie: "Audio", Actif: true},
		},
	}

	// affichage temporaire pour vérifier
	for _, p := range catalogue.Produits {
		fmt.Printf("[%d] %s - %s - %.2f€ (stock: %d)\n", p.ID, p.Nom, p.Marque, p.Prix, p.Stock)
	}
}
