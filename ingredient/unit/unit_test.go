package unit

import (
	"fmt"
	"testing"
)

func TestToKg(t *testing.T) {
	var tests = []struct {
		ingrName string
		unit     Unit
		want     float64
	}{
		{"farine", G, 0.001},
		{"any", G, 0.001},
		{"eau", Cas, 0.015},
		{"farine", Cac, DictToKg["farine"][Cas] / 3},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("(%v: %v --> Kg", test.ingrName, DictName[test.unit])
		t.Run(testname, func(t *testing.T) {
			have, err := ToKg(test.ingrName, test.unit)
			if err != nil {
				t.Errorf(err.Error())
			}
			if have != test.want {
				t.Errorf("Have %v, want %v", have, test.want)
			}
		})
	}
}
