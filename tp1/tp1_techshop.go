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
	// Ajoute le produit à la slice
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
			// Modifie directement le prix dans la slice
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
			// Modifie directement le stock dans la slice
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return errors.New("produit introuvable")
}

// Retourne un résumé : nb produits et valeur totale du stock
// Pas de pointer receiver : on lit seulement
func (c Catalogue) Rapport() string {
	total := 0.0
	for _, p := range c.Produits {
		total += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("Catalogue : %d produits | Valeur totale du stock : %.2f€", len(c.Produits), total)
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

	for {
		fmt.Println("\n=== TechShop ===")
		fmt.Println("[1] Ajouter  [2] Chercher  [3] Soldes  [4] Vendre  [5] Rapport  [0] Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			fmt.Println("Fermeture du menu")
			return

		case 1: // Ajouter un produit
			var p Produit
			fmt.Print("ID : ")
			fmt.Scan(&p.ID)
			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)
			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)
			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)
			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)
			fmt.Print("Categorie : ")
			fmt.Scan(&p.Categorie)
			p.Actif = true
			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté :", p.Nom)
			}

		case 2: // Chercher par ID
			var id int
			fmt.Print("ID du produit : ")
			fmt.Scan(&id)
			if p, err := catalogue.TrouverParID(id); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("[%d] %s - %s - %.2f€ (stock: %d, catégorie: %s)\n",
					p.ID, p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie)
			}

		case 3: // Appliquer une réduction sur une catégorie
			var cat string
			var pct float64
			fmt.Print("Catégorie : ")
			fmt.Scan(&cat)
			fmt.Print("Réduction (%) : ")
			fmt.Scan(&pct)
			n := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("%d produit(s) mis en solde (-%.0f%%)\n", n, pct)

		case 4: // Vendre
			var id, qte int
			fmt.Print("ID du produit : ")
			fmt.Scan(&id)
			fmt.Print("Quantité : ")
			fmt.Scan(&qte)
			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				p, _ := catalogue.TrouverParID(id)
				fmt.Printf("Vente OK — stock restant pour '%s' : %d\n", p.Nom, p.Stock)
			}

		case 5: // Rapport
			fmt.Println(catalogue.Rapport())

		default:
			fmt.Println("Choix invalide")
		}
	}
}
