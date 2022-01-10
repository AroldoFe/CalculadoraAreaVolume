package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	// Opções de figuras geométricas que o usuário poderá escolher
	opcoesMap := initializeOpcoesMap()

	var figuraGeometricaSelecionada = -1
	for {
		// imprimir opcoes
		{
			fmt.Println(OPCOES)
			for i := 1; i < 9; i++ {
				fmt.Printf("(%d) %s\n", i, opcoesMap[i])
			}
		}

		// Recuperar opcao do usuário
		fmt.Print(OPCAO_ENTRADA)
		var _, err = fmt.Scanf("%d", &figuraGeometricaSelecionada)
		stdin.ReadString('\n')

		// Não conseguiu converter, então figuraGeometricaSelecionada é definida fora do intervalo aceitado
		if err != nil {
			figuraGeometricaSelecionada = -1
			fmt.Println(err)
		}

		// Resolveu sair
		if figuraGeometricaSelecionada == 0 {
			fmt.Println("Obrigado!")
			break
		} else if figuraGeometricaSelecionada > 0 && figuraGeometricaSelecionada < 9 {
			figura := opcoesMap[figuraGeometricaSelecionada]

			var base = -1
			var largura = -1
			var altura = -1

			// Só precisa do raio
			if isCirculoOrEsfera(figuraGeometricaSelecionada) {
				fmt.Printf("%s %s: ", OPCAO_RAIO, figura)
				base = recuperarDado()
			} else {
				// Se é quadrado ou cubo, só requer base
				// Recuperando base
				fmt.Printf("%s %s: ", OPCAO_BASE, figura)

				base = recuperarDado()

				// Se é triângulo, retângulo, pirâmide ou paralelepípedo, requer também altura
				if !isQuadradoOrCubo(figuraGeometricaSelecionada) {
					fmt.Printf("%s %s: ", OPCAO_ALTURA, figura)

					altura = recuperarDado()

					// Se é pirâmide ou paralelepíledo, requer também largura
					if isPiramideOrParalelepipedo(figuraGeometricaSelecionada) {
						fmt.Printf("%s %s: ", OPCAO_LARGURA, figura)

						largura = recuperarDado()
					}
				}

			}

			// Calcular dados
			area := calcularArea(figuraGeometricaSelecionada, float64(base), float64(altura), float64(largura))
			perimetro := calcularPerimetro(figuraGeometricaSelecionada, float64(base), float64(altura))
			volume := calcularVolume(figuraGeometricaSelecionada, float64(base), float64(altura), float64(largura))

			// Imprimindo dados
			imprimirDadosCalculados(figura, area, perimetro, volume)
		} else {
			// Opção inválida
			fmt.Println(OPCAO_INVALIDA)
		}
	}
}
