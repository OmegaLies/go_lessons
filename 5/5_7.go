package main

import (
	"fmt"
)


func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want money back",
	}
	test(e, e)
}

