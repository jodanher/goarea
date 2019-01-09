package goarea

import "math"

// PI é razão entre a circunferência de um círculo e seu diâmetro
const PI = 3.141516

// Circ calcula a área da circunferência
func Circ(raio float64) float64 {
	return PI * math.Pow(raio, 2)
}

// Rect calcula a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// função privada (visível apenas no pacote)
func _TrianguloEquilatero(base, altura float64) float64 {
	return Rect(base, altura) / 2
}
