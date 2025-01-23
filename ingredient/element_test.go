package ingredient

import (
	"fmt"
	"testing"

	"ingredient-calculator/ingredient/unit"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		ingrName         string
		qty              float64
		fromUnit, toUnit unit.Unit
		want             float64
	}{
		{"farine", 100, unit.G, unit.Kg, 0.1},
		{"eau", 500, unit.G, unit.Ml, 500},
		{"eau", 1, unit.Cas, unit.Ml, 15},
		{"eau", 1, unit.Cas, unit.Cac, 3},
		{"farine", 1, unit.Cas, unit.Cac, 3},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("(%v: %v %v --> %v", test.ingrName, test.qty, unit.DictName[test.fromUnit], unit.DictName[test.toUnit])
		t.Run(testname, func(t *testing.T) {
			x := CreateElement(test.ingrName, test.qty, test.fromUnit)
			err := x.Convert(test.toUnit)
			if err != nil {
				t.Errorf(err.Error())
			}

			have := x.Quantity
			if have != test.want {
				t.Errorf("Have %v, test.want %v", have, test.want)
			}
		})
	}
}
