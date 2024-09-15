package scan

import (
	"reflect"
	"testing"
)

func TestParseAndAdapters(t *testing.T) {
	testData := &ParsedLine{
		Location: GridLocation{X: 85, Y: 42},
		Base:     "",
		Content:  "Radio transmissions",
		Ship:     "",
	}
	rules := buildAdapters()

	config := ParseConfig{
		Rules: rules,
		Line:  "Radio transmissions detected at 85,42.",
	}

	data, err := config.parse()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(data, testData) {
		t.Error(data, testData)
	}
}

func TestReadHeader(t *testing.T) {
	type test struct {
		line        string
		expected    string
		expectedErr string
	}
	tests := []test{
		{"SIG INT REPORT FOR Dec 07, 41", "Dec 07, 41", ""},
		{"Dec 07, 41", "", "no matches"},
		{"bwrtbtbtenytn", "", "no matches"},
	}
	for _, tc := range tests {
		date, err := readHeader(tc.line)

		if date != tc.expected {
			t.Errorf("Received %v, expected %v", date, tc.expected)
		}

		if err != nil && len(tc.expectedErr) == 0 {
			t.Errorf("Did not expect error, received %v", err)
		} else if err == nil && len(tc.expectedErr) > 0 {
			t.Errorf("Expected error, did not receive one")
		} else if err != nil && err.Error() != tc.expectedErr {
			t.Errorf("Expected error %v, but did not match received error: %v", tc.expectedErr, err.Error())
		}
	}
}

func TestDisplay(t *testing.T) {
	type test struct {
		parsed   ParsedLine
		expected string
	}
	tests := []test{
		{ParsedLine{Location: GridLocation{X: 1, Y: 2}, Base: "", Content: "test content", Ship: "", Turn: "test turn"}, "test turn: test content at 1,2"},
		{ParsedLine{Location: GridLocation{X: 1, Y: 2}, Base: "", Content: "test content", Ship: "abcd", Turn: "test turn"}, "test turn: test content at 1,2"},
		{ParsedLine{Location: GridLocation{X: 1, Y: 2}, Base: "Gotham", Content: "test content", Ship: "", Turn: "test turn"}, "test turn: test content at Gotham (1,2)"},
		{ParsedLine{Location: GridLocation{X: 1, Y: 2}, Base: "Gotham", Content: "test content", Ship: "abcd", Turn: "test turn"}, "test turn: test content at Gotham (1,2)"},
	}
	for _, tc := range tests {
		displayedString := tc.parsed.display()

		if displayedString != tc.expected {
			t.Errorf("Received %v, expected %v", displayedString, tc.expected)
		}
	}
}

func TestParseFile(t *testing.T) {
	type test struct {
		contents    []string
		gameDate    string
		rules       RegexRules
		expectedErr string
	}
	realRules := buildAdapters()
	tests := []test{
		{[]string{"a", "b"}, "abcd", realRules, ""},
	}
	for _, tc := range tests {
		_, err := parseFile(tc.contents, tc.gameDate, tc.rules)

		if err != nil {
			if len(tc.expectedErr) > 0 {
				if err.Error() != tc.expectedErr {
					t.Errorf("mismatched errors")
				}
			} else {
				t.Errorf("Expected no error, recieved %v", err.Error())
			}
		} else {
			if len(tc.expectedErr) > 0 {
				t.Errorf("Expected error %v, received no error", tc.expectedErr)
			}
		}
	}
}

func TestProcess(t *testing.T) {
	type test struct {
		folder    string
		side      string
		expectErr bool
	}
	tests := []test{
		{folder: "./test", side: "Allies", expectErr: false},
		{folder: "./test", side: "Japan", expectErr: false},
		{folder: "../test", side: "Allies", expectErr: true},
		{folder: "../test", side: "Japan", expectErr: true},
	}

	for _, tc := range tests {
		_, err := Process(tc.folder, tc.side)
		if tc.expectErr && err == nil {
			t.Errorf("Expected error, recieved %v", err)
		} else if !tc.expectErr && err != nil {
			t.Errorf("Did not expect error, received %v", err)
		}
	}

}

func BenchmarkParseAndAdapters(b *testing.B) {
	rules := buildAdapters()
	config := ParseConfig{
		Rules: rules,
		Line:  "Radio transmissions detected at 85,42.",
	}
	for i := 0; i < b.N; i++ {
		config.parse()
	}
}

func BenchmarkReadHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readHeader("SIG INT REPORT FOR Dec 07, 41")
	}
}

func BenchmarkDisplay(b *testing.B) {
	obj := ParsedLine{Location: GridLocation{X: 1, Y: 2}, Base: "", Content: "test content", Ship: "", Turn: "test turn"}
	for i := 0; i < b.N; i++ {
		obj.display()
	}
}

func BenchmarkParseLine(b *testing.B) {
	realRules := buildAdapters()
	for i := 0; i < b.N; i++ {
		parseFile([]string{"a", "b"}, "abcd", realRules)
	}
}

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Process("./test", "Allies")
	}
}
