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
