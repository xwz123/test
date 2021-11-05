package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("hello github actions %v", "s")
	a, _ := mut(1, 0)
	fmt.Println(a)
}

func mut(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error")
	}
	return a / b, nil
}
