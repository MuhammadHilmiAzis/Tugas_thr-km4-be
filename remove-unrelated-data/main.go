package main

import "fmt"

func removeUnrelatedData(mapData map[string]interface{}, key string) map[string]interface{} {
	delete(mapData, key)
	return mapData
}

func main() {
	map1 := map[string]interface{}{"name": "Edo", "age": 20, "address": "Jakarta"}
	fmt.Println(removeUnrelatedData(map1, "address")) // output: map[name:Edo age:20]

	map2 := map[string]interface{}{"run": "lari", "jump": "loncat", "swim": "berenang"}
	fmt.Println(removeUnrelatedData(map2, "jump")) // output: map[run:lari swim:berenang]

	map3 := map[string]interface{}{"satu": "ichi", "dua": "ni", "tiga": "san", "empat": "yon", "lima": "go"}
	fmt.Println(removeUnrelatedData(map3, "tiga")) // output: map[satu:ichi dua:ni empat:yon lima:go]
}
