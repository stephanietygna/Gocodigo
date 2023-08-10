package main

import (
	"fmt"
	"math"
)

const R = 6371 //raio da Terra em km

func Radianos(degress float64) float64 { //degress vai converter um ângulo em graus para radianos
	return degress * (math.Pi / 180.0) // multiplicar o ângulo em graus por pi/180
} // razão entre a circuferência de um círculo e seu diâmetro(pi) e 180 graus

func CoordenadasCartesianas(latitude, longitude float64) (x, y, z float64) {
	theta := Radianos(latitude) //teta
	phi := Radianos(longitude)  //fi

	x = R * math.Cos(theta) * math.Cos(phi)
	y = R * math.Cos(theta) * math.Sin(phi)
	z = R * math.Sin(theta)

	return x, y, z
}

func DistanciaEuclidiana(x1, y1, z1, x2, y2, z2 float64) float64 {

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2) + math.Pow(z2-z1, 2)) //math.Sqrt = calcular raiz quadrada
} //Pow = calcular potência de um número //"2" é o expoente que vai ser calculado antes da raiz quadrada

func main() {
	var latitudeA, longitudeA, latitudeB, longitudeB float64

	fmt.Println("Digite a latitude e a longitude do ponto A:") //pedindo a latitude e longitude do ponto A
	fmt.Print("Latitude A: ")
	fmt.Scan(&latitudeA)
	fmt.Print("Longitude A: ")
	fmt.Scan(&longitudeA)

	fmt.Println("Digite a latitude e a longitude do ponto B:") //pedindo a latitude e longitude do ponto A
	fmt.Print("Latitude B: ")
	fmt.Scan(&latitudeB)
	fmt.Print("Longitude B: ")
	fmt.Scan(&longitudeB)

	// converter as coordenadas geográficas em coordenadas cartesianas
	xA, yA, zA := CoordenadasCartesianas(latitudeA, longitudeA)
	xB, yB, zB := CoordenadasCartesianas(latitudeB, longitudeB)

	distancia := DistanciaEuclidiana(xA, yA, zA, xB, yB, zB) //calculando a distancia euclidiana entre os pontos A e B

	fmt.Printf("A distância euclidiana entre os pontos Ae B é: %.6f km\n", distancia)
}
