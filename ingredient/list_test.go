package ingredient

import (
	"fmt"
	"ingredient-calculator/ingredient/unit"
	"testing"
)

func TestSetSameQty(t *testing.T) {
	v1 := List{
		CreateElement("farine", 1000, unit.G),
		CreateElement("eau", 500, unit.G),
	}
	v2 := List{
		CreateElement("farine", 850, unit.G),
		CreateElement("eau", 400, unit.G),
	}
	v1.SetSameQuantity(v2, "farine")

	if v1[1].Quantity != 425 {
		t.Fatalf("Quantity should return 425")
	}
}

func TestGetElement(t *testing.T) {
	farine := CreateElement("farine", 1000, unit.G)
	eau := CreateElement("eau", 500, unit.G)

	l := List{farine, eau}

	var tests = []struct {
		ingrName string
		want     Element
		// wantError   error
	}{
		{"farine", farine},
		{"eau", eau},
		{"notPresent", Element{}},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("Getting element: %v", test.ingrName)
		t.Run(testname, func(t *testing.T) {
			el, _ := l.GetElement(test.ingrName)
			if el != test.want {
				t.Errorf("Have %v, want %v", el, test.want)
			}
		})

	}
}

// TODO
//func TestConvertList(t *testing.T) {
//	v1 := List{
//		CreateElement("farine", 1000, unit.G),
//		CreateElement("eau", 500, unit.G),
//	}
//	v2 := List{
//		CreateElement("farine", 850, unit.G),
//		CreateElement("eau", 400, unit.G),
//	}

//	var compareList = []List{
//		v1, v2,
//	}

//	equalizedList := ConvertList(compareList, 0, "farine")
//	vv1 := equalizedList[0]
//	vv2 := equalizedList[1]

//	//TODO: checks :
//	// vv2.farine.qty == 1000
//	// vv2.eau.qty == 400 * 1000/850

//}
