package SectorGenerator

import (
	"bytes"
	"fmt"
	"testing"
)


func TestGenerator_GeneratePoints(t *testing.T) {
	var gen = NewGenerator()
	gen.GeneratePoints(10, 100, 100, 15, "Test", "T")
}

func TestGenerator_GetLocationDataJSON(t *testing.T) {
	var gen = NewGenerator()
	gen.GeneratePoints(20, 40.0150, -105.2705, 1, "Boulder", "BLD")
	gen.GeneratePoints(20, 40.7128, -74.0060, 1, "New York", "NY")
	gen.GeneratePoints(20, 34.0522, -118.2437, 1, "Los Angeles", "LA")
	gen.GeneratePoints(20, 47.6062, -122.3321, 1, "Seattle", "SEA")
	object, _ := gen.GetLocationDataJSON()
	fmt.Println(string(object))
}

func TestGenerator_ParseLocationDataJSON(t *testing.T) {
	var gen0 = NewGenerator()
	gen0.GeneratePoints(10, 100, 100, 15, "Test", "T")
	var gen1 = NewGenerator()
	object0, _ := gen0.GetLocationDataJSON()
	gen1.ParseLocationDataJSON(object0)
	object1, _ := gen1.GetLocationDataJSON()
	fmt.Println(object0)
	fmt.Println(object1)

	if !bytes.Equal(object0, object1) {
		t.Fatalf("JSON byte slices were not equal.")
	}
}
