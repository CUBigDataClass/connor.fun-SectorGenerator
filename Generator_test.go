package SectorGenerator

import (
	"testing"
	"fmt"
)

var gen = NewGenerator()

func TestGenerator_GeneratePoints(t *testing.T) {
	gen.GeneratePoints(10, 100, 100, 15, "Test")
}

func TestGenerator_GetLocationDataJSON(t *testing.T) {
	object, _ := gen.GetLocationDataJSON()
	fmt.Println(string(object))
}