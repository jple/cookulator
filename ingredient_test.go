package main

import (
	"testing"
)

func TestConvertUnit(t *testing.T) {
	x := CreateIngredient("farine", 100, G)
	x.ConvertUnit(Kg)
	if x.Quantity != 0.1 {
		t.Fatalf("ConvertUnit should return 0.1")
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
