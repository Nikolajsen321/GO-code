package variableOpgave

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	//Arrange
	rect := Rectangle{Height: 10, Width: 5}
	//Act
	got := RectangleArea(rect)

	//Assert
	want := 50.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

/*
Test Noter:
 1. Kør tests i aktuel mappe
	[go test]
	Kører tests i den nuværende mappe
	Fungerer kun hvis du står i fx /variableFolder

2. Kør tests i hele projektet (rekursivt)
    [go test ./...]
	Kører alle tests i alle mapper
	Bruges typisk fra roden af projektet


3. Kør tests i en specifik mappe
	[go test ./variableFolder]
	Kører kun tests i variableFolder-mappen

4. Kør én bestemt testfunktion
	[go test -run TestRectangleArea]
	[go test -v]
	[go test -v -run TestRectangleArea]
	Kører kun testfunktionen med det navn
	Virker også med -v for at se output









*/
