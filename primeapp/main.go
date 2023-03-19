package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneCh := make(chan bool)

	go readUserInput(os.Stdin, doneCh)

	<-doneCh

	close(doneCh)

	fmt.Println("Goodbye.")
}

func readUserInput(in io.Reader, doneCh chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneCh <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numbToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numbToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
