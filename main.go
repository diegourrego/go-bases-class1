package main

import "fmt"

func main() {
	// Ejercicio #1:
	var name string = "Diego A. Urrego"
	var address string = "Calle 1 #2-3"
	fmt.Println("Mi nombre es: ", name)
	fmt.Println("Mi dirección es: ", address)

	// Ejercicio #2:
	var temperature float32 = 15
	var humidity float32 = 15
	var pressure float32 = 1030

	fmt.Printf("Clima Bogotá. Temperatura: %vºC, Humedad: %v%%, Presión: %v mb", temperature, humidity, pressure)

	// Ejercicio #3:
	var firstName string
	var lastName string
	var age uint8
	lastName = "Doe"
	var hasDriverLicense bool = true
	var personHeight float32
	childrenNumber := 2
	fmt.Println(firstName, lastName, age, hasDriverLicense, personHeight, childrenNumber)

	// Ejercicio #4:
	var lastName2 string = "Smith"
	var age2 uint8 = 35
	boolean := false
	var salary float32 = 45857.90
	var firstName2 string = "Mary"
	fmt.Println(lastName2, age2, boolean, salary, firstName2)
}
