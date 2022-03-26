package main

import (
	"fmt"
	"github.com/jacobscunn07/aws-lens/internal/ui"
	"os"
)

func main() {
	if err := ui.New(); err != nil {
		fmt.Println("Oh no, it didn't work:", err)
		os.Exit(1)
	}
}
