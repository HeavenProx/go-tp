package main

import (
	"fmt"
	"time"
)

func main() {
	// message et commentaires
	naissance := time.Date(2002, time.March, 11, 0, 0, 0, 0, time.UTC)
	maintenant := time.Now()

	years := maintenant.Sub(naissance).Hours() / 24 / 365
	mois := int(years*12) % 12
	jours := int(years*365)%365 - int(mois*30)
	heures := int(years*365*24) % 24

	fmt.Printf("Je suis né le 11 mars 2002.\n")
	fmt.Printf("J'ai %d ans, %d mois, %d jours et %d heures\n", int(years), mois, jours, heures)
}
