package main

import (
	"reflect"
	"testing"
)

func TestMergeCSVData(t *testing.T) {
	data1 := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "", "New York"},
		{"Bob", "30", ""},
	}

	data2 := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "25", ""},
		{"Bob", "", "Los Angeles"},
	}

	expected := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "25", "New York"},
		{"Bob", "30", "Los Angeles"},
	}

	result := mergeCSVData(data1, data2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}