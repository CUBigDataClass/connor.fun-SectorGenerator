package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

/* The generator object will save all previously calculated locations. */
type Generator struct {
	locData []LocationData
}

/* The exported, generic location data struct */
type LocationData struct {
	Name      string  `json:"name"`
	ID        string  `json:"ID"`
	CenterLat float64 `json:"centerLat"`
	CenterLon float64 `json:"centerLon"`
	North     float64 `json:"north"`
	East      float64 `json:"east"`
	South     float64 `json:"south"`
	West      float64 `json:"west"`
}

func main() {
	var gen = NewGenerator()
	// gen.GeneratePoints("United States", "USA", 37.8283, -96.5795,
	// 	815, 1400, 375, 600)
	// gen.GeneratePoints("United States", "USA", 37, -96.5795,
	// 	825, 1425, 825, 650)
	gen.GeneratePoints("United States", "USA", 37, -96.5795,
		850, 1425, 850, 1600)
	val, _ := gen.GetLocationDataJSON()
	fmt.Println(string(val))
}

/* Returns a new generator */
func NewGenerator() *Generator {
	return &Generator{
		locData: nil,
	}
}

/* Returns the generator's data */
func (gen *Generator) GetLocationData() []LocationData {
	return gen.locData
}

/*  This takes the size of a sector, the centerpoint of a city, the radius from the center, and the city's name.  It
 *  calculates a grid of bounding boxes with unique identifiers and stores them in the Generator object.
 */
func (gen *Generator) GeneratePoints(name string, abbreviation string, centerLat float64, centerLon float64,
	radialHeight float64, radialWidth float64, sectorHeight float64, sectorWidth float64) {
	numRows := math.Ceil((radialHeight * 2) / sectorHeight)
	numCols := math.Ceil((radialWidth * 2) / sectorWidth)

	latChange := calcLatitudeChange(sectorHeight)
	lonChange := calcLongitudeChange(sectorWidth, centerLat)

	initialLat := centerLat - calcLatitudeChange((numRows/2)*sectorHeight)
	initialLon := centerLon - calcLongitudeChange((numCols/2)*sectorWidth, centerLat)

	for i := 0; i < int(numRows); i++ {
		for j := 0; j < int(numCols); j++ {
			north := initialLat + (float64(i+1) * latChange)
			south := initialLat + (float64(i) * latChange)
			east := initialLon + (float64(j) * lonChange)
			west := initialLon + (float64(j+1) * lonChange)
			newBox := LocationData{
				Name:      name,
				ID:        abbreviation + strconv.Itoa(j+(int(numCols)*i)),
				CenterLat: (north + south) / 2,
				CenterLon: (east + west) / 2,
				North:     north,
				South:     south,
				East:      east,
				West:      west,
			}
			gen.locData = append(gen.locData, newBox)
		}
	}
}

/* nice */
func calcLatitudeChange(distance float64) float64 {
	return distance / 69
}

/* Longitude calculations are needlessly involved. */
func calcLongitudeChange(distance float64, lat float64) float64 {
	earthRadius := 3960.0
	radiansToDeg := 180 / math.Pi
	degToRadians := 1 / radiansToDeg
	r := earthRadius * math.Cos(lat*degToRadians)
	return (distance / r) * radiansToDeg
}

func (gen *Generator) GetLocationDataJSON() ([]byte, error) {
	object, err := json.MarshalIndent(&gen.locData, "", "	")
	return object, err
}

func (gen *Generator) ParseLocationDataJSON(raw []byte) error {
	return json.Unmarshal(raw, &gen.locData)
}
