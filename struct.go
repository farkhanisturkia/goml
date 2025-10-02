package goml

type ModelType string

const (
	FMamdani   ModelType = "Mamdani"
	FTsukamoto ModelType = "Tsukamoto"
)

type Models struct {
	Model ModelType `json:"model"`
}

type Datas struct {
    Input      map[string]float64   `json:"Input"`
    Dataset    []map[string][]float64 `json:"Dataset"`
}

type Declaration struct {
	Min      float64   `json:"Min"`
	Mid      float64   `json:"Mid"`
	Max      float64   `json:"Max"`
}

type Fuzzification struct {
	Plus      float64   `json:"Plus"`
	Normal    float64   `json:"Normal"`
	Minus      float64   `json:"Minus"`
}

type Mamdani struct {
	Declarations    []map[string]Declaration `json:"Declaration"`
	Fuzzifications   []map[string]Fuzzification `json:"Fuzzification"`
}