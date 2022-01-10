package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func nilInitializer() interface{} {
	return nil
}

/**
 * Inicializa um Map de opções para o usuário escolher
 */
func initializeOpcoesMap() map[int]string {
	return map[int]string{
		0: "Sair",
		1: "Triângulo equilátero",
		2: "Retângulo",
		3: "Quadrado",
		4: "Circulo",
		5: "Pirâmide com base quadrangular",
		6: "Cubo",
		7: "Paralelepípedo",
		8: "Esfera",
	}
}

/**
 * Recupera inteiro do STDIN se for maior que 0
 */
func recuperarDado() int {
	// enquanto err != nil, tem que pegar o raio novamente
	stdin := bufio.NewReader(os.Stdin)

	var erroDadoInvalido = nilInitializer()
	var dado = -1
	for {
		_, erroDadoInvalido = fmt.Scanf("%d", &dado)
		stdin.ReadString('\n')

		if erroDadoInvalido == nil && dado > 0 {
			return dado
		}

		fmt.Println(NEGATIVO)
	}
}

/**
 * Imprime os dados calculados da figura
 */
func imprimirDadosCalculados(figura string, area float64, perimetro float64, volume float64) {
	if area > 0 {
		fmt.Printf("Área do(a) %s: %f\n", figura, area)
	}

	if perimetro > 0 {
		fmt.Printf("Perímetro do(a) %s: %f\n", figura, perimetro)
	}

	if volume > 0 {
		fmt.Printf("Volume do(a) %s: %f\n", figura, volume)
	}

	fmt.Println()
}

// TODO implementar cálculo de área da pirâmide
/**
 * Calcula a área das figuras geométricas planas/espaciais
 */
func calcularArea(figuraGeometrica int, base float64, altura float64, largura float64) float64 {
	if isCirculo(figuraGeometrica) {
		return math.Pi * math.Pow(base, 2)
	}

	if isEsfera(figuraGeometrica) {
		return 4 * math.Pow(base, 2) * math.Pi
	}

	if isQuadrado(figuraGeometrica) {
		return math.Pow(base, 2)
	}

	if isRetangulo(figuraGeometrica) {
		return base * altura
	}

	if isTriangulo(figuraGeometrica) {
		return base * altura / 2
	}

	if isCubo(figuraGeometrica) {
		return 6 * math.Pow(base, 2)
	}

	if isParalelepipedo(figuraGeometrica) {
		return 2 * (base*altura + base*largura + altura*largura)
	}

	return 0
}

/**
 * Calcula o volume das figuras espaciais
 */
func calcularVolume(figuraGeometrica int, base float64, altura float64, largura float64) float64 {
	if isPlano(figuraGeometrica) {
		return 0
	}

	if isEsfera(figuraGeometrica) {
		return (4 * math.Pi * math.Pow(base, 3)) / 3
	}

	if isParalelepipedo(figuraGeometrica) {
		return base * altura * largura
	}

	if isPiramide(figuraGeometrica) { // Is pirâmide
		return base * largura * altura / 3
	}

	if isCubo(figuraGeometrica) {
		return math.Pow(base, 3)
	}

	return 0
}

/**
 * Calcula o volume das figuras planas
 */
func calcularPerimetro(figuraGeometrica int, base float64, altura float64) float64 {
	if isEspacial(figuraGeometrica) {
		return 0
	}

	if isQuadrado(figuraGeometrica) {
		return 4 * base
	}

	if isRetangulo(figuraGeometrica) {
		return 2 * (base + altura)
	}

	if isTriangulo(figuraGeometrica) {
		return 3 * base
	}

	if isCirculo(figuraGeometrica) {
		return 2 * math.Pi * base
	}

	return 0
}
