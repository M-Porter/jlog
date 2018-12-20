package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
)

var noColor bool

var bluePrintf = color.New(color.FgBlue).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()

func main() {
	flag.BoolVar(&noColor, "no-color", false, "Supresses color output")
	flag.Parse()

	color.NoColor = noColor

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleLogLine(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
