package unit

import (
	"fmt"
	"maps"
	"slices"
)

type Unit int

const (
	G Unit = iota
	Kg
	Cas
	Cac
	L
	Ml
)

var DictName = map[Unit]string{
	G:   "G",
	Kg:  "Kg",
	Cas: "Cas",
	Cac: "Cac",
	L:   "L",
	Ml:  "Ml",
}

var (
	// DictToKg[ingrName][unit] returns conversion from unit to Kg for ingrName
	DictToKg = map[string]map[Unit]float64{
		"eau": map[Unit]float64{
			Ml:  0.001,
			Cas: 0.015},
		"farine": map[Unit]float64{
			Cas: 4.0 / 5.0 / 1000},
		"sel": map[Unit]float64{
			Cas: 15 / 1000},
		"moutarde": map[Unit]float64{
			Cas: 20 / 1000},
		"miel": map[Unit]float64{
			Cas: 21 / 1000},
		// "huile": map[Unit]float64{
		// 	Cas: / 1000},
	}

	// DictToVolume[FromUnit][ToUnit] returns volume conversion from FromUnit to ToUnit
	DictToVolume map[Unit]map[Unit]float64
)

func init() {
	DictToVolume = initDictToVolume()
}

func initDictToVolume() map[Unit]map[Unit]float64 {
	v := make(map[Unit]map[Unit]float64)
	for _, unit := range []Unit{Cas, Cac, Ml} {
		v[unit] = make(map[Unit]float64)
	}

	v[Cas][Cac] = 3 // 1 Cac = 3 Cac
	v[Cac][Ml] = 5  // 1 Cac = 1 Ml

	v[Cas][Ml] = v[Cas][Cac] * v[Cac][Ml]

	v[Cac][Cas] = 1 / v[Cas][Cac]
	v[Ml][Cas] = 1 / v[Cas][Ml]
	v[Ml][Cac] = 1 / v[Cac][Ml]

	return v
}

func ToKg(item string, unit Unit) (kg float64, err error) {
	// If unit is G, no conversion needed
	if unit == G {
		kg = 0.001
		return
	}
	if unit == Kg {
		kg = 1
		return
	}

	ConvertItem, itemExist := DictToKg[item]
	if itemExist {
		value, unitExist := ConvertItem[unit]
		// If unit exists, returns value
		if unitExist {
			kg = value
			return

			// if unit is "liquid", checks for existing other liquid conversion
		} else if liquidUnit := unit == Cac || unit == Cas || unit == Ml; liquidUnit {
			LiqUnits := slices.Collect(maps.Keys(DictToKg[item]))
			if len(LiqUnits) == 0 {
				err = fmt.Errorf("Unit %v for item %v does not exist in DictToKg", unit, item)
				return
			}

			unit2 := LiqUnits[0]
			value2 := DictToKg[item][unit2]

			kg = DictToVolume[unit][unit2] * value2
			return
		}
	} else {
		err = fmt.Errorf("Item %v does not exist in DictToKg", item)
		return
	}

	return 0, fmt.Errorf("??? Unexpected error ????")
}
