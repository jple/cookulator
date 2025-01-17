package main

import (
	"fmt"
	"testing"
)

func CreateIngredient(name string, qty float64, unit Unit) IngredientBase {
	return IngredientBase{
		Name:     name,
		Quantity: qty,
		Unit:     unit,
		UnitName: UnitDict[unit],
	}
}

func TestGetInKg(t *testing.T) {
	var tests = []struct {
		ingrName string
		unit     Unit
		want     float64
	}{
		{"farine", G, 0.001},
		{"any", G, 0.001},
		{"eau", Cas, 0.015},
		{"farine", Cac, ConvertKg["farine"][Cas] / 3},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("(%v: %v --> Kg", test.ingrName, UnitDict[test.unit])
		t.Run(testname, func(t *testing.T) {
			have, err := GetInKg(test.ingrName, test.unit)
			if err != nil {
				t.Errorf(err.Error())
			}
			if have != test.want {
				t.Errorf("Have %v, want %v", have, test.want)
			}
		})
	}
}

func TestConvertUnit(t *testing.T) {
	var tests = []struct {
		ingrName         string
		qty              float64
		fromUnit, toUnit Unit
		want             float64
	}{
		{"farine", 100, G, Kg, 0.1},
		{"eau", 500, G, Ml, 500},
		{"eau", 1, Cas, Ml, 15},
		{"eau", 1, Cas, Cac, 3},
		{"farine", 1, Cas, Cac, 3},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("(%v: %v %v --> %v", test.ingrName, test.qty, UnitDict[test.fromUnit], UnitDict[test.toUnit])
		t.Run(testname, func(t *testing.T) {
			x := CreateIngredient(test.ingrName, test.qty, test.fromUnit)
			err := x.ConvertUnit(test.toUnit)
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

func TestSetSameQty(t *testing.T) {
	v1 := Ingredients{
		CreateIngredient("farine", 1000, G),
		CreateIngredient("eau", 500, G),
	}
	v2 := Ingredients{
		CreateIngredient("farine", 850, G),
		CreateIngredient("eau", 400, G),
	}
	v1.SetSameQuantity(v2, "farine")

	if v1[1].Quantity != 425 {
		t.Fatalf("Quantity should return 425")
	}
}
