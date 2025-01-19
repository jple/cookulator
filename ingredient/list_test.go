package ingredient

import (
	"fmt"
	"ingredient-calculator/ingredient/unit"
	"testing"
)

func TestSetSameQuantity(t *testing.T) {
	l1 := List{
		CreateElement("farine", 1000, unit.G),
		CreateElement("eau", 500, unit.G),
	}
	l2 := List{
		CreateElement("farine", 850, unit.G),
		CreateElement("eau", 400, unit.G),
	}
	l1.SetSameQuantity(l2, "farine")

	if l1[1].Quantity != 425 {
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
func TestConvertList(t *testing.T) {
	l1 := List{
		CreateElement("farine", 1000, unit.G),
		CreateElement("eau", 500, unit.G),
	}
	l2 := List{
		CreateElement("farine", 850, unit.G),
		CreateElement("eau", 400, unit.G),
	}

	// var compareList = []List{l1, l2}

	refIngr := "farine"
	stdList := ConvertList([]List{l1, l2}, 0, refIngr)
	ll2 := stdList[1]

	conv := func(x, y, refY float64) float64 {
		return x * refY / y
	}

	var have, want float64
	for _, ingr := range []string{"farine", "eau"} {
		x, _ := l2.GetElement(ingr)
		y, _ := l2.GetElement(refIngr)
		refY, _ := l1.GetElement(refIngr)
		want = conv(x.Quantity, y.Quantity, refY.Quantity)

		el, _ := ll2.GetElement(ingr)
		have = el.Quantity

		if have != want {
			t.Errorf("%v: Have %v, want %v", ingr, have, want)
		}

	}
}
