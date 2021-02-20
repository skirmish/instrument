package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	log.SetPrefix("instrumenter: ")
	log.SetOutput(os.Stderr)

	var logs strings.Builder
	log.SetOutput(&logs)

	args := os.Args[1:]

	// parse arguments
	// do stuff
	// finish

	log.Println("incoming args", args)

	err := execute(args)

	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(os.Stderr, &logs)
	if err != nil {
		os.Exit(5)
	}
}

// execute a command
func execute(args []string) error {
	path := args[0]
	args = args[1:]
	quotedArgs := fmt.Sprintf("%q", args)
	log.Printf("executing command `%s %s`", path, quotedArgs[1:len(quotedArgs)-1])
	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
