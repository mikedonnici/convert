package convert

import (
	"fmt"
)

// MassPerArea is a compound unit of mass per area
type MassPerArea string

// todo: probably remove these - unused but here for reference - too many combinations to list so we can rely on string
//
//	representations of compound units for the time being.
const (
	TonsPerAcre              MassPerArea = "ton1ac-1"
	PoundsPerAcre            MassPerArea = "lb1ac-1"
	OuncesMassPerAcre        MassPerArea = "ozm1ac-1"
	OuncesMassPerSquareFoot  MassPerArea = "ozm1[ft2]-1"
	TonnesPerHectare         MassPerArea = "t1ha-1"
	KilogramsPerHectare      MassPerArea = "kg1ha-1"
	KilogramsPerSquareMetre  MassPerArea = "kg1[m2]"
	GramsPerHectare          MassPerArea = "g1ha-1"
	GramsPerSquareMetre      MassPerArea = "g1[m2]-1"
	DecigramsPerHectare      MassPerArea = "dg1ha-1"
	DecigramsPerSquareMetre  MassPerArea = "dg1[m2]-1"
	MilligramsPerHectare     MassPerArea = "mg1ha-1"
	MilligramsPerSquareMetre MassPerArea = "mg1[m2]-1"
)

// MassAreaRatioUnit is a divisive unit with a mass numerator and an area denominator
type MassAreaRatioUnit struct {
	Numerator   MassUnit
	Denominator AreaUnit
}

// String returns the string representation of the MassAreaRatioUnit.
func (u MassAreaRatioUnit) String() string {
	return RatioUnit{
		Numerator:   u.Numerator,
		Denominator: u.Denominator,
	}.String()
}

// massAreaRatioUnitByName attempts to derive a MassAreaRatioUnit from a string.
func massAreaRatioUnitByName(s string) (MassAreaRatioUnit, error) {
	n, d, err := splitCompoundUnit(s)
	if err != nil {
		return MassAreaRatioUnit{}, err
	}
	un, err := massUnitByName(n)
	if err != nil {
		return MassAreaRatioUnit{}, fmt.Errorf("numerator of compound unit %s (%s) is not a mass unit", s, n)
	}
	ud, err := areaUnitByName(d)
	if err != nil {
		return MassAreaRatioUnit{}, fmt.Errorf("denominator of compound unit %s (%s) is not an area unit", s, d)
	}
	return MassAreaRatioUnit{
		Numerator:   un,
		Denominator: ud,
	}, nil
}

// MassAreaRatioMeasure represents a mass measurement per unit area
type MassAreaRatioMeasure struct {
	MassMeasurement
	unitArea AreaUnit
}

// NewMassAreaRatioMeasure creates a new MassAreaRatioMeasure with the specified value and units.
func NewMassAreaRatioMeasure(v float64, mu MassUnit, au AreaUnit) MassAreaRatioMeasure {
	return MassAreaRatioMeasure{
		MassMeasurement: MassMeasurement{
			Value: v,
			Unit:  mu,
		},
		unitArea: au,
	}
}

// NewMassAreaRatioMeasureFromUnitString creates a new MassAreaRatioMeasure with the specified value and compound unit.
func NewMassAreaRatioMeasureFromUnitString(v float64, compoundUnit string) (MassAreaRatioMeasure, error) {
	n, d, err := splitCompoundUnit(compoundUnit)
	if err != nil {
		return MassAreaRatioMeasure{}, err
	}
	un, err := massUnitByName(n)
	if err != nil {
		return MassAreaRatioMeasure{}, fmt.Errorf("numerator of compound unit %s (%s) is not a MassUnit", compoundUnit, n)
	}
	ud, err := areaUnitByName(d)
	if err != nil {
		return MassAreaRatioMeasure{}, fmt.Errorf("denominator of compound unit %s (%s) is not an AreaUnit", compoundUnit, d)
	}
	return NewMassAreaRatioMeasure(v, un, ud), nil
}

// To converts the MassAreaRatioMeasure to the specified mass and area units
func (mr MassAreaRatioMeasure) To(toMassUnit MassUnit, toAreaUnit AreaUnit) MassAreaRatioMeasure {
	toMass := mr.MassMeasurement.To(toMassUnit)
	toArea := AreaMeasurement{Value: 1, Unit: mr.unitArea}.To(toAreaUnit)
	toMass.Value = toMass.Value / toArea.Value
	return MassAreaRatioMeasure{
		MassMeasurement: toMass,
		unitArea:        toAreaUnit,
	}
}

// Value returns the value of the MassAreaRatioMeasure
func (mr MassAreaRatioMeasure) Value() float64 {
	return mr.MassMeasurement.Value
}

// Unit returns the unit of the MassAreaRatioMeasure
// TODO: this should return a compound unit
func (mr MassAreaRatioMeasure) Unit() (string, error) {
	return joinCompoundUnit(mr.MassMeasurement.Unit.String(), mr.unitArea.standard.String())
}
