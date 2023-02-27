package contas

import "github.com/mahauni/Alura-go/alura_OObj/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(valor float64) string {
	if conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (conta *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "Deposito feito com sucesso", conta.saldo
	} else {
		return "Deposito não foi realizado", conta.saldo
	}
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}
