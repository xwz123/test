package animal

import "fmt"

type Animaler interface {
	Run()
	Eate()
}

type Dog struct {
}

func (d Dog) Run() {
	fmt.Println("dog run with for lagger")
}

func (d Dog) Eate() {
	fmt.Println("eate dog")
}