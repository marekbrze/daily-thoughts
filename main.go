package main

import (
	"flag"
	"fmt"

	"github.com/marekbrze/daily-thoughts/cli"
)

func main() {
	flag.Parse()
	value := flag.Arg(0)
	if value == "" {
		fmt.Println(cli.NoParams())
	}

	if value != "" {
		fmt.Println(cli.ParseText(value))
	}
	fmt.Println("Test")
}
