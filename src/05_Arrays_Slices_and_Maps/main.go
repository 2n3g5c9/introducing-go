package main

import "fmt"

func main() {
	// Arrays
	var a [5]int
	a[4] = 100
	fmt.Println(a)

	var b[5]float64
	b[0] = 98
	b[1] = 93
	b[2] = 77
	b[3] = 82
	b[4] = 83

	var total float64 = 0
	for i:= 0; i < 5; i++ {
		total += b[i]
	}
	fmt.Println(total / 5)

	fmt.Println(total / float64(len(b)))

	total = 0
	for _, value := range b {
		total += value
	}
	fmt.Println(total / float64(len(b)))

	c := [5]float64{ 98, 93, 77, 82, 83 }
	d := [5]float64{
		98,
		93,
		77,
		83,
		83,
	}
	fmt.Println(c, d)

	// Slices
	var e []float64
	fmt.Println(e)

	f := make([]float64, 5)
	fmt.Println(f)

	g := make([]float64, 5, 10)
	fmt.Println(g)

	arr := [5]float64{1, 2, 3, 4, 5}
	h := arr[0:5]
	fmt.Println(h)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	//Maps
	i := make(map[string]int)
	i["key"] = 10
	fmt.Println(i["key"])

	j := make(map[int]int)
	j[1] = 10
	fmt.Println(j[1])

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

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	shorter_elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(shorter_elements["Li"])

	detailed_elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := detailed_elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}