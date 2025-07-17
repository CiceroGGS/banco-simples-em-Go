package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                             clientes.Titular
	NumeroAgecia, NumeroConta, Operacao int
	saldo                               float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	if c.saldo >= valorDoSaque && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return "Pode sacar!", valorDoSaque
	} else {
		return "Nao pode sacar", valorDoSaque
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Valor depositado! -", valorDoDeposito
	} else {
		return "Nao pode depositar! - ", valorDoDeposito
	}
}

func (c *ContaPoupanca) Transferir(valorTransferecia float64, contaDestino *ContaPoupanca) (string, float64) {
	if c.saldo >= valorTransferecia {
		c.saldo -= valorTransferecia
		contaDestino.saldo += valorTransferecia
		return "Valor transferido! -", valorTransferecia
	} else {
		return "Impossivel transferir! -", valorTransferecia
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
