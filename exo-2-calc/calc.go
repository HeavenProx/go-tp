package main

import (
	"errors"
	"fmt"
)

// fonction qui prend deux nombres et une opération, et retourne le résultat de l'opération
func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opération inconnue : %s", op)
	}
}

// fonction qui prend une opération et retourne une fonction qui effectue cette opération
func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return nil
	}
}

// fonction qui prend une opération et retourne une fonction qui effectue cette opération, en utilisant une closure
func main() {
	for {
		var a, b float64
		var op string

		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			break
		}

		result, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Résultat :", result)
		}
	}
}
