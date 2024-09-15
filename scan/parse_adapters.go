package scan

import (
	"fmt"
	"strconv"
)

/*
Adapters defined here, require regex that uses capture groups to identify important values.
Function defined of type parseAdapter that leverages important values and loads them into
a ParsedLine struct, with any needed error handling.
*/
func buildAdapters() RegexRules {
	rules := RegexRules{}
	rules[`^(Heavy Volume of Radio transmissions) detected at (\d+),(\d+)\.$`] = heavyRadioTransmissionsAdapter
	rules[`^.*(Radio transmissions) detected at (\d+),(\d+)\.$`] = radioTransmissionsAdapter
	rules[`^(Heavy Volume of Radio transmissions) detected at (.*) \((\d+),(\d+)\)\.$`] = heavyRadioTransmissionsAdapterBase
	rules[`^(Radio transmissions) detected at (.*) \((\d+),(\d+)\)\.$`] = radioTransmissionsAdapterBase
	rules[`^(.*) is located at (.*)\((\d+),(\d+)\)\.$`] = locatedAtAdapterBase
	rules[`^(.*) is loaded on a (.* .*) moving to (.*)\.$`] = loadedOnAdapterBase
	rules[`^(\d* men) are based at (.*) \((\d*),(\d*)\)\.$`] = menAtAdapterBase
	rules[`^(.*) is located at (\d*),(\d*)\.$`] = locatedAtAdapter
	rules[`^(.*) is loaded on (.*) moving to (.*)\.$`] = loadedOnAdapterBaseAlt
	rules[`^(.*) is planning for an attack on (.*)\.$`] = planningForAttackAdapter
	rules[`^a (.*) is moving to (\d*),(\d*)\.$`] = shipMovingToAdapter
	rules[`^a (.*) is moving to (.*) \((\d*),(\d*)\)\.$`] = shipMovingToAdapterBase
	rules[`^(\d* aircraft) are based at (.*) \((\d*),(\d*)\)\.$`] = aircraftAtAdapterBase
	rules[`^(Radio call sign) of (.*) detected at (\d*),(\d*)\.$`] = shipRadioCallSignAdapter
	rules[`^((.*) is moving) to (\d*),(\d*)\.$`] = shipMovingToAdapterAlt
	rules[`^((.*) is moving) to (.*) \((\d*),(\d*)\)\.$`] = shipMovingToAdapterBaseAlt
	rules[`^(.* is loaded) on (.*) at (.*) \((\d*),(\d*)\)\.$`] = shipAtAdapterBase

	return rules
}

func radioTransmissionsAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func heavyRadioTransmissionsAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func radioTransmissionsAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func heavyRadioTransmissionsAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func locatedAtAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func loadedOnAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		return &ParsedLine{
			Location: GridLocation{},
			Base:     submatches[0][3],
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}

func menAtAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func locatedAtAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}

		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func loadedOnAdapterBaseAlt(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		return &ParsedLine{
			Location: GridLocation{},
			Base:     submatches[0][3],
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}

func planningForAttackAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		return &ParsedLine{
			Location: GridLocation{},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func shipMovingToAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][2])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  fmt.Sprintf(`%v is moving`, submatches[0][1]),
			Ship:     submatches[0][1],
		}, nil
	}

	return nil, nil
}

func shipMovingToAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  fmt.Sprintf(`%v is moving`, submatches[0][1]),
			Ship:     submatches[0][1],
		}, nil
	}

	return nil, nil
}

func aircraftAtAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][2],
			Content:  submatches[0][1],
			Ship:     "",
		}, nil
	}

	return nil, nil
}

func shipRadioCallSignAdapter(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}

func shipMovingToAdapterAlt(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][3])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     "",
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}

func shipMovingToAdapterBaseAlt(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][5])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][3],
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}

func shipAtAdapterBase(submatches [][]string) (*ParsedLine, error) {
	if len(submatches) != 0 {
		x, err := strconv.Atoi(submatches[0][4])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(submatches[0][5])
		if err != nil {
			return nil, err
		}
		return &ParsedLine{
			Location: GridLocation{X: x, Y: y},
			Base:     submatches[0][3],
			Content:  submatches[0][1],
			Ship:     submatches[0][2],
		}, nil
	}

	return nil, nil
}
