package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["alice"] += 2
	ages["charlie"]++
	fmt.Println(ages)
	fmt.Println(ages["charlie"])
	delete(ages, "alice")
	fmt.Println(ages)
}
