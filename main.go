package main

import (
	"bufio"
	"fmt"
	"main/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	saySomething(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		if userInput == "quit" {
			break
		}
		fmt.Println(doctor.Response(userInput))

	}
}

func saySomething(something string) {
	fmt.Println(something)
}
