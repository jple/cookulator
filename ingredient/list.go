package ingredient

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"text/tabwriter"
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

	// TODO: regacto: use applyCrossMult
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

func (l List) Show() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '.', tabwriter.AlignRight|tabwriter.Debug)
	for _, el := range l {
		fmt.Fprintln(w, el.Name, "\t", el.Quantity, el.UnitName, "\t")
	}
	w.Flush()
}

type CompareList []List

func (ll CompareList) GetAllNames() []string {
	var names []string
	for _, l := range ll {
		for _, el := range l {
			if !slices.Contains(names, el.Name) {
				names = append(names, el.Name)
			}
		}
	}
	return names
}

func (l *List) Swap(i int, j int) error {
	ll := *l

	if i >= len(ll) || j >= len(ll) {
		return errors.New("i or j is greater than len(l)")
	}
	old := ll[i]
	ll[i] = ll[j]
	ll[j] = old
	return nil
}

func (ll CompareList) SortByName() CompareList {
	for j, ingr := range ll.GetAllNames() {
		for _, l := range ll {
			k, err := l.GetElementIdByName(ingr)
			if err == nil {
				l.Swap(j, k)
			} else {
				l = append(l, Element{})
				n := len(l)
				l.Swap(j, n-1)
			}
		}
	}
	return ll
}

func (ll CompareList) Show() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '.', tabwriter.AlignRight|tabwriter.Debug)
	ll = ll.SortByName()
	for _, name := range ll.GetAllNames() {
		fmt.Fprintf(w, "%v\t", name)
		for _, l := range ll {
			el, err := l.GetElement(name)
			if err == nil {
				fmt.Fprintf(w, "%.0f %v\t", el.Quantity, el.UnitName)
			} else {
				fmt.Fprintf(w, "0\t")
			}
		}
		fmt.Fprintf(w, "\n")
	}

	w.Flush()
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
			compareList[i][j].Convert(refEl.Unit)
			compareList[i][j].Unit = refEl.Unit
			compareList[i][j].UnitName = refEl.UnitName

			compareList[i][j].Quantity = applyCrossMult(
				compareList[i][j].Quantity,
				el.Quantity,
				refEl.Quantity,
			)
		}
	}
	return compareList
}
