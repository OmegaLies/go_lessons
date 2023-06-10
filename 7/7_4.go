package main

import "fmt"

func fizzbuzz() {
	for i := 0; i<=100; i++ {
		mod3 := i % 3
		mod5 := i % 5
		if mod3 == 0 && mod5 == 0{
			fmt.Println("fizzbuzz")
		} else if mod3 == 0 {
			fmt.Println("fizz")
		} else if mod5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
		
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
