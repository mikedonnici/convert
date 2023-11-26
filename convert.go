package convert

import (
	"fmt"
	"math"
	"strings"
)

// Two types of value / area conversions
type conversionType int

const (
	massAreaConversion = iota
	volumeAreaConversion
)

// splitConversionUnits splits a compound unit into its numerator and denominator units from a conversion from one
// compound unit to another. For example, if converting from kg/ha to lb1ac-1, the numerator units would be kg and lb,
// and the denominator units would be ha and ac.
type splitConversionUnits struct {
	fromNumerator   string
	fromDenominator string
	toNumerator     string
	toDenominator   string
}

// ValueFromTo converts a numerical value from one unit To another. Params fromUnit and toUnit can be
// simple units such as lb or kg, or compound units such as kg/ha or lb1ac-1.
// It will return an error if fromUnit and toUnit are not compatible for conversion.
func ValueFromTo(value float64, fromUnit string, toUnit string) (float64, error) {
	if fromUnit == toUnit {
		return value, nil
	}
	fn := conversionFunc(fromUnit, toUnit)
	if fn == nil {
		return 0, fmt.Errorf("cannot convert from %s to %s", fromUnit, toUnit)
	}
	return fn(value, fromUnit, toUnit)
}

// CropRate is a special conversion which can convert a MassMeasurement rate To a volume using known bushel conversions for
// certain crops. If crop Value is not provided it will still do MassMeasurement-MassMeasurement or volume-volume conversions.
func CropRate(crop string, value float64, fromCompoundUnit, toCompoundUnit string) (float64, error) {
	// Nothing To do
	if fromCompoundUnit == toCompoundUnit {
		return value, nil
	}

	// Converting MassMeasurement-MassMeasurement or volume-volume is a straightforward rate conversion
	if (IsMassAreaRatioUnit(fromCompoundUnit) && IsMassUnit(toCompoundUnit)) ||
		(IsVolumeAreaRatioUnit(fromCompoundUnit) && IsVolumeAreaRatioUnit(toCompoundUnit)) {
		return convertMassAreaMeasurement(value, fromCompoundUnit, toCompoundUnit)
	}

	// From here on we need To deal with a specific crop
	crop = strings.ToLower(crop)
	if !isBushelCrop(crop) && !isBaleCrop(crop) {
		return 0, fmt.Errorf("unknown crop: %s", crop)
	}
	return convertCropRate(crop, value, fromCompoundUnit, toCompoundUnit)
}

// Round rounds a float64 to the specified number of decimal places
func Round(f float64, places int) float64 {
	n := math.Pow10(places)
	return math.Round(f*n) / n
}

// convertAreaMeasurement converts an AreaMeasurement from one unit to another. Params fromUnit and toUnit can be
// simple AreaMeasurement units such as ha or ac, or compound units such as kg/ha or lb1ac-1.
func convertAreaMeasurement(areaValue float64, fromUnit string, toUnit string) (float64, error) {
	if fromUnit == toUnit {
		return areaValue, nil
	}
	var err error
	// Get the AreaMeasurement Unit from compound units
	if maybeCompoundUnit(fromUnit) {
		_, fromUnit, err = splitCompoundUnit(fromUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split fromUnit as compound Unit: %s", err)
		}
	}
	if maybeCompoundUnit(toUnit) {
		_, toUnit, err = splitCompoundUnit(toUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split toUnit as compound Unit: %s", err)
		}
	}
	from, err := areaUnitByName(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s is not an AreaUnit", fromUnit)
	}
	to, err := areaUnitByName(toUnit)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s is not an AreaUnit", toUnit)
	}
	a := AreaMeasurement{
		Value: areaValue,
		Unit:  from,
	}
	return a.To(to).Value, nil
}

// convertLineMeasurement converts a LineMeasurement from / to the specified units.
func convertLineMeasurement(value float64, fromUnit, toUnit string) (float64, error) {
	if (fromUnit == toUnit) || value == 0 {
		return value, nil
	}
	from, err := lineUnitByName(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s is not a LineUnit", fromUnit)
	}
	to, err := lineUnitByName(toUnit)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s is not a LineUnit", toUnit)
	}
	l := LineMeasurement{
		Value: value,
		Unit:  from,
	}
	return l.To(to).Value, nil
}

// convertMassMeasurement converts a MassMeasurement from one unit To another. Params fromUnit and toUnit can be
// simple MassMeasurement units such as lb or kg, or compound units such as kg/ha or lb1ac-1.
func convertMassMeasurement(value float64, fromUnit, toUnit string) (float64, error) {
	// Nothing To do
	if fromUnit == toUnit {
		return value, nil
	}
	var err error
	// Get the MassMeasurement Unit from compound units
	if maybeCompoundUnit(fromUnit) {
		fromUnit, _, err = splitCompoundUnit(fromUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split fromUnit as compound Unit: %s", err)
		}
	}
	if maybeCompoundUnit(toUnit) {
		toUnit, _, err = splitCompoundUnit(toUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split toUnit as compound Unit: %s", err)
		}
	}
	from, err := massUnitByName(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s is not a MassUnit", fromUnit)
	}
	to, err := massUnitByName(toUnit)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s is not a MassUnit", toUnit)
	}
	m := MassMeasurement{
		Value: value,
		Unit:  from,
	}
	return m.To(to).Value, nil
}

// convertVolumeMeasurement converts a VolumeMeasurement from one unit To another. Params fromUnit and toUnit can be
// simple VolumeMeasurement units such as floz of l, or compound units such as floz/ac or l1ha-1.
func convertVolumeMeasurement(value float64, fromUnit, toUnit string) (float64, error) {
	// Nothing To do
	if fromUnit == toUnit {
		return value, nil
	}
	var err error
	// Get the VolumeMeasurement Unit from compound units
	if maybeCompoundUnit(fromUnit) {
		fromUnit, _, err = splitCompoundUnit(fromUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split fromUnit as compound Unit: %s", err)
		}
	}
	if maybeCompoundUnit(toUnit) {
		toUnit, _, err = splitCompoundUnit(toUnit)
		if err != nil {
			return 0, fmt.Errorf("could not split toUnit as compound Unit: %s", err)
		}
	}
	from, err := volumeUnitByName(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s is not a VolumeUnit", fromUnit)
	}
	to, err := volumeUnitByName(toUnit)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s is not a VolumeUnit", toUnit)
	}
	m := VolumeMeasurement{
		Value: value,
		Unit:  from,
	}
	return m.To(to).Value, nil
}

// convertMassAreaMeasurement converts the value of a MassAreaRatioMeasure (mass / area) between units.
// The params fromCompoundUnit and toCompoundUnit can be specified using 'exponent' format, eg kg1ha-1, l1ac-1,
// or 'slash' format, eg kg/ha, l/ac.
func convertMassAreaMeasurement(value float64, fromUnit, toUnit string) (float64, error) {
	units, err := checkConversionUnits(massAreaConversion, fromUnit, toUnit)
	if err != nil {
		return 0, err
	}
	fromMassUnit, err := massUnitByName(units.fromNumerator)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s does not have a mass numerator", fromUnit)
	}
	fromAreaUnit, err := areaUnitByName(units.fromDenominator)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s does not have an area denominator", fromUnit)
	}
	toMassUnit, err := massUnitByName(units.toNumerator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s does not have a mass numerator", toUnit)
	}
	toAreaUnit, err := areaUnitByName(units.toDenominator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s does not have an area denominator", toUnit)
	}

	mam := NewMassAreaRatioMeasure(value, fromMassUnit, fromAreaUnit)
	return mam.To(toMassUnit, toAreaUnit).Value(), nil
}

// convertVolumeAreaMeasurement converts the value of a VolumeAreaRatioMeasurement (volume/area) between units.
// The params fromCompoundUnit and toCompoundUnit can be specified using 'exponent' format, eg kg1ha-1, l1ac-1,
// or 'slash' format, eg kg/ha, l/ac.
func convertVolumeAreaMeasurement(value float64, fromUnit, toUnit string) (float64, error) {
	units, err := checkConversionUnits(volumeAreaConversion, fromUnit, toUnit)
	if err != nil {
		return 0, err
	}
	fromVolumeUnit, err := volumeUnitByName(units.fromNumerator)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s does not have a volume numerator", fromUnit)
	}
	fromAreaUnit, err := areaUnitByName(units.fromDenominator)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s denominator is not an AreaUnit", fromUnit)
	}
	toVolumeUnit, err := volumeUnitByName(units.toNumerator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s does not have a volume numerator", toUnit)
	}
	toAreaUnit, err := areaUnitByName(units.toDenominator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s denominator is not an AreaUnit", toUnit)
	}
	vam := NewVolumeAreaMeasurement(value, fromVolumeUnit, fromAreaUnit)
	return vam.To(toVolumeUnit, toAreaUnit).Value(), nil
}

// unitType check ensures the from and to units can be converted, and if so it returns an empty value of the
// appropriate type so the caller can do a type check. If not, it returns false.
func conversionFunc(unit1, unit2 string) func(float64, string, string) (float64, error) {
	switch {
	case IsAreaUnit(unit1) && IsAreaUnit(unit2):
		return convertAreaMeasurement
	case IsLineUnit(unit1) && IsLineUnit(unit2):
		return convertLineMeasurement
	case IsMassUnit(unit1) && IsMassUnit(unit2):
		return convertMassMeasurement
	case IsVolumeUnit(unit1) && IsVolumeUnit(unit2):
		return convertVolumeMeasurement
	case IsMassAreaRatioUnit(unit1) && IsMassAreaRatioUnit(unit2):
		return convertMassAreaMeasurement
	case IsVolumeAreaRatioUnit(unit1) && IsVolumeAreaRatioUnit(unit2):
		return convertVolumeAreaMeasurement
	}
	return nil
}

// checkConversionUnits checks that the units are valid for the conversion type
func checkConversionUnits(conversion conversionType, fromUnit, toUnit string) (*splitConversionUnits, error) {
	var units splitConversionUnits
	var err error

	switch conversion {
	case massAreaConversion, volumeAreaConversion:
		units.fromNumerator, units.fromDenominator, err = splitValueAreaCompoundUnit(fromUnit)
		if err != nil {
			return nil, fmt.Errorf("incorrect source unit for conversion %s: %s", fromUnit, err)
		}

		// numerator and denominator of the toUnit
		units.toNumerator, units.toDenominator, err = splitValueAreaCompoundUnit(toUnit)
		if err != nil {
			return nil, fmt.Errorf("incorrect target unit for conversion %s: %s", toUnit, err)
		}

		// both denominators must be an area unit
		if !IsAreaUnit(units.fromDenominator) {
			return nil, fmt.Errorf("fromUnit denominator %s is not an area unit", units.fromDenominator)
		}
		if !IsAreaUnit(units.toDenominator) {
			return nil, fmt.Errorf("toUnit denominator %s is not an area unit", units.toDenominator)
		}

		// Cannot convert between mass and volume numerators
		if IsMassUnit(units.fromNumerator) && IsVolumeUnit(units.toNumerator) {
			return nil, fmt.Errorf("cannot convert from mass (%s) to volume (%s)", fromUnit, toUnit)
		}
		if IsVolumeUnit(units.fromNumerator) && IsMassUnit(units.toNumerator) {
			return nil, fmt.Errorf("cannot convert from volume (%s) to mass (%s)", fromUnit, toUnit)
		}

		// finally, do we have the correct units for the specified conversion
		switch conversion {
		case massAreaConversion:
			if !IsMassUnit(units.fromNumerator) || !IsMassUnit(units.toNumerator) {
				return nil, fmt.Errorf("incorrect units for mass / area conversion: %s and %s", fromUnit, toUnit)
			}
		case volumeAreaConversion:
			if !IsVolumeUnit(units.fromNumerator) || !IsVolumeUnit(units.toNumerator) {
				return nil, fmt.Errorf("incorrect units for volume / area conversion: %s and %s", fromUnit, toUnit)
			}
		}
		// otherwise, we're good!
		return &units, nil
	default:
		return nil, fmt.Errorf("unhandled conversion enum: %d", conversion)
	}
}
