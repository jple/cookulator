package main

import (
	"fmt"
	"log/slog"
	"maps"
	"slices"
)

type (
	// IngredientBase is the basic type struct for all ingredients (solid, liquid, ...)
	IngredientBase struct {
		Name     string
		Quantity float64
		Unit     Unit
		UnitName string
	}
	Solid  IngredientBase
	Farine Solid

	Ingredients []IngredientBase
)

var (
	// ConvertKg[ingrName][unit] returns conversion from unit to Kg for ingrName
	ConvertKg = map[string]map[Unit]float64{
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

	// ConvertVol[FromUnit][ToUnit] returns volume conversion from FromUnit to ToUnit
	ConvertVol map[Unit]map[Unit]float64
)

func init() {
	ConvertVol = initConvertVol()
}

func initConvertVol() map[Unit]map[Unit]float64 {
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

func GetInKg(item string, unit Unit) (kg float64, err error) {
	// If unit is G, no conversion needed
	if unit == G {
		kg = 0.001
		return
	}

	ConvertItem, itemExist := ConvertKg[item]
	if itemExist {
		value, unitExist := ConvertItem[unit]
		// If unit exists, returns value
		if unitExist {
			kg = value
			return

			// if unit is "liquid", checks for existing other liquid conversion
		} else if liquidUnit := unit == Cac || unit == Cas || unit == Ml; liquidUnit {
			LiqUnits := slices.Collect(maps.Keys(ConvertKg[item]))
			if len(LiqUnits) == 0 {
				err = fmt.Errorf("Unit %v for item %v does not exist in ConvertKg", unit, item)
				return
			}

			unit2 := LiqUnits[0]
			value2 := ConvertKg[item][unit2]

			kg = ConvertVol[unit][unit2] * value2
			return
		}
	} else {
		err = fmt.Errorf("Item %v does not exist in ConvertKg", item)
		return
	}

	return 0, fmt.Errorf("??? Unexpected error ????")
}

func CreateIngredient(name string, qty float64, unit Unit) IngredientBase {
	return IngredientBase{
		Name:     name,
		Quantity: qty,
		Unit:     unit,
		UnitName: UnitDict[unit],
	}
}

func (ing *IngredientBase) ConvertUnit(toUnit Unit) {
	switch {
	case ing.Unit == G && toUnit == Kg:
		ing.Quantity = ing.Quantity / 1000
		ing.Unit = toUnit
		ing.UnitName = UnitDict[toUnit]
	}
}

func (ing *Solid) ConvertUnit(toUnit Unit) {
	switch {
	case ing.Unit == G && toUnit == Kg:
		(*IngredientBase)(ing).ConvertUnit(toUnit)
	}
}
func (ing *Farine) ConvertUnit(toUnit Unit) {
	switch {
	case ing.Unit == G && toUnit == Kg:
		(*IngredientBase)(ing).ConvertUnit(toUnit)
	case ing.Unit == Ml && toUnit == G:
		ing.Quantity = ing.Quantity * 4 / 5
		ing.Unit = toUnit
		ing.UnitName = UnitDict[toUnit]
	}
}

// TODO
func (source *Ingredients) SetSameUnit(target Ingredients, ingrName string) {
}

func (ings Ingredients) GetIngredientIdByName(ingrName string) int {
	for i, ing := range ings {
		if ing.Name == ingrName {
			return i
		}
	}
	// panic("%v not in ingredients", ingrName)
	slog.Error(ingrName)
	slog.Error("not in ingredients")
	return -1
}

func (source *Ingredients) SetSameQuantity(target Ingredients, ingrName string) {
	source.SetSameUnit(target, ingrName)

	i := (*source).GetIngredientIdByName(ingrName)
	j := target.GetIngredientIdByName(ingrName)
	ratio := target[i].Quantity / (*source)[j].Quantity

	for i := range *source {
		(*source)[i].Quantity = (*source)[i].Quantity * ratio
	}
}
