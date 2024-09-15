package main

import (
	"fmt"
	"intel-scanner/conf"
	"intel-scanner/scan"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must pass a search string")
		os.Exit(1)
	}
	searchString := args[0]

	config, err := conf.Read("./intelscanner.conf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	intelData, err := scan.Process(config.Folder, config.Side)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	searcher := scan.NewSearcher(intelData)
	filteredValues := searcher.SearchBase(searchString)

	for _, value := range filteredValues {
		fmt.Println(value)
	}
}
