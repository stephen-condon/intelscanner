package conf

import "testing"

func TestRead(t *testing.T) {
	type test struct {
		filename       string
		expectedFolder string
		expectedSide   string
		expectedError  string
	}
	tests := []test{
		{filename: "./test/1.conf", expectedFolder: "test", expectedSide: "xyz", expectedError: ""},
		{filename: "./test/2.conf", expectedFolder: "test", expectedSide: "", expectedError: ""},
		{filename: "./test/3.conf", expectedFolder: "", expectedSide: "", expectedError: ""},
		{filename: "./test/4.conf", expectedFolder: "", expectedSide: "", expectedError: ""},
		{filename: "./test/5.conf", expectedFolder: "test", expectedSide: "xyz", expectedError: ""},
		{filename: "./test/6.conf", expectedFolder: "test", expectedSide: "x", expectedError: ""},
		{filename: "./test/not_a_file.conf", expectedFolder: "", expectedSide: "", expectedError: "open ./test/not_a_file.conf: no such file or directory"},
	}
	for _, tc := range tests {
		config, err := Read(tc.filename)
		if err != nil && len(tc.expectedError) != 0 {
			if err.Error() != tc.expectedError {
				t.Errorf(`Received %v, expected %v`, err.Error(), tc.expectedError)
			}
		} else if err == nil && len(tc.expectedError) != 0 {
			t.Errorf(`Received no error, expected %v`, tc.expectedError)
		} else if err != nil && len(tc.expectedError) == 0 {
			t.Errorf(`Received error: %v, expected no error`, err.Error())
		} else {
			if config.Folder != tc.expectedFolder {
				t.Errorf(`Expected %v, received %v for Folder`, tc.expectedFolder, config.Folder)
			}
			if config.Side != tc.expectedSide {
				t.Errorf(`Expected %v, received %v for Side`, tc.expectedSide, config.Side)
			}
		}
	}
}
