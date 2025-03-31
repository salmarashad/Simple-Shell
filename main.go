package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-tty"
)

var history []string
var historyIndex int

func main() {

	// to detect up and down arrows
	tty, err := tty.Open()
	if err != nil {
		log.Fatalf("Failed to open tty: %v", err)
	}
	defer tty.Close()

	// intilise history
	history = []string{}
	historyIndex = 0

	
	var currentInput string

	for {
		// get and display current directory
		rootDirectory, rootError := os.Getwd()
		if rootError != nil {
			log.Fatalf("Failed to get current directory: %v", rootError)
		}

		// clear the current line and reprint with current input
		fmt.Print("\r\033[K") // clear the line
		fmt.Print("> ", rootDirectory, " % ", currentInput)

		// read a single key press
		nextKey, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		// handle special keys
		switch nextKey {
			case 27: // arrow key start
			next, err := tty.ReadRune()
			if err != nil {
				continue
			}
		
			if next != 91 { // if the next character isn't '[', ignore
				continue
			}
		
			key, err := tty.ReadRune()
			if err != nil {
				continue
			}
		
			if key == 65 { // up arrow
				if historyIndex > 0 {
					historyIndex--
					currentInput = history[historyIndex]
				}
			} else if key == 66 { // down arrow
				if historyIndex < len(history)-1 {
					historyIndex++
					currentInput = history[historyIndex]
				} else if historyIndex == len(history)-1 {
					historyIndex++
					currentInput = ""
				}
			}
		case 13: // enter key
			fmt.Println() // move to next line
			if currentInput != "" {
				// add to history
				history = append(history, currentInput)
				historyIndex = len(history)

				// execjte the command
				err := executeInput(currentInput)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				currentInput = ""
			}
		case 8, 127: // backspace handling
			if len(currentInput) > 0 {
				currentInput = currentInput[:len(currentInput)-1]
			}
		default:
			currentInput += string(nextKey)
		}
	}
}

func executeInput(input string) error {
	// removing the new line character
	input = strings.TrimSuffix(input,"\n")

	args := strings.Split(input, " ")
	
	switch args[0] {
	case "cd":  

    	if len(args) < 2 {
        	return  errors.New("path required")
    	}

    	return os.Chdir(args[1])
	case "exit": 
		 os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}