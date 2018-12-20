package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

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

	indent := strings.Repeat(config.indent, depth)
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
