package convert

import (
	"fmt"
	"strings"
)

type Volume string

const (
	MicrolitreStandard      Volume = "ul"
	MillilitreStandard      Volume = "ml"
	CentilitreStandard      Volume = "cl"
	DecilitreStandard       Volume = "dl"
	LitreStandard           Volume = "l"
	KilolitreStandard       Volume = "kl"
	DecalitreStandard       Volume = "dal"
	HectolitreStandard      Volume = "hl"
	MegalitreStandard       Volume = "Ml"
	CubicCentimetreStandard Volume = "cm3"   // cubic centimetres
	CubicMetreStandard      Volume = "m3"    // cubic metres
	GallonStandard          Volume = "gal"   // US gallon
	FluidOunceStandard      Volume = "floz"  // fluid ounce
	QuartStandard           Volume = "qt"    // quarts
	PintStandard            Volume = "pt"    // pints
	CubicInchStandard       Volume = "in3"   // cubic inches
	CubicFootStandard       Volume = "ft3"   // cubic feet
	CubicYardStandard       Volume = "yd3"   // cubic yards
	AcreFootStandard        Volume = "ac-ft" // acre-feet
	AcreInchStandard        Volume = "ac-in" // acre-inches
	BushelStandard          Volume = "bu"    // grain
	BaleStandard            Volume = "bale"  // cotton
)

// String return the string representation of the volume unit
func (v Volume) String() string {
	return string(v)
}

type VolumeUnit struct {
	unit       Volume
	full       string
	fancy      string
	aliases    []string
	conversion float64
}

// String returns the string representation of the base volume unit.
func (u VolumeUnit) String() string {
	return u.unit.String()
}

// Matches returns true if s matches the volume unit - this must be case-sensitive because ml is not Ml.
func (u VolumeUnit) Matches(s string) bool {
	// Deal with the special case of Ml (megalitre) and ml (millilitre).
	if u.unit == MillilitreStandard && s == "Ml" {
		return false
	}
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

var volumeUnits = []VolumeUnit{
	Microlitre,
	Millilitre,
	Centilitre,
	Decilitre,
	Litre,
	Kilolitre,
	Decalitre,
	Hectolitre,
	Megalitre,
	CubicCentimetre,
	CubicMetre,
	Gallon,
	FluidOunce,
	Quart,
	Pint,
	CubicInch,
	CubicFoot,
	CubicYard,
	AcreFoot,
	AcreInch,
	Bushel,
	Bale,
}

// var volumeToLitres = map[Volume]float64{
// 	Microlitre:      0.000001,
// 	Millilitre:      0.001,
// 	Centilitre:      0.01,
// 	Decilitre:       0.1,
// 	Litre:           1,
// 	Kilolitre:       1000,
// 	Decalitre:       10,
// 	Hectolitre:      100,
// 	Megalitre:       1000000,
// 	CubicCentimetre: 0.001,
// 	CubicMetre:      1000,
// 	Gallon:          3.78541,
// 	FluidOunce:      0.0284131,
// 	Quart:           0.946353,
// 	Pint:            0.473176,
// 	CubicInch:       0.0163871,
// 	CubicFoot:       28.3168,
// 	CubicYard:       764.555,
// 	AcreFoot:        1233480,
// 	AcreInch:        102790.15312896,
// 	Bushel:          35.2391,
// 	Bale:            480.0, // ref: https://www.cotton.org/tech/bale/bale-description.cfm
// }

var Microlitre = VolumeUnit{
	unit:       MicrolitreStandard,
	full:       "microlitre",
	fancy:      string(MicrolitreStandard),
	aliases:    []string{"microlitres", "microliter", "microliters", "µl", "mcL"},
	conversion: 0.000001,
}

var Millilitre = VolumeUnit{
	unit:       MillilitreStandard,
	full:       "millilitre",
	fancy:      string(MillilitreStandard),
	aliases:    []string{"millilitres", "milliliter", "milliliters"},
	conversion: 0.001,
}

var Centilitre = VolumeUnit{
	unit:       CentilitreStandard,
	full:       "centilitre",
	fancy:      string(CentilitreStandard),
	aliases:    []string{"centilitres", "centiliter", "centiliters"},
	conversion: 0.01,
}

var Decilitre = VolumeUnit{
	unit:       DecilitreStandard,
	full:       "decilitre",
	fancy:      string(DecilitreStandard),
	aliases:    []string{"decilitres", "deciliter", "deciliters"},
	conversion: 0.1,
}

var Litre = VolumeUnit{
	unit:       LitreStandard,
	full:       "litre",
	fancy:      string(LitreStandard),
	aliases:    []string{"litres", "liter", "liters"},
	conversion: 1,
}

var Kilolitre = VolumeUnit{
	unit:       KilolitreStandard,
	full:       "kilolitre",
	fancy:      string(KilolitreStandard),
	aliases:    []string{"kilolitres", "kiloliter", "kiloliters"},
	conversion: 1000,
}

var Decalitre = VolumeUnit{
	unit:       DecalitreStandard,
	full:       "decalitre",
	fancy:      string(DecalitreStandard),
	aliases:    []string{"decalitres", "decaliter", "decaliters"},
	conversion: 10,
}

var Hectolitre = VolumeUnit{
	unit:       HectolitreStandard,
	full:       "hectolitre",
	fancy:      string(HectolitreStandard),
	aliases:    []string{"hectolitres", "hectoliter", "hectoliters", "100l", "100 litres", "100 liters", "100 litre", "100 liter"},
	conversion: 100,
}

var Megalitre = VolumeUnit{
	unit:       MegalitreStandard,
	full:       "megalitre",
	fancy:      string(MegalitreStandard),
	aliases:    []string{"megalitre", "megalitres", "megaliter", "megaliters"},
	conversion: 1000000,
}

var CubicCentimetre = VolumeUnit{
	unit:       CubicCentimetreStandard,
	full:       "cubic centimetre",
	fancy:      string(CubicCentimetreStandard),
	aliases:    []string{"cm^3", "cubic centimetres", "cubic centimeter", "cubic centimeters", "cc"},
	conversion: 0.001,
}

var CubicMetre = VolumeUnit{
	unit:       CubicMetreStandard,
	full:       "cubic metre",
	fancy:      string(CubicMetreStandard),
	aliases:    []string{"m^3", "cubic metres", "cubic meter", "cubic meters"},
	conversion: 1000,
}

var Gallon = VolumeUnit{
	unit:       GallonStandard,
	full:       "gallon",
	fancy:      string(GallonStandard),
	aliases:    []string{"us gal", "us-gal", "us gallon", "us gallons", "gallon", "gallons"},
	conversion: 3.78541,
}

var FluidOunce = VolumeUnit{
	unit:       FluidOunceStandard,
	full:       "fluid ounce",
	fancy:      string(FluidOunceStandard),
	aliases:    []string{"fl oz", "us fl oz", "us-fluid-ounce", "us fluid ounce", "us fluid ounces", "fluid ounce", "fluid ounces"},
	conversion: 0.0295735,
}

var Quart = VolumeUnit{
	unit:       QuartStandard,
	full:       "quart",
	fancy:      string(QuartStandard),
	aliases:    []string{"us qt", "us-quart", "us quarts", "quart", "quarts"},
	conversion: 0.946353,
}

var Pint = VolumeUnit{
	unit:       PintStandard,
	full:       "pint",
	fancy:      string(PintStandard),
	aliases:    []string{"us pt", "us-pint", "us pints", "pint", "pints"},
	conversion: 0.473176,
}

var CubicInch = VolumeUnit{
	unit:       CubicInchStandard,
	full:       "cubic inch",
	fancy:      string(CubicInchStandard),
	aliases:    []string{"in^3", "cubic inches"},
	conversion: 0.0163871,
}

var CubicFoot = VolumeUnit{
	unit:       CubicFootStandard,
	full:       "cubic foot",
	fancy:      string(CubicFootStandard),
	aliases:    []string{"ft^3", "cubic feet"},
	conversion: 28.3168,
}

var CubicYard = VolumeUnit{
	unit:       CubicYardStandard,
	full:       "cubic yard",
	fancy:      string(CubicYardStandard),
	aliases:    []string{"yd^3", "cubic yards"},
	conversion: 764.555,
}

var AcreInch = VolumeUnit{
	unit:       AcreInchStandard,
	full:       "acre inch",
	fancy:      string(AcreInchStandard),
	aliases:    []string{"acre inches"},
	conversion: 102790.15312896,
}

var AcreFoot = VolumeUnit{
	unit:       AcreFootStandard,
	full:       "acre foot",
	fancy:      string(AcreFootStandard),
	aliases:    []string{"acre feet"},
	conversion: 1233480,
}

var Bushel = VolumeUnit{
	unit:       BushelStandard,
	full:       "bushel",
	fancy:      string(BushelStandard),
	aliases:    []string{"bushels"},
	conversion: 35.2391,
}

var Bale = VolumeUnit{
	unit:       BaleStandard,
	full:       "bale",
	fancy:      string(BaleStandard),
	aliases:    []string{"bales"},
	conversion: 480,
}

// volumeUnitByName returns the first volume unit that is a case-sensitive match for s, or an error if no match is found.
func volumeUnitByName(s string) (VolumeUnit, error) {
	for _, u := range volumeUnits {
		if u.Matches(s) {
			return u, nil
		}
	}
	return VolumeUnit{}, fmt.Errorf("no volume unit found for %s", s)
}

// VolumeMeasurement represents a volume measurement.
type VolumeMeasurement struct {
	Value float64
	Unit  VolumeUnit
}

// To converts a volume measurement to the specified unit.
func (m VolumeMeasurement) To(unit VolumeUnit) VolumeMeasurement {
	if m.Value != 0 {
		m.Value = (m.Value * m.Unit.conversion) / unit.conversion
	}
	m.Unit = unit
	return m
}
