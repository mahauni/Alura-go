package main

import (
	"fmt"

	"github.com/mahauni/Alura-go/alura_OObj/clientes"
	c "github.com/mahauni/Alura-go/alura_OObj/contas"
)

type verificarConta interface {
	Sacar(float64) string
}

func pagarBoleto(conta verificarConta, valorDoVoleto float64) {
	conta.Sacar(valorDoVoleto)
}

func main() {
	cliente1 := clientes.Titular{Nome: "Lucas", CPF: "111.111.111.11", Profissao: "Desenvolvedor"}

	conta1 := c.ContaCorrente{
		Titular:       cliente1,
		NumeroAgencia: 12,
		NumeroConta:   1,
	}

	conta1.Depositar(12000)

	pagarBoleto(&conta1, 420)

	fmt.Println(conta1)
	fmt.Println(conta1.ObterSaldo())

	conta2 := c.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Davi",
			CPF:       "122.122.122.12",
			Profissao: "Desempregado",
		},
		NumeroAgencia: 69,
		NumeroConta:   2,
	}

	conta2.Depositar(420)

	pagarBoleto(&conta2, 69)

	fmt.Println(conta2)
	fmt.Println(conta2.ObterSaldo())
}
