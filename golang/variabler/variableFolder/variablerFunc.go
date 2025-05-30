package variableOpgave

import "fmt"

func RectangleArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func RenteRegning(pmt float64, rente float64, aar int) float64 {
	fremFaktor := 1 + rente
	fmt.Println(fremFaktor)

	for i := 0; i < aar-1; i++ {
		fremFaktor *= 1 + rente
		fmt.Println("Inde i for loop ", fremFaktor)
	}
	fmt.Println(fremFaktor)

	return pmt * fremFaktor

}

func RenteRegningRecur(pmt float64, rente float64, aar int) float64 {
	return recurHelp(pmt, 1+rente, 1, aar)

}

func recurHelp(pmt float64, rente1 float64, rente2 float64, aar int) float64 {
	if aar == 0 {
		return pmt * rente2
	}
	rente2 = rente1 * rente2
	fmt.Println("Dette er rente i rekursiv kald ", rente2)
	return recurHelp(pmt, rente1, rente2, aar-1)

}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}
