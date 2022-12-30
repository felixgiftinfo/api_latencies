package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileAccess_getAPILatenciesData(t *testing.T) {
	type fields struct {
		PathToDirectory string
	}
	tests := []struct {
		name     string
		fields   fields
		minCount int
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS", fields: fields{PathToDirectory: "../files"}, minCount: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &FileAccess{
				PathToDirectory: tt.fields.PathToDirectory,
			}
			got := pd.getAPILatenciesData()
			assert.GreaterOrEqual(t, len(got), tt.minCount)

		})
	}
}

func TestFileAccess_getTransactionData(t *testing.T) {
	type fields struct {
		PathToDirectory string
	}
	tests := []struct {
		name     string
		fields   fields
		minCount int
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS", fields: fields{PathToDirectory: "../files"}, minCount: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &FileAccess{
				PathToDirectory: tt.fields.PathToDirectory,
			}
			got := pd.getTransactionData()
			assert.Len(t, got, tt.minCount)
		})
	}
}
