package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"lnkparse/lnk"
	"os"
	"path/filepath"
	"strings"
)

const version string = "0.1.0"

var flags struct {
	Trim    bool
	Json    bool
	Version bool
}

func main() {
	Cmd()
}

func Cmd() {
	flag.BoolVar(&flags.Trim, "trim", false, "trim leading/trailing spaces in command line")
	flag.BoolVar(&flags.Json, "json", false, "show output in JSON format")
	flag.BoolVar(&flags.Version, "version", false, "show version and exit")
	flag.Usage = func() {
		fmt.Println(`Usage:
  lnkparse [flags] <PATH>

Examples:
  lnkparse C:\Users\user\Desktop\file.lnk
  lnkparse C:\Users\user\Desktop\*.lnk
  lnkparse C:\Users\user\**\*.lnk

Flags:
  -trim        trim leading/trailing spaces in command line
  -json        show output in JSON format
  -h, -help    show this message and exit
  -version     show version and exit`)
	}
	flag.Parse()

	if flags.Version {
		fmt.Printf("lnkparse %s\n", version)
		os.Exit(0)
	}

	var targetFiles []string

	for _, arg := range flag.CommandLine.Args() {
		matches, err := filepath.Glob(arg)
		if err != nil {
			fmt.Printf("Error: failed to find matches for %s\n", arg)
			continue
		}
		targetFiles = append(targetFiles, matches...)
	}

	if len(targetFiles) == 0 {
		fmt.Println("Error: no valid files were passed")
		os.Exit(1)
	}

	var lnks []*lnk.LnkFile

	for _, targetFile := range targetFiles {
		l, err := lnk.ParseFromFile(targetFile, flags.Trim)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		lnks = append(lnks, l)
	}

	if flags.Json {
		var jsonData []byte
		var err error
		if len(lnks) == 1 {
			jsonData, err = json.Marshal(lnks[0])
		} else {
			jsonData, err = json.Marshal(lnks)
		}
		if err != nil {
			fmt.Println("Error: failed to convert to JSON")
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	} else {
		for i, l := range lnks {
			fmt.Println(l)
			if i < len(lnks)-1 && len(lnks) != 1 {
				fmt.Println(strings.Repeat("-", 48))
			}
		}
	}

}
