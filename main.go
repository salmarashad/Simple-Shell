package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var rootDirectory, err2 = os.Getwd()

func main (){
	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("> ", rootDirectory ," ")

		// read the input 
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		// Handle the execution of the input.
        if err = executeInput(input); err != nil {
            fmt.Println(os.Stderr, err)
        }

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
