package main

import "fmt"

func mapDemo() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// IMP: if key does not exists in the map then it returns zero value

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}

// m := make(map[string]string)

// setting an element
// m["name"] = "golang"
// m["area"] = "backend"

// get an element
// fmt.Println(m["name"], m["area"])
// IMP: if key does not exists in the map then it returns zero value

// m := make(map[string]int)
// m["age"] = 30
// m["price"] = 50
// fmt.Println(m["phone"])
// fmt.Println(len(m))

// delete(m, "price")
// clear(m)

// fmt.Println(m)
// fmt.Println(m)

// m := map[string]int{"price": 40, "phones": 3}

// v, ok := m["phones"]
// fmt.Println(v)
// if ok {
// 	fmt.Println("all ok")
// } else {
// 	fmt.Println("not ok")
// }

// m1 := map[string]int{"price": 40, "phones": 3}
// m2 := map[string]int{"price": 40, "phones": 8}
// fmt.Println(maps.Equal(m1, m2)) // false if phonse value is same then true

// separate object or memory of m1 and m2. maps.Equal() just compares the values of the maps
