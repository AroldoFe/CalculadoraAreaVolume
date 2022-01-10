package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isTriangulo(opcao int) bool {
	return opcao == 1
}

func isRetangulo(opcao int) bool {
	return opcao == 2
}

func isQuadrado(opcao int) bool {
	return opcao == 3
}

func isCirculo(opcao int) bool {
	return opcao == 4
}

func isPiramide(opcao int) bool {
	return opcao == 5
}

func isCubo(opcao int) bool {
	return opcao == 6
}

func isParalelepipedo(opcao int) bool {
	return opcao == 7
}

func isEsfera(opcao int) bool {
	return opcao == 8
}

func isCirculoOrEsfera(opcao int) bool {
	return isCirculo(opcao) || isEsfera(opcao)
}

func isQuadradoOrCubo(opcao int) bool {
	return isQuadrado(opcao) || isCubo(opcao)
}

func isPiramideOrParalelepipedo(opcao int) bool {
	return isPiramide(opcao) || isParalelepipedo(opcao)
}

func isQuadrilatero(opcao int) bool {
	return isQuadrado(opcao) || isRetangulo(opcao)
}

func isPlano(opcao int) bool {
	return 1 <= opcao && opcao <= 4
}

func isEspacial(opcao int) bool {
	return 5 <= opcao && opcao <= 8
}

func nilInitializer() interface{} {
	return nil
}

func initializeOpcoesMap() map[int]string {
	return map[int]string{
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
func calcularArea(opcao int, base float64, altura float64, largura float64) float64 {
	if isCirculo(opcao) {
		return math.Pi * math.Pow(base, 2)
	} else if isEsfera(opcao) {
		return 4 * math.Pow(base, 2) * math.Pi
	} else if isQuadrado(opcao) {
		return math.Pow(base, 2)
	} else if isRetangulo(opcao) {
		return base * altura
	} else if isTriangulo(opcao) {
		return base * altura / 2
	} else if isCubo(opcao) {
		return 6 * math.Pow(base, 2)
	} else if isParalelepipedo(opcao) {
		return 2 * (base*altura + base*largura + altura*largura)
	} else {
		return 0
	}
}

func calcularVolume(opcao int, base float64, altura float64, largura float64) float64 {
	if isPlano(opcao) {
		return 0
	}

	if isEsfera(opcao) {
		return (4 * math.Pi * math.Pow(base, 3)) / 3
	} else if isParalelepipedo(opcao) {
		return base * altura * largura
	} else if isPiramide(opcao) { // Is pirâmide
		return base * largura * altura / 3
	} else if isCubo(opcao) {
		return math.Pow(base, 3)
	}

	return 0
}

func calcularPerimetro(opcao int, base float64, altura float64) float64 {
	if isEspacial(opcao) {
		return 0
	}

	if isQuadrado(opcao) {
		return 4 * base
	}

	if isRetangulo(opcao) {
		return 2 * (base + altura)
	}

	if isTriangulo(opcao) {
		return 3 * base
	}

	if isCirculo(opcao) {
		return 2 * math.Pi * base
	}

	return 0
}
