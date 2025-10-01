package goml

import (
    "fmt"
	"encoding/json"
)

func ParseJSON(jsonInput []byte) (*Datas, error) {
    var datas Datas
    err := json.Unmarshal(jsonInput, &datas)
    if err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %v", err)
    }
    return &datas, nil
}