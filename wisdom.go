package main

import (
	"./project"
	"github.com/fatih/color"
)

func main() {
	colored := color.New(color.FgWhite, color.Bold)
	colored.Printf(project.PrintRandomTweet())
	colored.Printf("\n")
}
