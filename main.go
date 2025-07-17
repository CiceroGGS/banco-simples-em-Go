package main

import (
	"banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoboleto float64) {
	conta.Sacar(valorDoboleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {

	contaCicero := contas.ContaCorrente{}

	println(contaCicero.Depositar(-100))
}
