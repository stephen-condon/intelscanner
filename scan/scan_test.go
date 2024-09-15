package scan

import (
	"testing"
)

func TestReadDirectoryAndFilter(t *testing.T) {
	type test struct {
		folder        string
		side          string
		firstExpected string
	}
	tests := []test{
		{folder: "/Volumes/[C] Windows 11/Matrix Games/Witp - 1126b - Ironman/SAVE/archive", side: "Allies", firstExpected: "asigint_411207.txt"},
		{folder: "/Volumes/[C] Windows 11/Matrix Games/Witp - 1126b - Ironman/SAVE/archive", side: "Japan", firstExpected: "jsigint_411207.txt"},
	}

	for _, tc := range tests {
		scanConfig := ScanConfig{
			Folder: tc.folder,
			Side:   tc.side,
		}

		files, err := scanConfig.readDirectoryAndFilter()
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(files) == 0 {
			t.Errorf("No files found")
		}

		if files[0] != tc.firstExpected {
			t.Errorf("Received %v, expected %v", files[0], tc.firstExpected)
		}
	}
}

func TestReadFile(t *testing.T) {
	type test struct {
		folder   string
		side     string
		filename string
	}
	tests := []test{
		{folder: "/Volumes/[C] Windows 11/Matrix Games/Witp - 1126b - Ironman/SAVE/archive", side: "Allies", filename: "asigint_411207.txt"},
		{folder: "/Volumes/[C] Windows 11/Matrix Games/Witp - 1126b - Ironman/SAVE/archive", side: "Japan", filename: "jsigint_411207.txt"},
	}

	for _, tc := range tests {
		scanConfig := ScanConfig{
			Folder: tc.folder,
			Side:   tc.side,
		}

		contents, err := scanConfig.readFile(tc.filename)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(contents) == 0 {
			t.Errorf("No files found")
		}

		for i, line := range contents {
			if i != 1 && len(line) == 0 {
				t.Errorf("SigInt report contains empty line")
			} else if i == 1 && len(line) > 0 {
				t.Errorf("Expected blank line, received %v", line)
			}
		}
	}
}
