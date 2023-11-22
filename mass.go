package convert

import (
	"fmt"
	"strings"
)

type Mass string

const (
	MilligramStandard Mass = "mg"
	DecigramStandard  Mass = "dg"
	GramStandard      Mass = "g"
	KilogramStandard  Mass = "kg"
	TonneStandard     Mass = "t"
	PoundStandard     Mass = "lb"
	OunceMassStandard Mass = "ozm"
	StoneStandard     Mass = "st"
	TonStandard       Mass = "ton"
)

// String returns the string representation of the mass unit.
func (m Mass) String() string {
	return string(m)
}

// MassUnit represents a mass unit.
type MassUnit struct {
	unit       Mass
	full       string
	fancy      string
	aliases    []string
	conversion float64
}

// String returns the string representation of the base mass unit.
func (u MassUnit) String() string {
	return u.unit.String()
}

// Matches returns true if s matches the mass unit.
func (u MassUnit) Matches(s string) bool {
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

var massUnits = []MassUnit{
	Milligram,
	Decigram,
	Gram,
	Kilogram,
	Tonne,
	Pound,
	OunceMass,
	Stone,
	Ton,
}

var Milligram = MassUnit{
	unit:  MilligramStandard,
	full:  "milligram",
	fancy: string(MilligramStandard),
	aliases: []string{
		"milligram",
		"milligrams",
		"mil",
		"mils",
	},
	conversion: 0.001,
}

var Decigram = MassUnit{
	unit:  DecigramStandard,
	full:  "decigram",
	fancy: string(DecigramStandard),
	aliases: []string{
		"decigrams",
	},
	conversion: 0.1,
}

var Gram = MassUnit{
	unit:  GramStandard,
	full:  "gram",
	fancy: string(GramStandard),
	aliases: []string{
		"grams",
	},
	conversion: 1,
}

var Kilogram = MassUnit{
	unit:  KilogramStandard,
	full:  "kilogram",
	fancy: string(KilogramStandard),
	aliases: []string{
		"kilograms",
		"kilo",
		"kilos",
	},
	conversion: 1000,
}

var Tonne = MassUnit{
	unit:  TonneStandard,
	full:  "tonne",
	fancy: string(TonneStandard),
	aliases: []string{
		"tonnes",
		"metric ton",
		"metric tons",
		"metric tonne",
		"metric tonnes",
	},
	conversion: 1000000,
}

var OunceMass = MassUnit{
	unit:  OunceMassStandard,
	full:  "ounce mass",
	fancy: string(OunceMassStandard),
	aliases: []string{
		"ounce",
		"ounces",
		"oz",
	},
	conversion: 28.3495,
}

var Pound = MassUnit{
	unit:  PoundStandard,
	full:  "pound",
	fancy: string(PoundStandard),
	aliases: []string{
		"pounds",
		"lbs",
	},
	conversion: 453.592,
}

var Stone = MassUnit{
	unit:  StoneStandard,
	full:  "stone",
	fancy: string(StoneStandard),
	aliases: []string{
		"stones",
	},
	conversion: 6350.29,
}

var Ton = MassUnit{
	unit:  TonStandard,
	full:  "ton",
	fancy: string(TonStandard),
	aliases: []string{
		"tons",
		"short ton",
		"short tons",
	},
	conversion: 907185,
}

// massUnitByName returns the first mass unit that matches s.
func massUnitByName(s string) (MassUnit, error) {
	for _, u := range massUnits {
		if u.Matches(s) {
			return u, nil
		}
	}
	return MassUnit{}, fmt.Errorf("no mass unit found for %s", s)
}

// MassMeasurement represents a mass measurement.
type MassMeasurement struct {
	Value float64
	Unit  MassUnit
}

// To converts a mass measurement To the specified unit.
func (m MassMeasurement) To(unit MassUnit) MassMeasurement {
	if m.Value != 0 {
		m.Value = (m.Value * m.Unit.conversion) / unit.conversion
	}
	m.Unit = unit
	return m
}
