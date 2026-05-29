package main

import "fmt"

// const + iota
type Niveau int

const (
	Debutant      Niveau = iota // 0
	Intermediaire               // 1
	Avance                      // 2
	Expert                      // 3
)

func (n Niveau) String() string {
	return [...]string{"Débutant", "Intermédiaire", "Avancé", "Expert"}[n]
}

// fallthrough dans le switch case
// va retourner les compétences de chaque niveau et ceux d'en dessous
func decrireNiveau(n Niveau) {
	fmt.Printf("Niveau %s :\n", n)
	switch n {
	case Expert:
		fmt.Println("  - Maîtrise complète")
		fallthrough
	case Avance:
		fmt.Println("  - Concepts avancés acquis")
		fallthrough
	case Intermediaire:
		fmt.Println("  - Bases solides")
		fallthrough
	case Debutant:
		fmt.Println("  - En apprentissage")
	}
}

func main() {
	// --- slice ---
	scores := []int{42, 87, 15, 93, 56, 71}

	// for classique (avec index)
	fmt.Println("=== Scores ===")
	for i, s := range scores {
		fmt.Printf("  [%d] %d\n", i, s)
	}

	// for comme while : trouver le max
	max := scores[0]
	i := 1
	for i < len(scores) {
		if scores[i] > max {
			max = scores[i]
		}
		i++
	}
	fmt.Println("Max :", max)

	// for unique
	// va filtrer les scores > 50
	fmt.Println("\n Scores > 50 :")
	idx := 0
	for {
		if idx >= len(scores) {
			break
		}
		if scores[idx] > 50 {
			fmt.Println(" ", scores[idx])
		}
		idx++
	}

	// array
	fmt.Println("\n=== Array ===")
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array :", arr)
	fmt.Println("len :", len(arr)) // toujours égal à la taille fixe

	// len() et cap()
	fmt.Println("\n=== len() et cap() ===")

	// slice littéral
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 = %v  len=%d  cap=%d\n", s1, len(s1), cap(s1))

	// slice avec make : make([]T, longueur, capacité)
	s2 := make([]int, 3, 7) // longueur 3, capacité 7
	fmt.Printf("s2 = %v  len=%d  cap=%d\n", s2, len(s2), cap(s2))
}
