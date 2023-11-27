package saySomething

import (
	"bufio"
	"fmt"
	"main/doctor"
	"os"
	"strings"
)

func SaySomething() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	say(whatToSay)

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

func say(something string) {
	fmt.Println(something)
}
