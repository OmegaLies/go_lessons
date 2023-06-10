package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("======================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 12345678,
		message: "Thanks for signing up!",
	})
	test(messageToSend{
		phoneNumber: 2458883993,
		message: "Love to have you abroad!",
	})
	test(messageToSend{
		phoneNumber: 575792958250,
		message: "We're so excited!",
	})
}