package csv

import "testing"

func TestWriteCsvFile(t *testing.T) {
	type args struct {
		fileName string
		row      []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "aaa", args: args{
			fileName: "./data.csv",
			row:      []string{"1", "2", "3", "4", "5"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteCsvFile(tt.args.fileName, tt.args.row)
		})
	}
}
