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
func UnitFromLabel(label string) (Unit, error) {
	switch {
	case IsAreaUnit(label):
		return areaUnitByName(label)
	case IsLineUnit(label):
		return lineUnitByName(label)
	case IsMassUnit(label):
		return massUnitByName(label)
	case IsTimeUnit(label):
		return timeUnitByName(label)
	case IsVolumeUnit(label):
		return volumeUnitByName(label)
	case IsMassAreaRatioUnit(label):
		return massAreaRatioUnitByName(label)
	case IsVolumeAreaRatioUnit(label):
		return volumeAreaRatioUnitByName(label)
	default:
		return nil, fmt.Errorf("unknown unit: %s", label)
	}
}

// StandardLabel returns a 'standard' label for the specified unit label
func StandardLabel(label string) (string, error) {
	if label == "" {
		return "", nil
	}
	u, err := UnitFromLabel(label)
	if err != nil {
		return "", fmt.Errorf("failed to get unit from label %s: %w", label, err)
	}
	return u.String(), nil
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