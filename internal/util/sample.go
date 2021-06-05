package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadSampleData loads sample data from a file
func LoadSampleData(path string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open a file %q: %w", path, err)
	}
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode a content of sample file: %w", err)
	}
	return data, nil
}

// Compare compares two sample sets
func Compare(sampleA, sampleB []map[string]interface{}) bool {
	if len(sampleA) != len(sampleB) {
		return false
	}
	dc := make(map[string]struct{})
	for _, m := range sampleA {
		b, _ := json.Marshal(m)
		dc[string(b)] = struct{}{}
	}
	for _, m := range sampleB {
		b, _ := json.Marshal(m)
		if _, ok := dc[string(b)]; !ok {
			return false
		}
	}
	return true
}
