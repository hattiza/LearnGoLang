package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err) // if we cannot open the keyboard package
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {

		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	/*
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("-> ")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.Replace(userInput, "\n", "", -1)
			userInput = strings.Replace(userInput, "\r\n", "", -1)

			if userInput == "quit" {
				break
			}
			fmt.Println(userInput)
		}
	*/
}
