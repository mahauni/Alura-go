package contas

import "github.com/mahauni/Alura-go/alura_OObj/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	if conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (conta *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "Deposito feito com sucesso", conta.saldo
	} else {
		return "Deposito não foi realizado", conta.saldo
	}
}

func (conta *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < conta.saldo && valorDaTransferencia > 0 {
		conta.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}
