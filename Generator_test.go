package generator

import (
	"fmt"
	"testing"
)

func TestGenerator_GeneratePoints(t *testing.T) {
	var gen = NewGenerator()
	gen.GeneratePoints("United States", "USA", 39.8283, -98.5795,
		791, 1340, 320, 540)
	val, _ := gen.GetLocationDataJSON()
	fmt.Println(string(val))
}
