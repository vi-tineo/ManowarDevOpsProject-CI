package main

import "testing"

func TestCalcularPreco(t *testing.T) {
	tests := []struct {
		qtd       int
		estudante bool
		esperado  float64
	}{
		{2, true, 1000.0},   // 2 ingressos com desconto de estudante
		{1, false, 1000.0},  // 1 ingresso preço cheio
		{3, true, 1500.0},   // 3 ingressos com desconto
		{4, false, 4000.0},  // 4 ingressos preço cheio
	}

	for _, tt := range tests {
		total := calcularPreco(tt.qtd, tt.estudante)
		if total != tt.esperado {
			t.Errorf("calcularPreco(%d, %v): esperado %.2f, obtido %.2f",
				tt.qtd, tt.estudante, tt.esperado, total)
		}
	}
}
