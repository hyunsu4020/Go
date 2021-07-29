package main

import "fmt"

func main() {
	names := []string{"hyunsu", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}
