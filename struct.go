package goml

type Datas struct {
    Input      map[string]float64   `json:"Input"`
    Fuzzification []map[string]map[string]float64 `json:"Fuzzification"`
}