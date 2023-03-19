package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name       string
		testNum    int
		wantNumber bool
		wantMsg    string
	}{
		{
			name:       "zero is not prime",
			testNum:    0,
			wantNumber: false,
			wantMsg:    "0 is not prime, by definition!",
		},
		{
			name:       "one is not prime",
			testNum:    1,
			wantNumber: false,
			wantMsg:    "1 is not prime, by definition!",
		},
		{
			name:       "negative number",
			testNum:    -1,
			wantNumber: false,
			wantMsg:    "Negative numbers are not prime, by definition!",
		},
		{
			name:       "prime number",
			testNum:    7,
			wantNumber: true,
			wantMsg:    "7 is a prime number",
		},
		{
			name:       "not prime number",
			testNum:    8,
			wantNumber: false,
			wantMsg:    "8 is not a prime number because it is divisible by 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotMsg := isPrime(tt.testNum)
			require.Equal(t, gotNum, tt.wantNumber)
			require.Equal(t, gotMsg, tt.wantMsg)
		})
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	require.Equal(t, "-> ", string(out))
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	require.Contains(t, string(out), "Enter a whole number")
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		wantMsg  string
		wantDone bool
	}{
		{
			name:     "success get prime number",
			in:       "7",
			wantMsg:  "7 is a prime number",
			wantDone: false,
		},
		{
			name:     "exit program",
			in:       "q",
			wantMsg:  "",
			wantDone: true,
		},
		{
			name:     "not a number",
			in:       "s",
			wantMsg:  "Please enter a whole number",
			wantDone: false,
		},
		{
			name:     "empty string",
			in:       "",
			wantMsg:  "Please enter a whole number",
			wantDone: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.in)
			reader := bufio.NewScanner(input)
			res, done := checkNumbers(reader)

			require.Equal(t, res, tt.wantMsg)
			require.Equal(t, done, tt.wantDone)
		})
	}
}

func Test_readUserInput(t *testing.T) {
	doneCh := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneCh)
	<-doneCh
	close(doneCh)

}
