package convert

import (
	"fmt"
	"strings"
)

// splitValueAreaCompoundUnit separates a compound unit string into numerator and denominator unit strings and verifies
// that the numerator is a mass or volume unit and the denominator is an area unit.
func splitValueAreaCompoundUnit(unit string) (string, string, error) {
	n, d, err := splitCompoundUnit(unit)
	if err != nil {
		return "", "", err
	}
	if !IsMassUnit(n) && !IsVolumeUnit(n) {
		return "", "", fmt.Errorf("compound unit %s has numerator %s, expecting a mass or volume unit", unit, n)
	}
	if !IsAreaUnit(d) {
		return "", "", fmt.Errorf("compound unit %s has denominator %s, expecting an area unit", unit, d)
	}
	return n, d, nil
}

// splitCompoundUnit separates a compound Unit string into numerator and denominator Unit strings.
// For example: "l1ha-1" OR l/ha -> "l", "ha"
func splitCompoundUnit(unit string) (string, string, error) {
	if strings.Contains(unit, "-1") {
		return splitCompoundUnitExponentForm(unit)
	}
	if strings.Contains(unit, "/") {
		return splitCompoundUnitSlashForm(unit)
	}
	return "", "", fmt.Errorf("splitCompoundUnit() expects Unit string in exponent form (eg kg1ha-1) or slash form (eg kg/ha), got %s", unit)
}

func splitCompoundUnitExponentForm(unit string) (string, string, error) {
	s := strings.TrimRight(unit, "-1")
	xs := strings.Split(s, "1")
	if len(xs) != 2 {
		return "", "", fmt.Errorf("compound Unit %s split into %d parts, should be 2", unit, len(xs))
	}

	// Units with an exponent will generally be enclosed in square brackets which need To be removed.
	// For example [m3]1[m2]-1 (cubic metres per square metre) should return "m3" and "m3"
	n := strings.ToLower(strings.TrimSpace(strings.TrimRight(strings.TrimLeft(xs[0], "["), "]")))
	d := strings.ToLower(strings.TrimSpace(strings.TrimRight(strings.TrimLeft(xs[1], "["), "]")))

	if !IsVolumeUnit(n) && !IsMassUnit(n) {
		return "", "", fmt.Errorf("invalid MassMeasurement or volume Unit: %s", n)
	}
	if !IsAreaUnit(d) {
		return "", "", fmt.Errorf("invalid AreaMeasurement Unit: %s", d)
	}
	return n, d, nil
}

func splitCompoundUnitSlashForm(unit string) (string, string, error) {
	xs := strings.Split(unit, "/")
	if len(xs) != 2 {
		return "", "", fmt.Errorf("compound Unit %s split into %d parts, should be 2", unit, len(xs))
	}

	n := strings.ToLower(strings.TrimSpace(xs[0]))
	d := strings.ToLower(strings.TrimSpace(xs[1]))

	if !IsVolumeUnit(n) && !IsMassUnit(n) {
		return "", "", fmt.Errorf("invalid MassMeasurement or volume Unit: %s", n)
	}
	if !IsAreaUnit(d) {
		return "", "", fmt.Errorf("invalid AreaMeasurement Unit: %s", d)
	}
	return n, d, nil
}

// joinCompoundUnit returns numerator and denominator strings joined as a compound Unit string
func joinCompoundUnit(numerator, denominator string) (string, error) {
	if !IsVolumeUnit(numerator) && !IsMassUnit(numerator) {
		return "", fmt.Errorf("invalid MassMeasurement or volume Unit: %s", numerator)
	}
	if !IsAreaUnit(denominator) {
		return "", fmt.Errorf("invalid AreaMeasurement Unit: %s", denominator)
	}

	// Units with an exponent, eg m2, m3, ft2, are wrapped in square brackets.
	wrapUnit := func(s string) string {
		if s[len(s)-1:] == "2" || s[len(s)-1:] == "3" {
			return fmt.Sprintf("[%s]", s)
		}
		return s
	}
	n := wrapUnit(numerator)
	d := wrapUnit(denominator)
	return fmt.Sprintf("%s1%s-1", n, d), nil
}

// maybeCompoundUnit is a helper that returns true if the string appears To be a compound Unit
func maybeCompoundUnit(unit string) bool {
	return strings.Contains(unit, "-1") || strings.Contains(unit, "/")
}
