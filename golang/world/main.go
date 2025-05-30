package main

import (
	"fmt"
	"math"
	"slices"
)

const s string = "constant"

func main() {
	fmt.Println("Hello from world mappe")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// fmt.Println(true && true)

	// fmt.Println(true || false)

	// fmt.Println(!true)

	var a = "Initial"
	fmt.Println(a)
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(n)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for i := range 3 {
		fmt.Println("range", i)

	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	if 8%4 == 0 {
		fmt.Println("8 er dividere bar med 4")
	}

	if 7%2 == 1 || 8%2 == 0 {
		fmt.Println("HEJ1")
	} else if 3+3 == 6 {
		fmt.Println("HEJ2")
	}

	var m int8 = 2
	switch m {
	case 1:
		fmt.Println("et")
	case int8(math.Pow(2, 7) - 1):
		fmt.Println("Maksimal positiv")
	}

	var ad [5]int
	fmt.Println("emp", ad)

	ad[4] = 100
	fmt.Println("Set", ad)
	fmt.Println("get:", ad[4])

	fmt.Println("len:", len(a)) //Retunere længde af array

	b := [5]int{1, 2, 3, 4, 5} //delcare og intalizer i en linje
	fmt.Println("dcl:", b)
	// compiler tæller antal af elmenter
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//specficer index med :, vil elementer mellem blive 0
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	//multi dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//intalseret med det samme
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	//Slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] //giver elementerne mellem lavest 2 og 5  aktså s[2] s[3] s[4]
	fmt.Println("sl1:", l)

	l = s[:5] //giver op til men ekskludere s[5]
	fmt.Println("sl2:", l)
	l = s[2:] //op til og inkludere s[2]
	fmt.Println("sl3:", l)

	//Declare
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//utility
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// Initialiserer hver indre slice
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD2)

	//Map
}
