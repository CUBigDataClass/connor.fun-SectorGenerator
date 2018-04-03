package SectorGenerator

import (
	"encoding/json"
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
	CenterLon float64 `json:"centerLat"`
	North     float64 `json:"north"`
	East      float64 `json:"east"`
	South     float64 `json:"south"`
	West      float64 `json:"west"`
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
func (gen *Generator) GeneratePoints(sectorSize float64, centerLat float64, centerLon float64,
	radius float64, name string, abbreviation string) {
	numSquares := math.Ceil((radius * 2) / sectorSize)

	latChange := calcLatitudeChange(sectorSize)
	lonChange := calcLongitudeChange(sectorSize, centerLat)

	initialLat := centerLat - calcLatitudeChange((numSquares/2)*sectorSize)
	initialLon := centerLon - calcLongitudeChange((numSquares/2)*sectorSize, centerLat)

	for i := 0; i < int(numSquares); i++ {
		for j := 0; j < int(numSquares); j++ {
			newBox := LocationData{
				Name:      name,
				ID:        abbreviation + strconv.Itoa(j + (int(numSquares) * i)),
				CenterLat: centerLat,
				CenterLon: centerLon,
				North:     initialLat + (float64(i+1) * latChange),
				South:     initialLat + (float64(i) * latChange),
				East:      initialLon + (float64(j) * lonChange),
				West:      initialLon + (float64(j+1) * lonChange),
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
	object, err := json.MarshalIndent(gen.locData, "", "	")
	return object, err
}

func (gen *Generator) ParseLocationDataJSON(raw []byte) (error) {
  return json.Unmarshal(raw, &gen.locData)
}
