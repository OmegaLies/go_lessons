package main

import "fmt"

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int	
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name: "John",
			number: 123124123,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name: "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 2500,
		user: user{
			name: "Sally",
			number: 19055558790,
		},
	})
}