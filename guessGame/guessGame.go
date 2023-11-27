package guessGame

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const prompt = "and do not type your number in, just press ENTER when ready."

func GuessGame() {
	// seed the random number generator
	// so it is actually random
	// rand.Seed(time.Now().UnixNano()) - no longer needed as of Go 1.20

	// Generate a random number between 1 and 100
	var firstNumber = rand.Intn(8) + 2 // from 2 to 10
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction
	play(firstNumber, secondNumber, substraction, answer)

}

func play(firstNumber, secondNumber, substraction, answer int) {
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

	fmt.Println("Divide the result by the number you originally thought off", prompt)
	reader.ReadString('\n')

	fmt.Fprintln(os.Stdout, "Now subtract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is:", answer, "!")
}
