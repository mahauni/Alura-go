package main

import (
	"fmt"
	c "projects/Alura-go/alura_OObj/contas"
)

func main() {
	conta1 := c.ContaCorrente{Titular: "Enrico", NumeroAgencia: 1, Saldo: 4.0}
	conta2 := c.ContaCorrente{Titular: "Lucas", NumeroAgencia: 12, NumeroConta: 12, Saldo: 69420}
	fmt.Println(conta1, conta2)

	var conta3 *c.ContaCorrente
	conta3 = new(c.ContaCorrente)
	conta3.Titular = "Davi"
	conta3.Saldo = 4.7

	fmt.Println(conta3, *conta3)

	fmt.Println(conta2.Sacar(420.69))
	fmt.Println(conta2)

	fmt.Println(conta1.Depositar(69))
	fmt.Println(conta1)

	fmt.Println(conta1.Transferir(69, conta3))
}
