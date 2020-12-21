package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type AllergenIngredient struct {
	allergen, ingredient string
}

func main() {
	var ingredientList [][]string
	allergenIngredients := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pieces := strings.Split(scanner.Text(), " (contains ")
		if len(pieces) > 1 {
			ingredients := strings.Split(pieces[0], " ")
			for _, allergen := range strings.Split(pieces[1][:len(pieces[1])-1], ", ") {
				if knownIngredients, ok := allergenIngredients[allergen]; ok {
					// no built-in intersection
					// no generic intersection
					// no conversion from []interface{} to {}string
					var intersectedIngredients []string
					for _, ingredient := range intersect.Simple(knownIngredients, ingredients).([]interface{}) {
						intersectedIngredients = append(intersectedIngredients, ingredient.(string))
					}
					allergenIngredients[allergen] = intersectedIngredients
				} else {
					allergenIngredients[allergen] = ingredients
				}
			}
			ingredientList = append(ingredientList, ingredients)
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	ingredientAllergen := make(map[string]string)
	for len(allergenIngredients) > 0 {
		for allergen, ingredients := range allergenIngredients {
			if len(ingredients) == 1 {
				ingredientAllergen[ingredients[0]] = allergen
				delete(allergenIngredients, allergen)
			} else {
				var filtered []string
				for _, ingredient := range ingredients {
					if _, ok := ingredientAllergen[ingredient]; !ok {
						filtered = append(filtered, ingredient)
					}
				}
				allergenIngredients[allergen] = filtered
			}
		}
	}

	nonAllergenic := 0
	for _, ingredients := range ingredientList {
		for _, ingredient := range ingredients {
			if _, ok := ingredientAllergen[ingredient]; !ok {
				nonAllergenic++
			}
		}
	}
	fmt.Println(nonAllergenic)

	var canonical []AllergenIngredient
	for ingredient, allergen := range ingredientAllergen {
		canonical = append(canonical, AllergenIngredient{allergen, ingredient})
	}
	sort.Slice(canonical, func(i, j int) bool { return canonical[i].allergen < canonical[j].allergen })

	for i, pair := range canonical {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(pair.ingredient)
	}
	fmt.Println()
}
