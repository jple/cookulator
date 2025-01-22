package ingredient

import (
	"fmt"
)

type (
	List []Element
)

// TODO
// MAY NOT BE WORKING
// func ListAdd(l List, c Element) List {
// 	return List{
// 		c.Name: c,
// 		l,
// 	}
// }
// func ListNew() List {
// }

func (ings List) GetElementIdByName(ingrName string) (int, error) {
	for i, ing := range ings {
		if ing.Name == ingrName {
			return i, nil
		}
	}

	return -1, fmt.Errorf("%v is in the List", ingrName)
}

func (ings List) GetElement(ingrName string) (Element, error) {
	i, err := ings.GetElementIdByName(ingrName)
	if err != nil {
		return Element{}, err
	}
	return ings[i], nil
}

func (source *List) SetSameQuantity(target List, ingrName string) {
	source.SetSameUnit(target, ingrName)

	i, err1 := (*source).GetElementIdByName(ingrName)
	if err1 != nil {
		panic(err1)
	}
	j, err2 := target.GetElementIdByName(ingrName)
	if err2 != nil {
		panic(err2)
	}

	ratio := target[i].Quantity / (*source)[j].Quantity

	for i := range *source {
		(*source)[i].Quantity = (*source)[i].Quantity * ratio
	}
}

// TODO
func (source *List) SetSameUnit(target List, ingrName string) {
}

func applyCrossMult(x, y, y2 float64) float64 {
	return x * y2 / y
}

// TODO: rename to Equalized
func ConvertList(compareList []List, refListId int, refIngr string) []List {
	refEl, err := compareList[refListId].GetElement(refIngr)
	if err != nil {
		panic(err.Error())
	}
	for i, _ := range compareList {
		el, err := compareList[i].GetElement(refIngr)
		if err != nil {
			panic(err.Error())
		}
		err = el.Convert(refEl.Unit)
		if err != nil {
			panic(err.Error())
		}

		for j, _ := range compareList[i] {
			compareList[i][j].Quantity = applyCrossMult(
				compareList[i][j].Quantity,
				el.Quantity,
				refEl.Quantity,
			)
			compareList[i][j].Unit = refEl.Unit
			compareList[i][j].UnitName = refEl.UnitName
		}
	}
	return compareList
}
