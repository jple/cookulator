package main

import (
	"fmt"
	"log/slog"
)

// Ingredient
type Ingredient interface {
	ConvertUnit(Unit)
}

type UnitToKg map[Unit]float64

var ConvertKg = map[string]UnitToKg{
	"eau":    UnitToKg{Ml: 0.001, Cas: 0.0015},
	"farine": UnitToKg{Cas: 4.0 / 5.0 / 1000},
}

// var GenericConvertKg = map[Unit]float64{
// 	G:        0.001,
// 	"farine": UnitToKg{Cas: 4.0 / 5.0 / 1000},
// }
//TODO:
// G = 0.001 Kg
// Cas = 3*Cac (volume)
// Cac = 5 ml

func f() {
	fmt.Println(ConvertKg)
	fmt.Println(ConvertKg["eau"][Cas])
	fmt.Println(ConvertKg["farine"][Cas])

	fmt.Println("===============")
	out, exist := ConvertKg["tot"][Cas]
	fmt.Println(out, exist)

	fmt.Println(GetInKg("farine", G))
}

func GetInKg(item string, unit Unit) (kg float64, err error) {
	ConvertItem, itemExist := ConvertKg[item]
	if itemExist {
		value, unitExist := ConvertItem[unit]
		if unitExist {
			kg = value
		} else {
			//TODO: add generic conversion, as G to Kg
			if unit == G {
				kg = 0.001
			} else {
				err = fmt.Errorf("Unit %v does not exist for item %v in ConvertKg", UnitDict[unit], item)
			}
		}
	} else {
		err = fmt.Errorf("Item %v does not exist in ConvertKg", item)
	}

	// kg = ConvertKg[item][unit]
	return
}

// IngredientBase is the basic type struct for all ingredients (solid, liquid, ...)
type IngredientBase struct {
	Name     string
	Quantity float64
	Unit     Unit
	UnitName string
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

type Solid IngredientBase
type Farine Solid

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

// {
// 	name="Farine",
// 	Qty = 100
// 	{from=G, to=Kg, result=

type Ingredients []IngredientBase
type IngredientsVariant []Ingredients

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
