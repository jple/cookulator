package ingredient

import (
	"fmt"
	"log/slog"

	"ingredient-calculator/ingredient/unit"
)

type (
	// Content is the basic type struct describing ingredient
	Content struct {
		Name     string
		Quantity float64
		Unit     unit.Unit
		UnitName string
	}

	List []Content
)

func (ing *Content) ConvertUnit(toUnit unit.Unit) error {

	if toUnit == ing.Unit {
		return nil
	}

	toKg1, err1 := unit.ToKg(ing.Name, ing.Unit)
	toKg2, err2 := unit.ToKg(ing.Name, toUnit)
	if err1 != nil {
		return err1
		return fmt.Errorf("1) %v", err1)
	}
	if err2 != nil {
		// return err2
		return fmt.Errorf("2) %v", err2)
	}

	if toKg2 == 0 {
		return fmt.Errorf("Dividing to zero")
	}

	ing.Quantity = ing.Quantity * toKg1 / toKg2
	ing.Unit = toUnit
	ing.UnitName = unit.DictName[toUnit]

	return nil
}

// TODO
func (source *List) SetSameUnit(target List, ingrName string) {
}

func (ings List) GetContentIdByName(ingrName string) int {
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

func (source *List) SetSameQuantity(target List, ingrName string) {
	source.SetSameUnit(target, ingrName)

	i := (*source).GetContentIdByName(ingrName)
	j := target.GetContentIdByName(ingrName)
	ratio := target[i].Quantity / (*source)[j].Quantity

	for i := range *source {
		(*source)[i].Quantity = (*source)[i].Quantity * ratio
	}
}
