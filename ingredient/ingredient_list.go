package ingredient

import "log/slog"

type (
	List []Content // TODO: change to map[string]Content
)

// TODO
// MAY NOT BE WORKING
func ListAdd(l List, c Content) List {
	return List{
		c.Name: c,
		l,
	}
}
func ListNew() List {
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

// TODO
func (source *List) SetSameUnit(target List, ingrName string) {
}

func ConvertList(compareList []List, refId int, refIngr string) []List {
	return nil
}
