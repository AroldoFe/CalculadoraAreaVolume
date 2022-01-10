package main

import (
	"bufio"
	"fmt"
	"os"
)

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

		// Resolveu sair
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

			// Só precisa do raio
			if isCirculoOrEsfera(opcaoEscolhida) {
				fmt.Printf("%s %s: ", OPCAO_RAIO, figura)
				raio = recuperarDado()
			} else {
				// Se é quadrado ou cubo, só requer base
				// Recuperando base
				fmt.Printf("%s %s: ", OPCAO_BASE, figura)

				base = recuperarDado()

				// Se é triângulo, retângulo, pirâmide ou paralelepípedo, requer também altura
				if !isQuadradoOrCubo(opcaoEscolhida) {
					fmt.Printf("%s %s: ", OPCAO_ALTURA, figura)

					altura = recuperarDado()

					// Se é pirâmide ou paralelepíledo, requer também largura
					if isPiramideOrParalelepipedo(opcaoEscolhida) {
						fmt.Printf("%s %s: ", OPCAO_LARGURA, figura)

						largura = recuperarDado()
					}
				}

			}

			// Calcular dados
			if isCirculoOrEsfera(opcaoEscolhida) {
				area = calcularArea(opcaoEscolhida, float64(raio), 0, 0)
				perimetro = calcularPerimetro(opcaoEscolhida, float64(raio), 0)
				volume = calcularVolume(opcaoEscolhida, float64(raio), 0, 0)
			} else {
				area = calcularArea(opcaoEscolhida, float64(base), float64(altura), float64(largura))
				perimetro = calcularPerimetro(opcaoEscolhida, float64(base), float64(altura))
				volume = calcularVolume(opcaoEscolhida, float64(base), float64(altura), float64(largura))
			}

			// Imprimindo dados
			imprimirDadosCalculados(figura, area, perimetro, volume)
		} else {
			// Opção inválida
			fmt.Println(OPCAO_INVALIDA)
		}
	}
}
