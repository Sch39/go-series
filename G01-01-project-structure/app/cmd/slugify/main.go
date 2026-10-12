package main

import (
	"fmt"
	"os"

	"github.com/sch39/go-series/g0101/textkit"
)

func main() {
	inputs := os.Args[1:]
	for _, input := range inputs {
		fmt.Println(textkit.Slugify(input))
	}
}
