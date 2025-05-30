package array

func Sum(numbers [5]int) int {
	sum := 0
	//_, er blank identifier hvis index er ligegyldig fx
	for _, number := range numbers {
		sum += number
	}
	return sum
}
