package main

import "fmt"

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  string
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your bithday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s"report is ready, You've sent %v reports`,sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("=====================================")

}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  "12",
	})
}