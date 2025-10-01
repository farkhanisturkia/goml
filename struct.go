package goml

type Datas struct {
    Input      map[string]float64   `json:"Input"`
    Fuzzifikasi []map[string]map[string]float64 `json:"Fuzzifikasi"`
}