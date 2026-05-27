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

func feuilleteInverse() {
	inverse := ingredient.CompareList{
		// source : https://www.mercotte.fr/2025/01/09/le-meilleur-patissier-saison-13-emission-14-lmp-se-met-sur-son-31-la-finale-bonne-annee-et-la-galette-des-rois-de-nina-metayer/

		// cercle 22 cm + 20 cm
		ingredient.List{
			// qté pour beurre manié à gauche du "+"
			// qté pour détrempe à droite du "+"
			ingredient.CreateElement("beurre", 225+60, unit.G),
			ingredient.CreateElement("farine", 100+185, unit.G),
			ingredient.CreateElement("eau", 85, unit.G),
			ingredient.CreateElement("vinaigre", 2.5, unit.G),
			ingredient.CreateElement("sel", 4, unit.G),
		},

		// source : https://www.mercotte.fr/2008/04/08/la-pate-feuilletee-inversee/

		// Version Christophe Felder
		// 2 foix 24 cm
		// 350 g de pâte
		ingredient.List{
			// qté pour beurre manié à gauche du "+"
			// qté pour détrempe à droite du "+"
			ingredient.CreateElement("eau", 150, unit.G),
			ingredient.CreateElement("vinaigre", 1, unit.Cas),
			ingredient.CreateElement("sel", 15, unit.G),
			ingredient.CreateElement("farine", 150+350, unit.G),
			ingredient.CreateElement("beurre", 375+115, unit.G),
		},

		// Version Pierre Hermé
		//
		//
		ingredient.List{
			// qté pour beurre manié à gauche du "+"
			// qté pour détrempe à droite du "+"
			ingredient.CreateElement("eau", 150, unit.G),
			ingredient.CreateElement("vinaigre", 1, unit.Cas),
			ingredient.CreateElement("sel", 15, unit.G),
			ingredient.CreateElement("farine", 75*2+175*2, unit.G),
			ingredient.CreateElement("beurre", 375+110, unit.G),
		},

		// https://www.chefnini.com/pate-feuilletee-inversee/
		// Version chefnini (inversée)
		//
		// 750g de pate
		ingredient.List{
			ingredient.CreateElement("beurre", 185+55, unit.G),
			ingredient.CreateElement("farine", 75+175, unit.G),
			ingredient.CreateElement("eau", 75, unit.G),
			ingredient.CreateElement("vinaigre", 0, unit.G),
			ingredient.CreateElement("sel", 5, unit.G),
		},
	}

	classique := ingredient.CompareList{
		// repri
		inverse[0],
		// Version livre (classique)
		//
		// 1kg de pate
		ingredient.List{
			ingredient.CreateElement("farine", 500, unit.G),
			ingredient.CreateElement("eau", 230, unit.G),
			ingredient.CreateElement("sel", 10, unit.G),
			ingredient.CreateElement("beurre", 60+300, unit.G),
		},

		// https://www.mercotte.fr/2005/06/02/la-pate-feuilletee-pourquoi-tant-de-haine/
		// Version mercote (classique)
		//
		// 750g de pate
		ingredient.List{
			ingredient.CreateElement("farine", 250, unit.G),
			ingredient.CreateElement("eau", 125, unit.G),
			ingredient.CreateElement("vinaigre", 5, unit.G),
			ingredient.CreateElement("sel", 7, unit.G),
			ingredient.CreateElement("beurre", 50+250, unit.G),
		},

		// https://www.mamiecaillou.com/2020/01/pate-feuilletee-de-philippe-conticini.il-faut-bien-se-lancer-un-jour-alors-autant-que-ce-soit-avec-les-conseils-d-un-maitre.html
		// Version Conchini(classique)
		//
		// 750g de pate
		ingredient.List{
			ingredient.CreateElement("farine", 500, unit.G),
			ingredient.CreateElement("eau", 250, unit.G),
			ingredient.CreateElement("sel", 12, unit.G),
			ingredient.CreateElement("beurre", 400, unit.G),
		},

		// https://www.chefnini.com/pate-feuilletee/
		// Version chefnini (classique)
		//
		// 750g de pate
		ingredient.List{
			ingredient.CreateElement("beurre", 250+50, unit.G),
			ingredient.CreateElement("farine", 250, unit.G),
			ingredient.CreateElement("eau", 125, unit.G),
			ingredient.CreateElement("sel", 7, unit.G),
		},

		// https://encoreungateau.com/pate-feuilletee-maison-pur-beurre/
		// Version CAP
		//
		// 750g de pate
		ingredient.List{
			ingredient.CreateElement("farine", 125*2, unit.G),
			ingredient.CreateElement("eau", 125, unit.G),
			ingredient.CreateElement("sel", 5, unit.G),
			ingredient.CreateElement("beurre", 175, unit.G),
		},

		// https://lemondeculinairedesamia.com/comment-faire-une-pate-feuilletee-classique-ou-ordinaire-cap-patisserie/
		// Version CAP 2
		ingredient.List{
			ingredient.CreateElement("farine", 500, unit.G),
			ingredient.CreateElement("eau", 230, unit.G),
			ingredient.CreateElement("vinaigre", 12, unit.G),
			ingredient.CreateElement("beurre", 60+250, unit.G),
			ingredient.CreateElement("sel", 12, unit.G),
		},

		// https://www.meilleurduchef.com/fr/recette/pate-feuilletee-cap-patissier-video.html
		// Version CAP 3
		ingredient.List{
			ingredient.CreateElement("farine", 300, unit.G),
			ingredient.CreateElement("eau", 160, unit.G),
			ingredient.CreateElement("beurre", 230, unit.G),
			ingredient.CreateElement("sel", 7, unit.G),
		},

		// https://lesplaisirssucresdantoine.over-blog.com/2020/01/pate-feuilletee-facon-cap-patissier.html
		// Version CAP 4
		ingredient.List{
			ingredient.CreateElement("farine", 500, unit.G),
			ingredient.CreateElement("eau", 250, unit.G),
			ingredient.CreateElement("beurre", 375, unit.G),
			ingredient.CreateElement("sel", 10, unit.G),
		},
	}

	fmt.Println("Inverse")
	inverse.Show()
	fmt.Println("===============")
	inverse = ingredient.ConvertList(inverse, 0, "beurre")
	inverse.Show()
	fmt.Println("===============")
	fmt.Println("Classique")
	classique.Show()
	fmt.Println("===============")
	classique = ingredient.ConvertList(classique, 0, "beurre")
	classique.Show()

}

func main() {
	feuilleteInverse()
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
