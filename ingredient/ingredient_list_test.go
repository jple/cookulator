package ingredient

import (
	"ingredient-calculator/ingredient/unit"
	"testing"
)

func TestSetSameQty(t *testing.T) {
	v1 := List{
		CreateContent("farine", 1000, unit.G),
		CreateContent("eau", 500, unit.G),
	}
	v2 := List{
		CreateContent("farine", 850, unit.G),
		CreateContent("eau", 400, unit.G),
	}
	v1.SetSameQuantity(v2, "farine")

	if v1[1].Quantity != 425 {
		t.Fatalf("Quantity should return 425")
	}
}

func TestConvertList(t *testing.T) {
	v1 := List{
		CreateContent("farine", 1000, unit.G),
		CreateContent("eau", 500, unit.G),
	}
	v2 := List{
		CreateContent("farine", 850, unit.G),
		CreateContent("eau", 400, unit.G),
	}

	var compareList = []List{
		v1, v2,
	}

	equalizedList := ConvertList(compareList, 0, "farine")
	vv1 := equalizedList[0]
	vv2 := equalizedList[1]

	//TODO: checks :
	// vv2.farine.qty == 1000
	// vv2.eau.qty == 400 * 1000/850

}
