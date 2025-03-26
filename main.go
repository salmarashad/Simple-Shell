package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func main (){
	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("> ")

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

	
	cmd := exec.Command(input)
	
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}
