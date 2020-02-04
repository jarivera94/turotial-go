package main

import "fmt"

func main() {
	//Arrays
	//var x [5]float64
	x := [5]float64{
		99, 98, 97, 96, 94,
	}
	var total float64
	for _, value := range x {
		total += value
	}

	fmt.Println("The averange is ", total/float64(len(x)))

	//Slices
	slice := make([]float64, 5, 10)
	slice[4] = 10
	fmt.Println(slice)
	//slice from arrays another way to declare, where de first number is de first of copy
	sliceFromArray := x[0:3]
	fmt.Println(sliceFromArray)
	//append function to slice
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6, 7)
	fmt.Println("Slice 1", slice1)
	fmt.Println("Slice 2", slice2)

	//copy functions on slice

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3)
	fmt.Println(slice4)

	//maps

	mapX := make(map[string]int)
	mapX["key1"] = 10
	mapX["key2"] = 20
	fmt.Println(mapX)
	delete(mapX, "key1")
	fmt.Println(mapX)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])

	name, ok := elements["UN"]

	fmt.Println(name, ok)
	if name, ok := elements["UN"]; ok {
		fmt.Println(name, ok)

	}

	if name, ok := elements["Li"]; ok {
		fmt.Println(name, ok)
	}

	temperatureElement := map[string]map[string]string{
		"H": map[string]string{
			"name":  "hidrogen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "hidrogenagaing",
			"state": "gas",
		},
	}

	fmt.Println(temperatureElement)
}
