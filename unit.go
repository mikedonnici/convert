// Package unit provides functions for handling units of measurement and conversions
package convert

import (
	"fmt"
	"strings"
)

type Unit interface {
	String() string
}

// RatioUnit is a unit with a numerator and a denominator
type RatioUnit struct {
	Numerator   Unit
	Denominator Unit
}

// String returns the string representation of ratio unit and satisfied the Unit interface.
func (u RatioUnit) String() string {
	n := wrapUnitWithExponent(u.Numerator)
	d := wrapUnitWithExponent(u.Denominator)
	return fmt.Sprintf("%s1%s-1", n, d)
}

// IsAreaUnit returns true if s is a valid area unit.
func IsAreaUnit(s string) bool {
	_, err := areaUnitByName(s)
	return err == nil
}

// IsLineUnit returns true if the given string is a valid line unit.
func IsLineUnit(s string) bool {
	_, err := lineUnitByName(s)
	return err == nil
}

// IsMassUnit returns true if s is a valid mass unit.
func IsMassUnit(s string) bool {
	_, err := massUnitByName(s)
	return err == nil
}

// IsVolumeUnit returns true if s is a valid volume unit.
func IsVolumeUnit(s string) bool {
	_, err := volumeUnitByName(s)
	return err == nil
}

// IsTimeUnit returns true if s is a valid time unit.
func IsTimeUnit(s string) bool {
	_, err := timeUnitByName(s)
	return err == nil
}

// IsMassAreaRatioUnit returns true if the unit arg can be identified as a mass/area, otherwise false.
func IsMassAreaRatioUnit(unit string) bool {
	n, d, err := splitCompoundUnit(unit)
	if err != nil {
		return false
	}
	return IsMassUnit(n) && IsAreaUnit(d)
}

// IsVolumeAreaRatioUnit returns true if the unit arg can be identified as a volume/area, otherwise false.
func IsVolumeAreaRatioUnit(unit string) bool {
	n, d, err := splitCompoundUnit(unit)
	if err != nil {
		return false
	}
	return IsVolumeUnit(n) && IsAreaUnit(d)
}

// UnitFromLabel returns the standard unit for the given unit string.
func UnitFromLabel(unit string) (Unit, error) {
	switch {
	case IsAreaUnit(unit):
		return areaUnitByName(unit)
	case IsLineUnit(unit):
		return lineUnitByName(unit)
	case IsMassUnit(unit):
		return massUnitByName(unit)
	case IsTimeUnit(unit):
		return timeUnitByName(unit)
	case IsVolumeUnit(unit):
		return volumeUnitByName(unit)
	// case IsMassAreaRatioUnit(unit):
	// 	// an arbitrary conversion from/to the same unit should yield the standard unit in the response
	// 	au, err := areaUnitByName(unit)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	u, err := MassPerAreaMeasurement{
	// 		MassMeasurement: MassMeasurement{Unit: Gram},
	// 		unitArea:        au,
	// 	}.Unit()
	// 	m, err := ValueFromTo(1, u, u)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return m.Unit(), nil
	// case IsVolumeAreaRatioUnit(unit):
	// 	return volumeAreaUnitByName(unit)
	default:
		return nil, fmt.Errorf("unknown unit: %s", unit)
	}
}

// wrapUnitWithExponent wraps the unit string with square brackets if it ends with a '2' or '3' or if it has a space.
// This is to avoid confusion in compound units - eg [m3]1ha-1, [fl oz]1[ft3]-1, [m3]1[m2]-1 etc.
func wrapUnitWithExponent(u Unit) string {
	s := u.String()
	if s[len(s)-1:] == "2" || s[len(s)-1:] == "3" || strings.Contains(s, " ") {
		return fmt.Sprintf("[%s]", s)
	}
	return s
}
