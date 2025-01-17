package main

import (
	"fmt"
	. "ingredient-calculator/ingredient/unit"
)

func f() {
	fmt.Println(ConvertKg)
	fmt.Println(ConvertKg["eau"][Cas])
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
