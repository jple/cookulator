package main

import "fmt"

type Unit int

const (
	G Unit = iota
	Kg
	Cas
	Cac
	L
	Ml
)

var UnitDict = map[Unit]string{
	G:   "G",
	Kg:  "Kg",
	Cas: "Cas",
	Cac: "Cac",
	L:   "L",
	Ml:  "Ml",
}

func main() {
	v1 := Ingredients{
		CreateIngredient("farine", 1000, G),
		CreateIngredient("eau", 500, G),
	}
	v2 := Ingredients{
		CreateIngredient("farine", 850, G),
		CreateIngredient("eau", 400, G),
	}
	fmt.Printf("%+v\n", v1)
	v1.SetSameQuantity(v2, "farine")
	fmt.Printf("%+v\n", v1)

}
