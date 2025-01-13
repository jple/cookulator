package main

import (
	"fmt"
	"maps"
	"slices"
)

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

func f() {
	fmt.Println(ConvertKg)
	fmt.Println(ConvertKg["eau"][Cas])
	fmt.Println(ConvertKg["farine"][Cas])

	fmt.Println("===============")
	out, exist := ConvertKg["tot"][Cas]
	fmt.Println(out, exist)

	fmt.Println(GetInKg("farine", G))

	fmt.Println("===============")
	fmt.Println(slices.Collect(maps.Keys(ConvertKg["eau"])))
	fmt.Println(ConvertVol[Ml][Cas])
}

func main() {
	f()
	// v1 := Ingredients{
	// 	CreateIngredient("farine", 1000, G),
	// 	CreateIngredient("eau", 500, G),
	// }
	// v2 := Ingredients{
	// 	CreateIngredient("farine", 850, G),
	// 	CreateIngredient("eau", 400, G),
	// }
	// fmt.Printf("%+v\n", v1)
	// v1.SetSameQuantity(v2, "farine")
	// fmt.Printf("%+v\n", v1)

}
