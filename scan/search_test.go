package scan

import (
	"reflect"
	"testing"
)

func TestNewSearcher(t *testing.T) {
	testData := []*ParsedLine{
		{GridLocation{1, 2}, "abc", "bcd", "cde", "Octember 1st"},
	}
	type test struct {
		data []*ParsedLine
	}
	tests := []test{
		{testData},
	}
	for _, tc := range tests {
		searcher := NewSearcher(tc.data)
		if !reflect.DeepEqual(searcher.Data, tc.data) {
			t.Errorf(`Expected %v, received %v`, tc.data, searcher.Data)
		}
	}
}

func TestSearchBase(t *testing.T) {
	testData := []*ParsedLine{
		{GridLocation{1, 2}, "my base", "bcd", "cde", "Octember 1st"},
		{GridLocation{1, 2}, "my base1", "bcd", "cde", "Octember 1st"},
	}
	type test struct {
		data            []*ParsedLine
		searchString    string
		expectedRecords int
	}
	tests := []test{
		{testData, "base", 2},
		{testData, "my", 2},
		{testData, "y b", 2},
		{testData, "se1", 1},
		{testData, "abc", 0},
	}
	for _, tc := range tests {
		searcher := NewSearcher(tc.data)
		result := searcher.SearchBase(tc.searchString)
		if len(result) != tc.expectedRecords {
			t.Errorf(`Expected %v record, received %v`, tc.expectedRecords, len(result))
		}
	}
}
