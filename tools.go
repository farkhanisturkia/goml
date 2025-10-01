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

func ValidationJSON(d *Datas) error {
    inputKeys := make([]string, 0, len(d.Input))
    for key := range d.Input {
        inputKeys = append(inputKeys, key)
    }
    sort.Strings(inputKeys)

    if len(inputKeys) != len(d.Fuzzification) {
        return fmt.Errorf("mismatch in Input and Fuzzification: input keys length (%d) != fuzzification entries (%d)", len(inputKeys), len(d.Fuzzification))
    }

    for i, fuzz := range d.Fuzzification {
        if len(fuzz) != 1 {
            return fmt.Errorf("fuzzification entry %d should contain exactly one key, got %d", i, len(fuzz))
        }
        for fuzzKey := range fuzz {
            if fuzzKey != inputKeys[i] {
                return fmt.Errorf("mismatch in Input and Fuzzification: input key %s != fuzzification key %s at index %d", inputKeys[i], fuzzKey, i)
            }
        }
    }

    return nil
}