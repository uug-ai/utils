package main

import (
	"fmt"
	"io"
	"os"
)

func printEnvVar(name string, w io.Writer) {
	v := os.Getenv(name)
	if v == "" {
		fmt.Fprintln(w, "Environment variable "+name+" is not set.")
		return
	}
	fmt.Fprintf(w, "The value of %s is: %s\n", name, v)
}

func main() {
	printEnvVar("MY_ENV_VAR", os.Stdout)
}
