func calcularPreco(qtd int, estudante bool) float64 {
	preco := 1000.0
	if estudante {
		preco = 500.0
	}
	return preco * float64(qtd)
}
