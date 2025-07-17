package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                   clientes.Titular
	NumeroAgecia, NumeroConta int
	saldo                     float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque >= 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado!!!"
	} else {
		return "Nada de Saque!!!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Pode depositar", c.saldo
	} else {
		return "Pode n", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferecia float64, contaDestino *ContaCorrente) bool {
	if valorTransferecia <= c.saldo && valorTransferecia > 0 {
		c.saldo -= valorTransferecia
		contaDestino.saldo += valorTransferecia
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
