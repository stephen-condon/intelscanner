package main

import (
	"fmt"
	"intel-scanner/cli"
	"intel-scanner/conf"
	"intel-scanner/scan"
	"os"
)

func main() {
	args, err := cli.Process()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	filteredValues := searcher.SearchBase(args.BaseSearch)

	for _, value := range filteredValues {
		fmt.Println(value.Display())
	}
}
