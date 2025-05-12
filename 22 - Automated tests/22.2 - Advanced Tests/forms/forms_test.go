package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		// Arrange
		rectangle := Rectangle{10, 12}
		expect := float64(120)

		// Act
		result := rectangle.Area()

		// Assert
		if expect != result {
			t.Fatalf("expect area %f but got %f", expect, result)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		// Arrange
		circle := Circle{10}
		expect := float64(math.Pi * 100)

		// Act
		result := circle.Area()

		// Assert
		if expect != result {
			t.Fatalf("expect area %f but got %f", expect, result)
		}
	})
}
