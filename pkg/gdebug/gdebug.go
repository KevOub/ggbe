package gdebug

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

//OpcodesNoPrefix stores opcodes without prefixes
var OpcodesNoPrefix map[string]json.RawMessage

//OpcodesPrefixed stores opcodes with prefixes
var OpcodesPrefixed map[string]json.RawMessage

// below embeds the big json files into the binary...

//go:embed opcodes0.json
var F0 embed.FS

//go:embed opcodes1.json
var F1 embed.FS

// It is a debugger. Should be big and clunky

func init() {
	file0, _ := F0.ReadFile("opcodes0.json")
	err := json.Unmarshal([]byte(file0), &OpcodesNoPrefix)
	if err != nil {
		log.Println(err)
	}

	file1, _ := F1.ReadFile("opcodes1.json")
	err = json.Unmarshal([]byte(file1), &OpcodesPrefixed)
	if err != nil {
		log.Println(err)
	}
}

//WhatIsThisCode looks up instruction and returns info. Useful for later
func WhatIsThisCode(num int, prefix bool) string {

	h := fmt.Sprintf("0x%02x", num)
	// int to hex string: hex string look up in map
	var lookup string
	if prefix {
		lookup = string(OpcodesPrefixed[h])
	} else {
		lookup = string(OpcodesNoPrefix[h])

	}

	// strip dumb characters with regex
	reg, err := regexp.Compile("[,+{}+\"+\\t+]")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(lookup, "")
	return "-----\n" + processedString
}
