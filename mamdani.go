package goml

import "slices"

func UseDeclarations(d *Datas) Mamdani {
	var mamdani Mamdani
	for _, dataset := range d.Dataset {
		for itemName, item := range dataset {
			mamdani.Declarations = append(mamdani.Declarations, map[string]Declaration{
				itemName: Declaration{
					Min: slices.Min(item),
					Mid: (slices.Min(item) + slices.Max(item)) / 2,
					Max: slices.Max(item),
				},
			})
		}
	}

	return mamdani
}

func UseFuzzification(d *Datas, m Mamdani) Mamdani {
	for inputName, inputValue := range d.Input {
		for _, declaration := range m.Declarations {
			for key, decValue := range declaration {
				if key == inputName {
					if inputValue <= decValue.Min {
						m.Fuzzifications = append(m.Fuzzifications, map[string]Fuzzification{
							inputName: Fuzzification{
								Minus: 1,
								Normal: 0,
								Plus: 0,
							},
						})
					} else if inputValue >= decValue.Min && inputValue < decValue.Mid {
						m.Fuzzifications = append(m.Fuzzifications, map[string]Fuzzification{
							inputName: Fuzzification{
								Minus: (decValue.Mid - inputValue) / (decValue.Mid - decValue.Min),
								Normal: (inputValue - decValue.Min) / (decValue.Mid - decValue.Min),
								Plus: 0,
							},
						})
					} else if inputValue >= decValue.Mid && inputValue < decValue.Max {
						m.Fuzzifications = append(m.Fuzzifications, map[string]Fuzzification{
							inputName: Fuzzification{
								Minus: 0,
								Normal: (decValue.Max - inputValue) / (decValue.Max - decValue.Mid),
								Plus:  (inputValue - decValue.Mid) / (decValue.Max - decValue.Mid),
							},
						})
					} else {
						m.Fuzzifications = append(m.Fuzzifications, map[string]Fuzzification{
							inputName: Fuzzification{
								Minus: 0,
								Normal: 0,
								Plus: 1,
							},
						})
					}
				}
			}
		}
	}

	return m
}

func UseMamdani(d *Datas) (Mamdani, error) {
	var mamdani Mamdani
	mamdani = UseDeclarations(d)
	mamdani = UseFuzzification(d, mamdani)

	return mamdani, nil
}
