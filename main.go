package main

import (
	"fmt"
	"main/doctor"
)

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	saySomething(whatToSay)
}

func saySomething(something string) {
	fmt.Println(something)
}
