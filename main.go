package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
)

var config = new(struct {
	noColor bool
	indent  string
})

var bluePrintf = color.New(color.FgBlue).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()

func main() {
	flag.BoolVar(&config.noColor, "no-color", false, "Supresses color output")
	flag.StringVar(&config.indent, "indent", "    ", "Sequence to use as indentation of each level")
	flag.Parse()

	color.NoColor = config.noColor

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleLogLine(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
