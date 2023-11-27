package guessGame

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func guessGame() {
	var firstNumber = 2
	var secondNumber = 5
	var substraction = 7
	var answer int

	var reader = bufio.NewReader(os.Stdin)

	// display a welcome/instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------\n")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')
	// take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply it by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the resault by the number you originally thought off", prompt)
	reader.ReadString('\n')

	fmt.Fprintln(os.Stdout, "Now subtract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstNumber*secondNumber - substraction
	fmt.Println("The answer is:", answer, "!")
}
