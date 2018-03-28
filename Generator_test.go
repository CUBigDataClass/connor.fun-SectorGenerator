package SectorGenerator

import "testing"

func TestGenerator_GeneratePoints(t *testing.T) {
	gen := NewGenerator()
	gen.GeneratePoints(10, 100, 100, 15, "Test")
}