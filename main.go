package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

var noColor bool

var bluePrintf = color.New(color.FgBlue).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()

func unmarshalLine(line []byte) {
	var strMap map[string]interface{}

	json.Unmarshal(line, &strMap)
	printDivider()
	printMap(strMap, 0)
}

func printDivider() {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}
	divider := strings.Repeat("-", width)
	fmt.Printf("\n%s\n\n", divider)
}

func printMap(strMap map[string]interface{}, depth int) {
	// max depth of 255 so we don't get too wild
	if depth > 255 {
		return
	}

	indent := strings.Repeat("\t", depth)
	for k, v := range strMap {
		if v == nil {
			printValue(k, "null", indent)
		} else if reflect.TypeOf(v).Kind() != reflect.Map {
			printValue(k, v, indent)
		} else {
			printValue(k, "", indent)
			printMap(v.(map[string]interface{}), depth+1)
		}
	}
}

func printValue(k string, v interface{}, indent string) {
	fmt.Print(indent)
	bluePrintf("%s", k)
	fmt.Print(": ")
	greenPrintf("%v", v)
	fmt.Print("\n")
}

func main() {
	flag.BoolVar(&noColor, "no-color", false, "Supresses color output")
	flag.Parse()

	color.NoColor = noColor

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		unmarshalLine(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
