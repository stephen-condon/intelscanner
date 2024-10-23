package main

import (
	"fmt"
	"intel-scanner/cli"
	"intel-scanner/conf"
	"intel-scanner/scan"
	"os"
	"runtime/debug"
)

func main() {
	args, err := cli.Process()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.Version {
		// raw, err := os.ReadFile("./version")
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// fmt.Println(string(raw))
		// os.Exit(0)

		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("Unable to determine version information.")
			os.Exit(1)
		}

		if buildInfo.Main.Version != "" {
			fmt.Printf("Version: %s\n", buildInfo.Main.Version)
			os.Exit(0)
		} else {
			fmt.Println("Version: unknown")
			os.Exit(1)
		}
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
