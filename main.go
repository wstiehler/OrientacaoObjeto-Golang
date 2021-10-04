package main

import (
	"fmt"

	"github.com/wstiehler/OrientacaoObjeto-Golang/contas"
)

func main() {
	contaDoDario := contas.ContaCorrente{Titular: "Dario", Saldo: 400}
	contaDaJuce := contas.ContaCorrente{Titular: "Juce", Saldo: 100}

	status := contaDoDario.Transferir(200, &contaDaJuce)

	fmt.Println(status)
	fmt.Println(contaDoDario)
	fmt.Println(contaDaJuce)

}
