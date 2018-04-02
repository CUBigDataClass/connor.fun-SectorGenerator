package SectorGenerator

import (
	"bytes"
	"fmt"
	"testing"
)

var gen = NewGenerator()

func TestGenerator_GeneratePoints(t *testing.T) {
	gen.GeneratePoints(10, 100, 100, 15, "Test")
}

func TestGenerator_GetLocationDataJSON(t *testing.T) {
	object, _ := gen.GetLocationDataJSON()
	fmt.Println(string(object))
}

func TestGenerator_ParseLocationDataJSON(t *testing.T) {
	var gen1 = NewGenerator()
	object0, _ := gen.GetLocationDataJSON()
	gen1.ParseLocationDataJSON(object0)
	object1, _ := gen1.GetLocationDataJSON()

	if !bytes.Equal(object0, object1) {
		t.Fatalf("JSON byte slices were not equal.")
	}
}
