package goml

type Datas struct {
    Input      map[string]float64   `json:"Input"`
    Dataset    []map[string][]float64 `json:"Dataset"`
}