package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func contains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func findAllergens(allergens map[string]bool, ingredients map[string]bool, foods []food) ([]string, map[string][]string) {
	defAll := []string{}
	allMap := make(map[string][]string)
	for all := range allergens {
		couldBe := ingredients

		for _, f := range foods {
			if !contains(f.allergens, all) {
				continue
			}
			ing := f.ingredients
			newMap := make(map[string]bool)
			for i := range couldBe {
				if contains(ing, i) {
					newMap[i] = true
				}
			}
			couldBe = newMap
		}

		possibles := []string{}
		for i := range couldBe {
			if !contains(defAll, i) {
				defAll = append(defAll, i)
			}
			possibles = append(possibles, i)
		}
		allMap[all] = possibles
	}
	return defAll, allMap
}

func reduceRecipe(foods []food, badIng []string) []food {
	newFoods := []food{}
	for _, f := range foods {
		ing := []string{}
		for _, i := range f.ingredients {
			if contains(badIng, i) {
				ing = append(ing, i)
			}
		}
		newFoods = append(newFoods, food{ingredients: ing, allergens: f.allergens})
	}
	return newFoods
}

func main() {
	file := utils.ReadFile(os.Args[1])

	re := regexp.MustCompile(`^((?:(?:[a-z]+)\s)+)\(contains ((?:[a-z]+(?:, )?)+)\)$`)
	ingredients := make(map[string]bool)
	allergens := make(map[string]bool)
	foods := []food{}
	for _, line := range file {
		result := re.FindSubmatch([]byte(line))
		ing := strings.Split(string(result[1]), " ")
		all := strings.Split(string(result[2]), ", ")

		listIng := []string{}

		for _, x := range ing {
			if x == "" {
				continue
			}
			listIng = append(listIng, x)
			if _, ok := ingredients[x]; !ok {
				ingredients[x] = true
			}
		}
		for _, x := range all {
			if x == "" {
				continue
			}
			if _, ok := allergens[x]; !ok {
				allergens[x] = true
			}
		}
		newFood := food{ingredients: listIng, allergens: all}
		foods = append(foods, newFood)
	}

	defAll, _ := findAllergens(allergens, ingredients, foods)

	var safe []string
	for in := range ingredients {
		// if in not contained in defAll then safe
		if !contains(defAll, in) {
			safe = append(safe, in)
		}
	}

	count := 0
	for _, s := range safe {
		for _, f := range foods {
			if contains(f.ingredients, s) {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)

	recipes := reduceRecipe(foods, defAll)
	badIng := make(map[string]bool)
	for _, i := range defAll {
		badIng[i] = true
	}
	_, allMap := findAllergens(allergens, badIng, recipes)

	goodMap := make(map[string]string)
	for len(goodMap) != len(allMap) {
		// find item in allMap that has len 1
		var all string
		var ing string
		for a, i := range allMap {
			if len(i) == 1 {
				all = a
				ing = i[0]
				break
			}
		}
		//remove ing in all items of allMap
		for key, list := range allMap {
			newList := []string{}
			for _, item := range list {
				if item != ing {
					newList = append(newList, item)
				}
			}
			allMap[key] = newList
		}
		goodMap[all] = ing
	}

	allList := []string{}
	for key := range allMap {
		allList = append(allList, key)
	}
	sort.Strings(allList)

	part2 := ""
	for i, a := range allList {
		if i == 0 {
			part2 = goodMap[a]
		} else {
			part2 = part2 + "," + goodMap[a]
		}
	}

	fmt.Println("Part 2:", part2)
}
