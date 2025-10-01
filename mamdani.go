package goml

import "slices"

func UseFuzzification(d *Datas) Mamdani {
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

func UseMamdani(d *Datas) (Mamdani, error) {
	var mamdani Mamdani
	mamdani = UseFuzzification(d)

	return mamdani, nil
}
