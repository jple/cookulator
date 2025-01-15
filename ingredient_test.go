package main

import (
	"fmt"
	"strings"
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

func TestConvertUnit(t *testing.T) {
	x := CreateIngredient("farine", 100, G)
	err := x.ConvertUnit(Kg)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if x.Quantity != 0.1 {
		t.Fatalf("ConvertUnit should return 0.1, but return %v", x.Quantity)
	}

	ing2 := CreateIngredient("eau", 500, G)
	ing2.ConvertUnit(Ml)

	if !(ing2.Quantity == 500 && ing2.Unit == Ml && ing2.UnitName == UnitDict[ing2.Unit]) {
		t.Fatalf("Did not work")
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

type tester struct {
	call string
	have float64
	want float64
}

func createTesterInput(ingrName string, unit Unit, want float64) tester {
	have, err := GetInKg(ingrName, unit)
	if err != nil {
		panic(err)
	}

	return tester{
		call: fmt.Sprintf(`GetInKg("%v", %v)`, ingrName, UnitDict[unit]),
		have: have,
		want: want,
	}
}

func TestGetInKg(t *testing.T) {
	var errMsg []string
	var tests []tester

	tests = append(tests,
		createTesterInput("farine", G, 0.001),
		createTesterInput("any", G, 0.001),
		createTesterInput("eau", Cas, 0.015),
		createTesterInput("eau", Cac, 0.005),
		createTesterInput("farine", Cac, ConvertKg["farine"][Cas]/3),
	)

	for _, test := range tests {
		if test.have != test.want {
			errMsg = append(errMsg,
				fmt.Sprintf("%v returns %v, while expected %v\n",
					test.call, test.have, test.want))
		}
	}

	if errMsg != nil {
		t.Fatalf(strings.Join(errMsg, ""))
	}

}
