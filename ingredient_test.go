package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertUnit(t *testing.T) {
	x := CreateIngredient("farine", 100, G)
	x.ConvertUnit(Kg)
	if x.Quantity != 0.1 {
		t.Fatalf("ConvertUnit should return 0.1")
	}
}

func TestSolidConvertUnit(t *testing.T) {
	x := Solid{
		Name:     "farine",
		Quantity: 100,
		Unit:     G,
		UnitName: UnitDict[G],
	}
	x.ConvertUnit(Kg)
	if x.Quantity != 0.1 {
		t.Fatalf("ConvertUnit should return 0.1")
	}
}
func TestFarineConvertUnit(t *testing.T) {
	x := Farine{
		Name:     "farine",
		Quantity: 5,
		Unit:     Ml,
		UnitName: UnitDict[Ml],
	}
	x.ConvertUnit(G)
	if x.Quantity != 4 {
		t.Fatalf("ConvertUnit should return 4")
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
