package convert

import (
	"fmt"
	"strings"
)

type Area string

const (
	SquareCentimetreStandard Area = "cm2"
	SquareMetreStandard      Area = "m2"
	SquareKilometreStandard  Area = "km2"
	HectareStandard          Area = "ha"
	SquareInchStandard       Area = "in2"
	SquareFootStandard       Area = "ft2"
	SquareYardStandard       Area = "yd2"
	SquareMileStandard       Area = "mi2"
	AcreStandard             Area = "ac"
)

// String returns the string representation of the area unit.
func (a Area) String() string {
	return string(a)
}

// AreaUnit contains the data for a area unit.
type AreaUnit struct {
	standard   Area
	full       string
	fancy      string
	aliases    []string
	conversion float64
}

// String returns the string representation of the base area unit.
func (u AreaUnit) String() string {
	return u.standard.String()
}

// Matches returns true if s matches the area unit.
func (u AreaUnit) Matches(s string) bool {
	if strings.EqualFold(u.String(), s) ||
		strings.EqualFold(u.fancy, s) ||
		strings.EqualFold(u.full, s) {
		return true
	}
	for _, alias := range u.aliases {
		if strings.EqualFold(alias, s) {
			return true
		}
	}
	return false
}

var areaUnits = []AreaUnit{
	SquareCentimetre,
	SquareMetre,
	SquareKilometre,
	Hectare,
	SquareInch,
	SquareFoot,
	SquareYard,
	SquareMile,
	Acre,
}

var SquareCentimetre = AreaUnit{
	standard: SquareCentimetreStandard,
	full:     "square centimetre",
	fancy:    "cm²",
	aliases: []string{
		"cm^2",
		"centimetre squared",
		"centimetres squared",
		"centimeter squared",
		"square centimetres",
		"squared centimeters",
	},
	conversion: 0.0001,
}

var SquareMetre = AreaUnit{
	standard: SquareMetreStandard,
	full:     "square metre",
	fancy:    "m²",
	aliases: []string{
		"m^2",
		"metre squared",
		"metres squared",
		"meter squared",
		"square metres",
		"squared meters",
		"square meter",
		"squared meters",
	},
	conversion: 1,
}

var SquareKilometre = AreaUnit{
	standard: SquareKilometreStandard,
	full:     "square kilometre",
	fancy:    "km²",
	aliases: []string{
		"km^2",
		"kilometre squared",
		"kilometres squared",
		"kilometer squared",
		"square kilometres",
		"squared kilometers",
	},
	conversion: 1000000,
}

var Hectare = AreaUnit{
	standard: HectareStandard,
	full:     "hectare",
	fancy:    "ha",
	aliases: []string{
		"hectares",
	},
	conversion: 10000,
}

var SquareInch = AreaUnit{
	standard: SquareInchStandard,
	full:     "square inch",
	fancy:    "in²",
	aliases: []string{
		"in^2",
		"inch squared",
		"inches squared",
		"square inches",
	},
	conversion: 0.00064516,
}

var SquareFoot = AreaUnit{
	standard: SquareFootStandard,
	full:     "square foot",
	fancy:    "ft²",
	aliases: []string{
		"ft^2",
		"foot squared",
		"feet squared",
		"square feet",
	},
	conversion: 0.092903,
}

var SquareYard = AreaUnit{
	standard: SquareYardStandard,
	full:     "square yard",
	fancy:    "yd²",
	aliases: []string{
		"yd^2",
		"yard squared",
		"yards squared",
		"square yards",
	},
	conversion: 0.836127,
}

var SquareMile = AreaUnit{
	standard: SquareMileStandard,
	full:     "square mile",
	fancy:    "mi²",
	aliases: []string{
		"mi^2",
		"mile squared",
		"miles squared",
		"square miles",
	},
	conversion: 2589988.11,
}

var Acre = AreaUnit{
	standard: AcreStandard,
	full:     "acre",
	fancy:    "ac",
	aliases: []string{
		"acres",
	},
	conversion: 4046.86,
}

// areaUnitByName returns the first areaUnit that matches the search string, or nil if no match is found.
func areaUnitByName(s string) (AreaUnit, error) {
	for _, u := range areaUnits {
		if u.Matches(s) {
			return u, nil
		}
	}
	return AreaUnit{}, fmt.Errorf("no area unit found for %s", s)
}

// AreaMeasurement represents an area measurement.
type AreaMeasurement struct {
	Value float64
	Unit  AreaUnit
}

// To converts an area measurement to the specified unit.
func (m AreaMeasurement) To(unit AreaUnit) AreaMeasurement {
	if m.Value != 0 {
		m.Value = (m.Value * m.Unit.conversion) / unit.conversion
	}
	m.Unit = unit
	return m
}
