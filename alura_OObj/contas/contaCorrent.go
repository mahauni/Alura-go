package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	if conta.Saldo >= valor && valor > 0 {
		conta.Saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (conta *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		conta.Saldo += valor
		return "Deposito feito com sucesso", conta.Saldo
	} else {
		return "Deposito não foi realizado", conta.Saldo
	}
}

func (conta *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < conta.Saldo && valorDaTransferencia > 0 {
		conta.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
