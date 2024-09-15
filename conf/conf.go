package conf

import (
	"bufio"
	"os"
	"strings"
)

type IntelscannerConfig struct {
	Folder string
	Side   string
}

type rawConfig []string

func Read(filename string) (*IntelscannerConfig, error) {

	// fullPath := fmt.Sprintf(`%v/%v`, sc.Folder, fileName)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var data rawConfig

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	config := data.parse()

	return config, nil
}

func (rc rawConfig) parse() *IntelscannerConfig {
	config := &IntelscannerConfig{}
	for _, line := range rc {
		splitLine := strings.Split(line, "=")
		if splitLine[0] == "folder" {
			config.Folder = splitLine[1]
		} else if splitLine[0] == "side" {
			config.Side = splitLine[1]
		}
	}

	return config
}
