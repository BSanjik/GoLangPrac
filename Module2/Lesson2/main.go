package main

import "fmt"

func main() {
	// country := map[string]string{"Kazakhstan": "Astana", "Turkey": "Istambul", "Russia": "Moscow"}
	// fmt.Println(country)
	// 	beerMenu := make(map[string]string)
	// 	beerMenu["Shymkent's Beer"] = "fish" // add to map
	// 	beerMenu["Blanc 1669"] = "cheese"
	// 	fmt.Println(beerMenu)
	// 	delete(beerMenu, "Shymkent's Beer") //delete to map
	// 	fmt.Println(beerMenu)
	// 	fmt.Println(beerMenu["Shymkent's Beer"])
	// map1 := map[string]int{}
	// map2 := map[string]string{}
	// map3 := map[string]bool{}

	// map2["banana"] = "2"

	// fmt.Println(map1["apple"], map2["banana"], map3["orange"])
	// _, ok := map2["banana"]
	// if ok {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }
	// fmt.Println(ok)

	country := map[string]string{"Kazakhstan": "Astana", "Turkey": "Istambul", "Russia": "Moscow"}
	for i, j := range country {
		fmt.Println("The capital of ", i, "is:", j)
	}
}
