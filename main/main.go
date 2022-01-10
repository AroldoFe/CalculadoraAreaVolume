package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const OPCOES = "Calculadora de Geometria Plana e Espacial"

const OPCAO_INVALIDA = "Opção inválida. Por favor, atente-se as opções listadas!"

const OPCAO_SAIR = "(0) Sair"

const OPCAO_ENTRADA = "Digite a sua opção: "

const OPCAO_RAIO = "Digite o tamanho do raio do(a)"

const OPCAO_BASE = "Digite o tamanho da base do(a)"

const OPCAO_LARGURA = "Digite o tamanho da largura do(a)"

const OPCAO_ALTURA = "Digite o tamanho da altura do(a)"

const NEGATIVO = "Por favor, insira um dado maior que 0!"

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

func isPlano(opcao int) bool {
	return 1 <= opcao && opcao <= 4
}

func isEspacial(opcao int) bool {
	return 5 <= opcao && opcao <= 8
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	opcoesMap := initializeOpcoesMap()

	var opcaoEscolhida = -1
	for {
		// imprimir opcoes
		{
			fmt.Println(OPCOES)
			for i := 1; i < 9; i++ {
				fmt.Printf("(%d) %s\n", i, opcoesMap[i])
			}
			fmt.Println(OPCAO_SAIR)
		}

		// Recuperar opcao do usuário
		fmt.Print(OPCAO_ENTRADA)
		var _, err = fmt.Scanf("%d", &opcaoEscolhida)
		stdin.ReadString('\n')

		// Não conseguiu converter, então opcaoEscolhida é definida fora do intervalo aceitado
		if err != nil {
			opcaoEscolhida = -1
			fmt.Println(err)
		}

		// Fazer o que tem que fazer se número está no intervalo aceitado
		if opcaoEscolhida == 0 {
			fmt.Println("Obrigado!")
			break
		} else if opcaoEscolhida > 0 && opcaoEscolhida < 9 {
			figura := opcoesMap[opcaoEscolhida]
			var base = -1
			var largura = -1
			var altura = -1
			var raio = -1

			var area = -1.0
			var perimetro = -1.0
			var volume = -1.0

			var erroDadoInvalido = nilInitializer()

			// Só precisa do raio
			if isCirculoOrEsfera(opcaoEscolhida) {
				fmt.Printf("%s %s: ", OPCAO_RAIO, figura)

				// enquanto err != nil, tem que pegar o raio novamente
				for {
					_, erroDadoInvalido = fmt.Scanf("%d", &raio)
					stdin.ReadString('\n')

					if erroDadoInvalido == nil && raio > 0 {
						break
					}

					fmt.Println(NEGATIVO)
				}

				// Círculo
				if isCirculo(opcaoEscolhida) {
					area = math.Pi * math.Pow(float64(raio), 2)
					perimetro = float64(2*raio) * math.Pi
				} else { // Esfera
					area = 4 * math.Pow(float64(raio), 2) * math.Pi
					volume = (4 * math.Pi * math.Pow(float64(raio), 3)) / 3
				}
			} else {

				// Recuperando base
				fmt.Printf("%s %s: ", OPCAO_BASE, figura)

				for {
					_, erroDadoInvalido = fmt.Scanf("%d", &base)
					stdin.ReadString('\n')

					if erroDadoInvalido == nil && base > 0 {
						break
					}

					fmt.Println(NEGATIVO)
				}

				if !isTriangulo(opcaoEscolhida) && !isQuadrado(opcaoEscolhida) && !isCubo(opcaoEscolhida) { // Não é quadrado, nem cubo
					// Recuperando altura
					fmt.Printf("%s %s: ", OPCAO_ALTURA, figura)

					for {
						_, erroDadoInvalido = fmt.Scanf("%d", &altura)
						stdin.ReadString('\n')

						if erroDadoInvalido == nil && altura > 0 {
							break
						}

						fmt.Println(NEGATIVO)
					}
				} else if isTriangulo(opcaoEscolhida) {
					altura = base * int(math.Sqrt(3)) / 2
				} else {
					// Para um quadrado, a altura é igual a base
					altura = base
					// Para um cubo: altura, largura e base são iguais
					largura = base
				}

				// Recuperando largura, se necessário (Volume)
				if isPiramide(opcaoEscolhida) || isParalelepipedo(opcaoEscolhida) {
					fmt.Printf("%s %s: ", OPCAO_LARGURA, figura)

					for {
						_, erroDadoInvalido = fmt.Scanf("%d", &largura)
						stdin.ReadString('\n')

						if erroDadoInvalido == nil && largura > 0 {
							break
						}

						fmt.Println(NEGATIVO)
					}
				}

				// Calcular o que é preciso

				if isPlano(opcaoEscolhida) {
					area = float64(base * altura)
					perimetro = float64(2*base + 2*altura)
					// Se é triângulo, remove uma base, pois 2*base + 2*altura = 4*base = 4*lado
					// Mas P(Triangulo) = 3*lado

					if isTriangulo(opcaoEscolhida) {
						perimetro = float64(3 * base)
					}
				} else if isParalelepipedo(opcaoEscolhida) || isCubo(opcaoEscolhida) {
					volume = float64(base * altura * largura)
					area = float64(2 * (base*altura + base*largura + largura*altura))
				} else { // Pirâmide
					// TODO Calcular area lateral da pirâmide
					// Como calcular a altura da face do triângulo tendo base, largura e altura da pirâmide

					volume = float64(base*largura*altura) / 3
				}

			}

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
			// De 1 a 3 preciso de base x altura (1 -> /2)
			// 4 Preciso de raio
			// de 5 a 7 preciso de base x altura x largura (5 -> / 3)
			// de 8 preciso de r
		} else {
			fmt.Println(OPCAO_INVALIDA)
		}
	}
}
