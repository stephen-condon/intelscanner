package scan

import (
	"fmt"
	"regexp"
)

type GridLocation struct {
	X int
	Y int
}

type LineContents string

type ParsedLine struct {
	Location GridLocation
	Base     string
	Content  string
	Ship     string
	Turn     string
}

type RegexRules map[string]parseAdapter

type ParseConfig struct {
	Rules RegexRules
	Line  LineContents
	Turn  string
}

type parseAdapter func([][]string) (*ParsedLine, error)

func (pc *ParseConfig) parse() (*ParsedLine, error) {
	for rule, adapter := range pc.Rules {
		re := regexp.MustCompile(rule)
		submatches := re.FindAllStringSubmatch(string(pc.Line), -1)
		parsed, err := adapter(submatches)
		if err != nil {
			return nil, err
		}

		if parsed != nil {
			parsed.Turn = pc.Turn
			return parsed, nil
		}
	}

	return nil, nil
}

func readHeader(line string) (string, error) {
	re := regexp.MustCompile(`^SIG INT REPORT FOR ([[:alpha:]]{3} \d{2}, \d{2})$`)
	submatches := re.FindAllStringSubmatch(line, -1)
	if len(submatches) == 0 {
		return "", fmt.Errorf("no matches")
	}

	return submatches[0][1], nil
}

func (pl *ParsedLine) Display() string {
	var display string
	if len(pl.Base) > 0 {
		display = fmt.Sprintf(`%v: %v at %v (%v,%v)`, pl.Turn, pl.Content, pl.Base, pl.Location.X, pl.Location.Y)
	} else {
		display = fmt.Sprintf(`%v: %v at %v,%v`, pl.Turn, pl.Content, pl.Location.X, pl.Location.Y)
	}
	return display
}

func parseFile(contents []string, gameDate string, rules RegexRules) ([]*ParsedLine, error) {
	intelData := make([]*ParsedLine, 0)
	for _, line := range contents {
		parseConfig := ParseConfig{
			Rules: rules,
			Line:  LineContents(line),
			Turn:  gameDate,
		}

		data, err := parseConfig.parse()
		if err != nil {
			return []*ParsedLine{}, err
		}

		if data != nil {
			// fmt.Println(data.display())
			intelData = append(intelData, data)
		} else {
			fmt.Println("no match")
			fmt.Println(line)
		}
	}

	return intelData, nil
}

func Process(folder string, side string) ([]*ParsedLine, error) {
	intelData := make([]*ParsedLine, 0)
	scanConfig := ScanConfig{
		// Folder: "/Volumes/[C] Windows 11/Matrix Games/Witp - 1126b - Ironman/SAVE/archive",
		Folder: folder,
		Side:   side,
	}
	rules := buildAdapters()

	files, err := scanConfig.readDirectoryAndFilter()
	if err != nil {
		return []*ParsedLine{}, err
	}

	for _, file := range files {
		contents, err := scanConfig.readFile(file)
		if err != nil {
			return []*ParsedLine{}, err
		}

		gameDate, err := readHeader(contents[0])
		if err != nil {
			return []*ParsedLine{}, err
		}
		trimmedContents := contents[2:]

		data, err := parseFile(trimmedContents, gameDate, rules)
		if err != nil {
			return []*ParsedLine{}, err
		}
		intelData = append(intelData, data...)
	}

	return intelData, nil
}
