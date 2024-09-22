package cli

import (
	"os"
	"testing"
)

func TestProcess(t *testing.T) {
	type test struct {
		testArgs    []string
		expectedErr string
		expected    Arguments
	}
	tests := []test{
		{testArgs: []string{"program_name", "testB", "testC=testD"}, expectedErr: "", expected: Arguments{"testB", "testD", "testC"}},
		{testArgs: []string{"program_name", "testB", "testC"}, expectedErr: "pass properly formatted argument", expected: Arguments{}},
		{testArgs: []string{"program_name", "testB"}, expectedErr: "missing search string", expected: Arguments{}},
		{testArgs: []string{"program_name"}, expectedErr: "missing search string", expected: Arguments{}},
	}

	for _, tc := range tests {
		os.Args = tc.testArgs
		result, err := Process()
		if len(tc.expectedErr) > 0 {
			if err == nil {
				t.Fatalf(`Expected error %v, received no error`, tc.expectedErr)
			} else if tc.expectedErr != err.Error() {
				t.Fatalf(`Expected error %v, received %v`, tc.expectedErr, err.Error())
			}
		} else if len(tc.expectedErr) == 0 {
			if err != nil {
				t.Fatalf(`Expected no error, received %v`, err.Error())
			}
		}

		if result != nil {
			if tc.expected.Operation != result.Operation {
				t.Errorf(`Expected operation %v, received %v`, tc.expected.Operation, result.Operation)
			}
			if tc.expected.BaseSearch != result.BaseSearch {
				t.Errorf(`Expected baseSearch %v, received %v`, tc.expected.BaseSearch, result.BaseSearch)
			}
			if tc.expected.SearchType != result.SearchType {
				t.Errorf(`Expected searchType %v, received %v`, tc.expected.SearchType, result.SearchType)
			}
		}
	}
}
