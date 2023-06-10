package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	switch tmp := e.(type) {
	case email:
		return tmp.toAddress, tmp.cost()
	case sms:
		return tmp.toPhoneNumber, tmp.cost()
	default:
		return  "", 0.0
	}
}


func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	}
	return float64(len(e.body)) * 0.05
}

func (s sms) cost() float64 {
	if s.isSubscribed {
		return float64(len(s.body)) * 0.03
	}
	return float64(len(s.body)) * 0.1
}

func (i invalid) cost() float64 {
	return 0.0
}

type invalid struct {}

type expense interface {
	cost() float64
}


type email struct {
	isSubscribed bool
	body         string
	toAddress     string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}


func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Wanna catch up later?",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "I don't need this",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}