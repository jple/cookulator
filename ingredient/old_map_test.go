package ingredient

import (
	"ingredient-calculator/ingredient/unit"
	"testing"
)

type Map map[string]Element

func (m Map) GetElement(ingrName string) Element {
	return m[ingrName]
}

func BenchmarkGetElementFromMap(b *testing.B) {
	m := Map{
		"farine":    CreateElement("farine", 1000, unit.G),
		"eau":       CreateElement("eau", 1000, unit.G),
		"aze":       CreateElement("aze", 1000, unit.G),
		"toto":      CreateElement("toto", 1000, unit.G),
		"ajf":       CreateElement("ajf", 1000, unit.G),
		"ojzepv":    CreateElement("ojzepv", 1000, unit.G),
		"zporvnio":  CreateElement("zporvnio", 1000, unit.G),
		"zoaeijvn":  CreateElement("zoaeijvn", 1000, unit.G),
		"xcvmlkndf": CreateElement("xcvmlkndf", 1000, unit.G),
	}

	for i := 0; i < b.N; i++ {
		m.GetElement("farine")
	}
}

func BenchmarkGetElementFromList(b *testing.B) {
	l := List{
		CreateElement("farine", 1000, unit.G),
		CreateElement("eau", 1000, unit.G),
		CreateElement("aze", 1000, unit.G),
		CreateElement("toto", 1000, unit.G),
		CreateElement("ajf", 1000, unit.G),
		CreateElement("ojzepv", 1000, unit.G),
		CreateElement("zporvnio", 1000, unit.G),
		CreateElement("zoaeijvn", 1000, unit.G),
		CreateElement("xcvmlkndf", 1000, unit.G),
	}

	for i := 0; i < b.N; i++ {
		l.GetElement("farine")
	}
}
