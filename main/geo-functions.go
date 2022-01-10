package main

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

func isPlano(opcao int) bool {
	return 1 <= opcao && opcao <= 4
}

func isEspacial(opcao int) bool {
	return 5 <= opcao && opcao <= 8
}
