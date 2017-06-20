package main

import "fmt"

type Pato interface {
	Quack()
}

type Galinha interface {
	Cisca()
}

type Animal struct {
	nome string
}

func (o *Animal) SetNome(nome string) {
	o.nome = nome
}

func (o *Animal) GetNome() string {
	return o.nome
}

func (o Animal) Quack() {
	fmt.Println("Quack! ", o.nome, " faz quack!")
}

func (o Animal) Cisca() {
	fmt.Println(o.nome, " também cisca \\\\////\\\\////")
}

func main() {
	gg := Animal{}
	gg.SetNome("Gertrudes")
	QueroVerGalinhaCiscar(gg)
	VouOuvirUmPatoNoLago(gg)
	fmt.Println("Qual meu nome mesmo? ", gg.GetNome())
}

func VouOuvirUmPatoNoLago(p Pato) {
	p.Quack()
}

func QueroVerGalinhaCiscar(g Galinha) {
	g.Cisca()
}
