package convert

import (
	"fmt"
	"strings"
)

type Line string

const (
	MillimetreStandard Line = "mm"
	CentimetreStandard Line = "cm"
	MetreStandard      Line = "m"
	KilometreStandard  Line = "km"
	InchStandard       Line = "in"
	FootStandard       Line = "ft"
	YardStandard       Line = "yd"
	MileStandard       Line = "mi"
)

// String returns the string representation of the line unit.
func (l Line) String() string {
	return string(l)
}

// LineUnit contains the data for a line unit.
type LineUnit struct {
	unit       Line
	full       string
	fancy      string
	aliases    []string
	conversion float64
}

// String returns the string representation of the base line unit.
func (u LineUnit) String() string {
	return u.unit.String()
}

// Matches returns true if s matches the line unit.
func (u LineUnit) Matches(s string) bool {
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

var lineUnits = []LineUnit{
	Millimetre,
	Centimetre,
	Metre,
	Kilometre,
	Inch,
	Foot,
	Yard,
	Mile,
}

var Millimetre = LineUnit{
	unit:  MillimetreStandard,
	full:  "millimetre",
	fancy: "millimetre",
	aliases: []string{
		"millimeter",
		"millimeters",
		"millimetres",
	},
	conversion: 0.001,
}

var Centimetre = LineUnit{
	unit:  CentimetreStandard,
	full:  "centimetre",
	fancy: "centimetre",
	aliases: []string{
		"centimeter",
		"centimeters",
		"centimetres",
	},
	conversion: 0.01,
}

var Metre = LineUnit{
	unit:  MetreStandard,
	full:  "metre",
	fancy: "metre",
	aliases: []string{
		"meter",
		"meters",
		"metres",
	},
	conversion: 1,
}

var Kilometre = LineUnit{
	unit:  KilometreStandard,
	full:  "kilometre",
	fancy: "kilometre",
	aliases: []string{
		"kilometer",
		"kilometers",
		"kilometres",
	},
	conversion: 1000,
}

var Inch = LineUnit{
	unit:  InchStandard,
	full:  "inch",
	fancy: "inch",
	aliases: []string{
		"inches",
	},
	conversion: 0.0254,
}

var Foot = LineUnit{
	unit:  FootStandard,
	full:  "foot",
	fancy: "foot",
	aliases: []string{
		"feet",
	},
	conversion: 0.3048,
}

var Yard = LineUnit{
	unit:  YardStandard,
	full:  "yard",
	fancy: "yard",
	aliases: []string{
		"yards",
	},
	conversion: 0.9144,
}

var Mile = LineUnit{
	unit:  MileStandard,
	full:  "mile",
	fancy: "mile",
	aliases: []string{
		"miles",
	},
	conversion: 1609.34,
}

// lineUnitByName returns the first lineUnit that matches the search string, or nil if no match is found.
func lineUnitByName(s string) (LineUnit, error) {
	for _, u := range lineUnits {
		if u.Matches(s) {
			return u, nil
		}
	}
	return LineUnit{}, fmt.Errorf("no line unit found for %s", s)
}

// LineMeasurement represents a linear measurement such as height, depth, distance etc.
type LineMeasurement struct {
	Value float64
	Unit  LineUnit
}

// To converts a line measurement to the specified unit.
func (m LineMeasurement) To(unit LineUnit) LineMeasurement {
	if m.Value != 0 {
		m.Value = (m.Value * m.Unit.conversion) / unit.conversion
	}
	m.Unit = unit
	return m
}
