package goml

import (
    "fmt"
	"encoding/json"
    "sort"
)

func ParseJSON(jsonInput []byte) (*Datas, error) {
    var datas Datas
    err := json.Unmarshal(jsonInput, &datas)
    if err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %v", err)
    }
    return &datas, nil
}

func ValidationJSON(d *Datas, models Models) error {
    switch models.Model {
	case FMamdani:
        inputKeys := make([]string, 0, len(d.Input))
        for key := range d.Input {
            inputKeys = append(inputKeys, key)
        }
        sort.Strings(inputKeys)
    
        if len(inputKeys) != len(d.Dataset) {
            return fmt.Errorf("mismatch in Input and Dataset: input keys length (%d) != dataset entries (%d)", len(inputKeys), len(d.Dataset))
        }
    
        for i, dataset := range d.Dataset {
            for datasetKey := range dataset {
                if datasetKey != inputKeys[i] {
                    return fmt.Errorf("mismatch in Input and Dataset: input key %s != dataset key %s at index %d", inputKeys[i], datasetKey, i)
                }
            }
        }
    case FTsukamoto:
		return fmt.Errorf("Tsukamoto validation not implemented yet")
	default:
		return fmt.Errorf("unknown model: %s", models.Model)
	}

    return nil
}