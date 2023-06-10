package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + auth.username + ":" + auth.password
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("===========================")
}

func main() {
	test(authenticationInfo{
		username: "name",
		password: "password",
	})
	test(authenticationInfo{
		username: "John",
		password: "qwerty12345",
	})
}