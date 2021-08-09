package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	generatedNumber := rand.Intn(max-min) + min
	fmt.Println("The new generated number is", generatedNumber)

	try := 0
	for {
		try++
		fmt.Println("Please guess a any number between 1 and 100")
		fmt.Prln("Please give me your guess:")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading your number, try again")
			return
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Not valid input, please enter positive number")
			return
		}

		fmt.Println("Your guess is", guess)

		if guess > generatedNumber {
			fmt.Println("You guess is higher")
		} else if guess < generatedNumber {
			fmt.Println("You guess is smaller")
		} else {
			fmt.Println("You are Genius!")
			break
		}

	}

}
