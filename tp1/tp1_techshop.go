package main

import (
	"errors"
	"fmt"
	"strings"
)

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

// Ajoute un produit au catalogue
// *Catalogue : pointer receiver car on modifie la slice
func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, existing := range c.Produits {
		if existing.ID == p.ID {
			return fmt.Errorf("produit avec l'ID %d déjà existant", p.ID)
		}
	}
	c.Produits = append(c.Produits, p)
	return nil
}

// Retourne le produit correspondant à l'ID, ou une erreur si introuvable
// Retour (Produit, error) : idiome Go pour retourner une valeur + erreur
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.Produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("produit avec l'ID %d introuvable", id)
}

// Retourne tous les produits d'une catégorie
// strings.EqualFold : comparaison insensible à la casse ("audio" == "Audio")
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit
	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultats = append(resultats, p)
		}
	}
	return resultats
}

// Applique un % de réduction sur tous les produits d'une catégorie
// *Catalogue : on modifie les prix
// Retourne le nombre de produits modifiés
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for i, p := range c.Produits {
		if strings.EqualFold(p.Categorie, categorie) {
			c.Produits[i].Prix -= c.Produits[i].Prix * pct / 100
			count++
		}
	}
	return count
}

// Réduit le stock d'un produit
// *Catalogue : on modifie le stock
// Erreur si produit introuvable ou stock insuffisant
func (c *Catalogue) Vendre(id int, qte int) error {
	for i, p := range c.Produits {
		if p.ID == id {
			if p.Stock < qte {
				return fmt.Errorf("stock insuffisant pour '%s' (stock: %d, demandé: %d)", p.Nom, p.Stock, qte)
			}
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return errors.New("produit introuvable")
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

	// test AjouterProduit
	nouveau := Produit{ID: 6, Nom: "Sony WH-1000XM5", Marque: "Sony", Prix: 349.00, Stock: 7, Categorie: "Audio", Actif: true}
	if err := catalogue.AjouterProduit(nouveau); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Produit ajouté :", nouveau.Nom)
	}

	// test doublon
	if err := catalogue.AjouterProduit(Produit{ID: 1, Nom: "Doublon"}); err != nil {
		fmt.Println("Erreur :", err)
	}

	// test TrouverParID
	fmt.Println("\nTrouverParID")
	if p, err := catalogue.TrouverParID(3); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("Trouvé : [%d] %s - %.2f€\n", p.ID, p.Nom, p.Prix)
	}
	if _, err := catalogue.TrouverParID(99); err != nil {
		fmt.Println("Erreur :", err)
	}

	// test TrouverParCategorie
	fmt.Println("\n TrouverParCategorie : audio")
	resultats := catalogue.TrouverParCategorie("audio")
	for _, p := range resultats {
		fmt.Printf("  [%d] %s\n", p.ID, p.Nom)
	}

	// test AppliquerReduction
	fmt.Println("\nAppliquerReduction : -20% sur Smartphone")
	n := catalogue.AppliquerReduction("Smartphone", 20)
	fmt.Printf("%d produit(s) modifié(s)\n", n)
	for _, p := range catalogue.TrouverParCategorie("Smartphone") {
		fmt.Printf("  [%d] %s → %.2f€\n", p.ID, p.Nom, p.Prix)
	}

	// test Vendre
	fmt.Println("\nVendre")
	if err := catalogue.Vendre(1, 3); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		p, _ := catalogue.TrouverParID(1)
		fmt.Printf("Vendu 3x iPhone 15 → stock restant : %d\n", p.Stock)
	}
	// stock insuffisant
	if err := catalogue.Vendre(2, 100); err != nil {
		fmt.Println("Erreur :", err)
	}
}
