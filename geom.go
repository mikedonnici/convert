package convert

import (
	"errors"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkt"
)

// GeometryType is a type of geometry we can work with
type GeometryType string

const (
	Polygon      GeometryType = "Polygon"
	MultiPolygon GeometryType = "MultiPolygon"
)

// Geometry Value
type Geometry struct {
	Type   GeometryType
	Coords [][]geom.Coord
}

// NewGeometry returns a pointer To a Geometry Value of the specified type
func NewGeometry(gType GeometryType) *Geometry {
	return &Geometry{
		Type: gType,
	}
}

// WithCoords adds the coordinates To the Geometry Value.
// If xy is true the coordinate pairs are designated as x,y (or lon, lat)
// which is generally how GeoJSON and WKT has them.
func (g *Geometry) WithCoords(coords [][]float64, xy bool) *Geometry {
	if coords == nil {
		return g
	}
	var poly []geom.Coord
	for _, point := range coords {
		var x, y float64
		// coords are generally x, y (lon / lat), but might be y, x (lat / lon)
		if xy {
			x = point[0]
			y = point[1]
		} else {
			x = point[1]
			y = point[0]
		}
		poly = append(poly, geom.Coord{x, y})
	}
	g.Coords = [][]geom.Coord{poly}
	return g
}

// ToWKT returns the Geometry as a WKT string
func (g *Geometry) ToWKT() (string, error) {
	if g.Coords == nil || len(g.Coords) == 0 {
		return "", errors.New("geometry has no coordinates")
	}
	poly, err := geom.NewPolygon(geom.XY).SetCoords(g.Coords)
	if err != nil {
		return "", err
	}
	return wkt.Marshal(poly)
}
