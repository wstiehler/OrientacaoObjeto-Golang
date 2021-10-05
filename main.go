package main

import (
	"fmt"

	"github.com/wstiehler/OrientacaoObjeto-Golang/contas"
	"github.com/wstiehler/OrientacaoObjeto-Golang/clientes"

)

func main() {
	contaDoFillipi := contas.ContaCorrente{Titular: clientes.Titular{
		Nome : "Fillipi",
		CPF: "123.456.789-12",
		Profissao: "Estudante",}
	NumeroAgencia:123, NumeroConta:123, Saldo:100}

	fmt.Println(contaDoFillipi)
}

