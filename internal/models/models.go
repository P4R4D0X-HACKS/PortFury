package models

import (
	"encoding/json"
	"fmt"
)

// Struct 

type ScanResult struct {
	Port int `json:"port"`
	Status string `json:"status"`
	Banner string `json:"banner,omitempty"`
}

type ScanResults []ScanResult

// Function

func (s ScanResults) ToString() string {
	var output string
	for _, result := range s {
		output += fmt.Sprintf("Port %d: %s\n", result.Port, result.Status)
		if result.Banner != "" {
			output += fmt.Sprintf("  %s\n", result.Banner)
		}
	}
	return output
}

// To Json 

func (s ScanResults) ToJSON() string {
	jsonOutput, _ := json.MarshalIndent(s, "", "  ")
	return string(jsonOutput)
}
