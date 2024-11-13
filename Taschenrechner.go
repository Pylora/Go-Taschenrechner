package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Benutzer nach der ersten Zahl fragen
	fmt.Print("Gib die erste Zahl ein: ")
	var num1 string
	fmt.Scanln(&num1)

	// Benutzer nach der Operation fragen
	fmt.Print("Gib die Operation ein (+, -, *, /): ")
	var operator string
	fmt.Scanln(&operator)

	// Benutzer nach der zweiten Zahl fragen
	fmt.Print("Gib die zweite Zahl ein: ")
	var num2 string
	fmt.Scanln(&num2)

	// Umwandlung der Eingaben in Fließkommazahlen
	num1Float, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		log.Fatal("Ungültige Zahl: ", err)
	}

	num2Float, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		log.Fatal("Ungültige Zahl: ", err)
	}

	var result float64

	// Die Berechnung je nach gewähltem Operator
	switch operator {
	case "+":
		result = num1Float + num2Float
	case "-":
		result = num1Float - num2Float
	case "*":
		result = num1Float * num2Float
	case "/":
		if num2Float == 0 {
			fmt.Println("Fehler: Division durch Null!")
			os.Exit(1)
		}
		result = num1Float / num2Float
	default:
		fmt.Println("Ungültiger Operator!")
		os.Exit(1)
	}

	// Ausgabe des Ergebnisses
	fmt.Printf("Das Ergebnis ist: %f\n", result)
}
