package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)


var history []string

func main (){
	reader := bufio.NewReader(os.Stdin)

	for{

		var rootDirectory, rootError = os.Getwd()

		if rootError != nil {
			log.Fatalf("Failed to get current directory: %v", rootError)
		}

		fmt.Print("> ", rootDirectory ," % ")

		// read the input 
		input, err := reader.ReadString('\n')

		// adding to my history of commands 
		history = append(history, input)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		// handle the execution of the input
        if err = executeInput(input); err != nil {
            fmt.Fprint(os.Stderr, err)
        }

		fmt.Println(history)

	}

}

func executeInput (input string ) error{
	
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
