package main

import (
	"fmt"
	"ingredient-calculator/ingredient"
	"ingredient-calculator/ingredient/unit"
	"maps"
	"slices"
)

func f() {
	l1 := ingredient.List{
		ingredient.CreateElement("farine", 1000, unit.G),
		ingredient.CreateElement("eau", 500, unit.G),
		ingredient.CreateElement("lait", 50, unit.G),
	}

	// l1.Show()

	// fmt.Println("=============")

	l2 := ingredient.List{
		ingredient.CreateElement("farine", 600, unit.G),
		ingredient.CreateElement("sucre", 10, unit.G),
		ingredient.CreateElement("lait", 3, unit.Cas),
		ingredient.CreateElement("eau", 250, unit.G),
	}

	c := ingredient.CompareList{l1, l2}
	c.Show()
	fmt.Println("==============")
	c = ingredient.ConvertList(c, 0, "farine")
	c.Show()

	l1 = ingredient.List{
		ingredient.CreateElement("farine", 1000, unit.G),
		ingredient.CreateElement("eau", 500, unit.G),
	}
	l2 = ingredient.List{
		ingredient.CreateElement("farine", 850, unit.G),
		ingredient.CreateElement("eau", 400, unit.G),
	}
	l3 := ingredient.List{
		ingredient.CreateElement("farine", 850, unit.G),
		ingredient.CreateElement("eau", 80, unit.Cac),
	}
	c = ingredient.CompareList{l1, l2, l3}
	c.Show()
	fmt.Println("=================")
	c = ingredient.ConvertList(c, 0, "farine")
	c.Show()

	fmt.Println("=================")
	fmt.Println(unit.ToKg("eau", unit.Cas))
	fmt.Println(unit.ToKg("lait", unit.Cas))

	fmt.Println("=================")

	fmt.Printf("%+v\n", slices.Collect(maps.Keys(unit.DictToKg["eau"])))
	fmt.Printf("%+v\n", unit.DictToKg["eau"])

}

func main() {
	f()
	// v1 := List{
	// 	CreateContent("farine", 1000, G),
	// 	CreateContent("eau", 500, G),
	// }
	// v2 := List{
	// 	CreateContent("farine", 850, G),
	// 	CreateContent("eau", 400, G),
	// }
	// fmt.Printf("%+v\n", v1)
	// v1.SetSameQuantity(v2, "farine")
	// fmt.Printf("%+v\n", v1)

}
