package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) sacar(valor float64) {
	if conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		fmt.Println("Saque realizado com sucesso")
	} else {
		fmt.Println("Saldo insuficiente")
	}
}

func (conta *ContaCorrente) depositar(valor float64) {
	if valor > 0 {
		conta.saldo += valor
		fmt.Println("Deposito feito com sucesso")
	} else {
		fmt.Println("Deposito não foi realizado")
	}
}

func (conta *ContaCorrente) transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < conta.saldo && valorDaTransferencia > 0 {
		conta.saldo -= valorDaTransferencia
		contaDestino.depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	conta1 := ContaCorrente{titular: "Enrico", numeroAgencia: 1, saldo: 4.0}
	conta2 := ContaCorrente{"Lucas", 12, 12, 69420}
	fmt.Println(conta1, conta2)

	var conta3 *ContaCorrente
	conta3 = new(ContaCorrente)
	conta3.titular = "Davi"
	conta3.saldo = 4.7

	fmt.Println(conta3, *conta3)

	conta2.sacar(420.69)
	fmt.Println(conta2)

	conta1.depositar(69)
	fmt.Println(conta1)
}
