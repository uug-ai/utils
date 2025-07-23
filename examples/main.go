package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the value of the environment variable "MY_ENV_VAR"
	envVar := os.Getenv("MY_ENV_VAR")

	// Check if the environment variable is set
	if envVar == "" {
		fmt.Println("Environment variable MY_ENV_VAR is not set.")
	} else {
		fmt.Printf("The value of MY_ENV_VAR is: %s\n", envVar)
	}
}
