package Forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Square", func(t *testing.T) {

		s := Square{10, 12}

		expectedArea := float64(120)
		receivedArea := s.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The recevied area [%f] is not equal expectedArea [%f]", receivedArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}

		expectedArea := float64(math.Pi + 100)
		receivedArea := c.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The recevied area [%f] is not equal expectedArea [%f]", receivedArea, expectedArea)
		}
	})
}
