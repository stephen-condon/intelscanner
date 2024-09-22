package cli

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type rawArguments []string

type Arguments struct {
	Operation  string
	BaseSearch string
	SearchType string
}

func Process() (*Arguments, error) {
	args, err := processCliArgs()
	if err != nil {
		return nil, err
	}
	// fmt.Println(args)

	return args, nil
}

func processCliArgs() (*Arguments, error) {
	var args rawArguments = os.Args[1:]
	var baseSearch string
	var searchType string
	// fmt.Println(args)

	if len(args) < 2 {
		return nil, fmt.Errorf("missing search string")
		// return nil, fmt.Errorf("You must pass a search string")
	}

	operation := args[0]

	if !containsEquals(args[1]) {
		return nil, fmt.Errorf("pass properly formatted argument")
		// return nil, fmt.Errorf("You must pass a search formatted as {type}={value}")
	}
	values := strings.Split(args[1], "=")

	searchType = values[0]
	baseSearch = values[1]

	return &Arguments{
		Operation:  operation,
		BaseSearch: baseSearch,
		SearchType: searchType,
	}, nil
}

func containsEquals(s string) bool {
	re := regexp.MustCompile(`^\w+={1}\w+$`)
	result := re.FindAllStringIndex(s, -1)
	return len(result) > 0
}
