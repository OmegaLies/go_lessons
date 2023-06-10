package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	mail, isEmail := e.(email)
	number, isSms := e.(sms)
	if isEmail {
		return mail.toAdress, mail.cost()
	}
	if isSms {
		return number.toPhoneNumber, number.cost()
	}
	return  "", 0.0
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
	toAdress     string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}


func test(e expense) {
	adress, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", adress, cost)
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", adress, cost)
	default:
		fmt.Println("Report: Invalid expense\n")
	}
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
		toAdress:     "1@1.1",
	}
	test(e)
	s := sms{
		isSubscribed:  false,
		body:          "I want money back",
		toPhoneNumber: "28489323",
	}
	test(s)
	i := invalid{}
	test(i)
}
