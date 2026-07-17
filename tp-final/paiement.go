package main

import (
	"fmt"
	"math"
	"strings"
)

// Taux de conversion : 1 BTC vaut 50000€
const TauxBTC = 50000.0

// Interface (liste de méthodes)
type Payeur interface {
	Payer(montant float64) (string, error)
}

// Compilation : si erreurs elles apparaissent la
var _ Payeur = &CarteCredit{}
var _ Payeur = &PayPal{}
var _ Payeur = &Crypto{}

// Carte
type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

// *CarteCredit : pointer receiver, car on modifie le solde
func (cc *CarteCredit) Payer(montant float64) (string, error) {
	if montant > cc.Solde {
		return "", fmt.Errorf("solde de la carte bancaire insuffisant (solde: %.2f€, demandé: %.2f€)", cc.Solde, montant)
	}

	cc.Solde -= montant

	quatreDerniers := cc.Numero[len(cc.Numero)-4:]

	return fmt.Sprintf("Transaction CB #%s confirmée", quatreDerniers), nil
}

// 1er chiffre indique le réseau de la carte
func (cc CarteCredit) Reseau() string {
	if strings.HasPrefix(cc.Numero, "4") {
		return "Visa"
	}
	if strings.HasPrefix(cc.Numero, "5") {
		return "Mastercard"
	}
	return "Inconnu"
}

// Paypal
type PayPal struct {
	Email string
	Solde float64
}

func (pp *PayPal) Payer(montant float64) (string, error) {
	if montant > pp.Solde {
		return "", fmt.Errorf("solde PayPal insuffisant (solde: %.2f€, demandé: %.2f€)", pp.Solde, montant)
	}

	pp.Solde -= montant

	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, pp.Email), nil
}

// Crypto
type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	// On convertit les euros en BTC, arrondi au millième
	enCrypto := math.Round(montant/TauxBTC*1000) / 1000

	if enCrypto > c.Solde {
		return "", fmt.Errorf("solde %s insuffisant (solde: %.3f, demandé: %.3f)", c.Monnaie, c.Solde, enCrypto)
	}

	c.Solde -= enCrypto

	return fmt.Sprintf("Paiement de %.3f %s (%.2f€) vers %s", enCrypto, c.Monnaie, montant, c.Adresse), nil
}

// Processer le panier

// Payeur (une interface) marche avec la CB, PayPal ou la crypto
func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}
	fmt.Printf("\nPanier : %d article(s) — Total : %.2f€\n", len(articles), total)

	switch v := payeur.(type) {
	case *CarteCredit:
		fmt.Printf("Mode : Carte %s de %s\n", v.Reseau(), v.Titulaire)
	case *PayPal:
		fmt.Printf("Mode : PayPal (%s)\n", v.Email)
	case *Crypto:
		fmt.Printf("Mode : Crypto %s\n", v.Monnaie)
	}

	recu, err := payeur.Payer(total)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println(recu)
}

// Créer les types de paiement et passe par le panier
func main() {
	cb := &CarteCredit{Numero: "4970123456789012", Titulaire: "Hugo Dubois", Solde: 1500}
	pp := &PayPal{Email: "hugo@example.com", Solde: 300}
	btc := &Crypto{Adresse: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", Solde: 0.5, Monnaie: "BTC"}

	panier := []float64{29.99, 149.50, 12.00}

	moyens := []Payeur{cb, pp, btc}

	for _, moyen := range moyens {
		ProcesserPanier(moyen, panier)
	}

	ProcesserPanier(pp, []float64{5000})
}
