package main

import (
	"fmt"
	"math"
)

func main() {

	var capacidadeTotal float64
	fmt.Println("Digite o tamano do tanque em litros:")
	fmt.Scanln(&capacidadeTotal)

	leiturasPercentuais := []float64{66.0, 66.0, 59.0, 59.0, 57.5, 57.5, 59.0, 59.0, 72.9, 72.9, 61.5, 61.5, 62.5, 62.5, 61.5, 61.5, 62.5, 62.5, 61.5, 61.5, 62.5, 58.0, 58.0, 61.2, 61.2}

	distanciaPercorrida := 10.0

	KalmanFilter(capacidadeTotal, distanciaPercorrida, leiturasPercentuais)

}

func KalmanFilter(capacidade float64, distancia float64, medições []float64) {

	var media = mediavector(medições)

	Gerro := errovector(medições, media) // erro global

	auxe := []float64{}
	auxe = append(auxe, medições[0])
	Lerro := errovector(auxe, media) // erro local

	var estima float64
	leiturasPercentuaispos := []float64{}

	for _, leitura := range medições {
		var k = math.Pow(Lerro, 2)/math.Pow(Lerro, 2) + math.Pow(Gerro, 2)
		estima = media + k*(leitura-media)

		Lerro = (1 - k) * Lerro
		leiturasPercentuaispos = append(leiturasPercentuaispos, estima)

	}
	fmt.Printf("\n")
	fmt.Printf("Medições pelo OBD %.2f", medições)
	fmt.Printf("\n")
	fmt.Printf("Medições corrigidas %.2f", leiturasPercentuaispos)
	fmt.Printf("\n")
	fmt.Printf("Estimativa final de restante de combustivel %.2f", estima)
	fmt.Printf("\n")

	resultadotanque := (medições[1] - estima) * capacidade

	eficienciaEnergetica := distancia / resultadotanque

	fmt.Printf("Eficiência Energética: %.2f km/l\n", eficienciaEnergetica)
}

func mediavector(a []float64) float64 {
	media := 0.0
	for _, leitura := range a {
		media = media + leitura
	}

	media = media / float64(len(a))
	return media

}

func errovector(a []float64, m float64) float64 {
	var aux float64
	errovec := []float64{}
	for _, leitura := range a {
		errovec = append(errovec, leitura-m)
	}
	for _, leitura := range errovec {
		aux = +leitura
	}
	aux = aux / float64(len(a))
	return aux
}

func filtro(a []float64) []float64 {

	return []float64{66.0}

}
