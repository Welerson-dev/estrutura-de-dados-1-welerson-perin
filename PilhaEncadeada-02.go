package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type PilhaEncadeada struct {
	topo    *No
	tamanho int
}

func (p *PilhaEncadeada) Empilhar(v int) {
	novo := No{valor: v, proximo: p.topo}
	p.topo = &novo
	p.tamanho++
}

func (p *PilhaEncadeada) Desempilhar() (int, bool) {
	if p.Vazia() {
		return 0, false
	}

	valor := p.topo.valor
	p.topo = p.topo.proximo
	p.tamanho--

	return valor, true
}

func (p *PilhaEncadeada) Topo() (int, bool) {
	if p.Vazia() {
		return 0, false
	}

	return p.topo.valor, true
}

func (p *PilhaEncadeada) Vazia() bool {
	return p.topo == nil
}

func (p *PilhaEncadeada) Tamanho() int {
	return p.tamanho
}

func main() {
	var p PilhaEncadeada

	p.Empilhar(1)
	p.Empilhar(2)
	p.Empilhar(3)

	valor, ok := p.Topo()

	fmt.Println("Tamanho:", p.Tamanho())

	if ok {
		fmt.Println("Topo:", valor)
	}

	fmt.Print("Desempilhados: ")

	for p.Vazia() == false {
		valor, ok := p.Desempilhar()

		if ok {
			fmt.Print(valor, " ")
		}
	}

	fmt.Println()
	fmt.Println("Tamanho após esvaziar:", p.Tamanho())

	p.Empilhar(10)
	p.Empilhar(20)
	p.Empilhar(30)
	p.Empilhar(40)
	p.Empilhar(50)

	fmt.Print("Desempilhados novamente: ")

	for p.Vazia() == false {
		valor, ok := p.Desempilhar()

		if ok {
			fmt.Print(valor, " ")
		}
	}

	fmt.Println()

	valor, ok = p.Desempilhar()

	fmt.Println("Valor:", valor)
	fmt.Println("Operação válida:", ok)
	fmt.Println("Tamanho final:", p.Tamanho())
}
