// Um professor de ensno médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água de Kelvin para Graus Celsius.
package main

import "fmt"


const Kelvin float64 = 373.15

// Função para converter Kelvin para Celsius
func ConvertKelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return celsius
}

// Função principal
	func main() {

	kelvin := Kelvin // Convertendo o ponto de ebulição da água de Kelvin para Celsius
	celsius := ConvertKelvinToCelsius(kelvin) // Ponto de ebulição da água em Celsius

	fmt.Printf("O ponto de ebulição da água em Kelvin é: %.2f K\n", kelvin)
	fmt.Printf("O ponto de ebulição da água em Graus Celsius é: %.2f °C\n", celsius)

	}