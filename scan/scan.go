package scan

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type ScanConfig struct {
	Folder string
	Side   string
}

func (sc *ScanConfig) readDirectoryAndFilter() ([]string, error) {
	scanFilter := `^jsigint_.+.txt$`
	if sc.Side == "Allies" {
		scanFilter = `^asigint_.+.txt$`
	}
	filteredFileNames := []string{}
	files, err := os.ReadDir(sc.Folder)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		re, err := regexp.Compile(scanFilter)
		if err != nil {
			return nil, err
		}

		hasMatch := re.MatchString(file.Name())
		if hasMatch {
			filteredFileNames = append(filteredFileNames, file.Name())
		}
	}

	return filteredFileNames, nil
}

func (sc *ScanConfig) readFile(fileName string) ([]string, error) {
	fullPath := fmt.Sprintf(`%v/%v`, sc.Folder, fileName)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var data []string

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ^.+ moving to \w+\.$
// ^.+ Radio transmissions detected at \w+ \(\d+,\d+\)\.$
// ^.*Radio transmissions detected at \d+,\d+\.$
// ^.* is located at .+\(\d+,\d+\)\.$
// ^.* is loaded on a .+ moving to .+\.$
