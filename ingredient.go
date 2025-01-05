package main

import "log/slog"

type Ingredient struct {
	Name     string
	Quantity float64
	Unit     Unit
	UnitName string
}

func CreateIngredient(name string, qty float64, unit Unit) Ingredient {
	return Ingredient{
		Name:     name,
		Quantity: qty,
		Unit:     unit,
		UnitName: UnitDict[unit],
	}
}

func (ing *Ingredient) ConvertUnit(toUnit Unit) {
	switch {
	case ing.Unit == G && toUnit == Kg:
		ing.Quantity = ing.Quantity / 1000
		ing.Unit = toUnit
		ing.UnitName = UnitDict[toUnit]
	}
}

type Ingredients []Ingredient
type IngredientsVariant []Ingredients

func (source *Ingredients) SetSameUnit(target Ingredients, ingrName string) {
	// TODO
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
