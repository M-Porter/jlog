package main

import (
	"encoding/json"
)

func handleLogLine(line []byte) {
	var strMap map[string]interface{}

	err := json.Unmarshal(line, &strMap)
	if err != nil {
		panic(err)
	}

	printDivider()
	printMap(strMap, 0)
}
