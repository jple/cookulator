package ingredient

import (
	"fmt"

	"ingredient-calculator/ingredient/unit"
)

type (
	// Element is the basic type struct describing ingredient
	Element struct {
		Name     string
		Quantity float64
		Unit     unit.Unit
		UnitName string
	}
)

func (ing *Element) Convert(toUnit unit.Unit) error {

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
