package main

import (
	"bufio"
	"fmt"
	"main/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	saySomething(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		response := doctor.Response(userInput)
		fmt.Println(response)

	}
}

func saySomething(something string) {
	fmt.Println(something)
}
