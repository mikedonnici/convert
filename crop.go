package convert

import (
	"errors"
	"fmt"
	"strings"
)

// Crop represents crops for which measurements can be converted between
// MassMeasurement and volume rates, via 'bushels' for grains / seeds and bales for cotton.
type Crop string

const (
	Alfalfa  Crop = "alfalfa"
	Barley   Crop = "barley"
	Corn     Crop = "corn"
	Cotton   Crop = "cotton"
	Flax     Crop = "flax"
	Lucerne  Crop = "lucerne"
	Maize    Crop = "maize"
	Millet   Crop = "millet"
	Oats     Crop = "oats"
	Rye      Crop = "rye"
	Sorghum  Crop = "sorghum"
	Soybean  Crop = "soybean"
	Soybeans Crop = "soybeans"
	Spelt    Crop = "spelt"
	Wheat    Crop = "wheat"
)

// bushelCrops are crops for which harvest can be recorded in bushels.
var bushelCrops = []Crop{
	Alfalfa,
	Barley,
	Corn,
	Flax,
	Lucerne,
	Maize,
	Millet,
	Oats,
	Rye,
	Sorghum,
	Soybean,
	Soybeans,
	Spelt,
	Wheat,
}

// baleCrops are crops for which harvest can be recorded in bales.
var baleCrops = []Crop{
	Cotton,
}

// cropBushelsToGrams provides a factor for converting from 1 Bushel of the specified crop, To grams.
var cropBushelsToGrams = map[Crop]float64{
	Alfalfa:  27215.5,
	Barley:   21772,
	Corn:     25400,
	Flax:     25401.2,
	Lucerne:  27215.5,
	Maize:    25400,
	Millet:   22679.6,
	Oats:     14515, // US (32lb), Canada is 15.4221 (34lb)
	Rye:      25401.2,
	Sorghum:  25400,
	Soybean:  27215.5,
	Soybeans: 27215.5,
	Spelt:    18143.7,
	Wheat:    27215.5,
}

// cropBalesToGrams provides a factor for converting from 1 Bale of the specified crop, To grams.
// Only cotton for now but may also be applicable To hay and similar.
var cropBalesToGrams = map[Crop]float64{
	Cotton: 226800, // ref: https://en.wikipedia.org/wiki/Cotton_bale (500lb)
}

func isBushelCrop(s string) bool {
	for _, u := range bushelCrops {
		if strings.EqualFold(string(u), s) {
			return true
		}
	}
	return false
}

func isBaleCrop(s string) bool {
	for _, u := range baleCrops {
		if strings.EqualFold(string(u), s) {
			return true
		}
	}
	return false
}

// bushelsToGrams converts the given number of crop bushels To grams.
func bushelsToGrams(bushels float64, crop Crop) (MassMeasurement, error) {
	f, ok := cropBushelsToGrams[crop]
	if !ok {
		return MassMeasurement{}, fmt.Errorf("cannot convert bushels To grams for %s", crop)
	}
	return MassMeasurement{
		Value: bushels * f,
		Unit:  Gram,
	}, nil
}

// cropGramsToBushels coverts grams To crop bushels.
func cropGramsToBushels(crop Crop, grams float64) (VolumeMeasurement, error) {
	f, ok := cropBushelsToGrams[crop]
	if !ok {
		return VolumeMeasurement{}, fmt.Errorf("cannot convert grams To bushels for %s", crop)
	}
	return VolumeMeasurement{
		Value: grams * (1 / f),
		Unit:  Bushel,
	}, nil
}

// cropBushelsToMass converts the given number of crop bushels To the specified MassUnit.
func cropBushelsToMass(crop Crop, bushels float64, unit MassUnit) (MassMeasurement, error) {
	m, err := bushelsToGrams(bushels, crop)
	if err != nil {
		return MassMeasurement{}, err
	}
	return m.To(unit), nil
}

// cropMassToBushels converts the given crop MassMeasurement To bushels.
func cropMassToBushels(crop Crop, cropMass MassMeasurement) (VolumeMeasurement, error) {
	m := cropMass.To(Gram)
	return cropGramsToBushels(crop, m.Value)
}

// balesToGrams converts the given number of crop bales To grams.
func balesToGrams(bales float64, crop Crop) (MassMeasurement, error) {
	f, ok := cropBalesToGrams[crop]
	if !ok {
		return MassMeasurement{}, fmt.Errorf("cannot convert bales To grams for %s", crop)
	}
	return MassMeasurement{
		Value: bales * f,
		Unit:  Gram,
	}, nil
}

// cropGramsToBales coverts grams To crop bales.
func cropGramsToBales(crop Crop, grams float64) (VolumeMeasurement, error) {
	f, ok := cropBalesToGrams[crop]
	if !ok {
		return VolumeMeasurement{}, fmt.Errorf("cannot convert grams To bales for %s", crop)
	}
	return VolumeMeasurement{
		Value: grams * (1 / f),
		Unit:  Bale,
	}, nil
}

// cropBalesToMass converts the given number of crop bales To the specified MassUnit.
func cropBalesToMass(crop Crop, bales float64, unit MassUnit) (MassMeasurement, error) {
	m, err := balesToGrams(bales, crop)
	if err != nil {
		return MassMeasurement{}, err
	}
	return m.To(unit), nil
}

// cropMassToBales converts the given crop MassMeasurement To bales.
func cropMassToBales(crop Crop, cropMass MassMeasurement) (VolumeMeasurement, error) {
	m := cropMass.To(Gram)
	return cropGramsToBales(crop, m.Value)
}

// convertCropRate handles conversion between MassMeasurement and volume for crops whose yield
// can be measured in either bushels or bales.
func convertCropRate(crop string, value float64, fromUnit, toUnit string) (float64, error) {
	if crop == "" {
		return 0, errors.New("crop cannot be nil")
	}
	if !isBushelCrop(crop) && !isBaleCrop(crop) {
		return 0, fmt.Errorf("crop must be one of the bushel crops %v, or a bale crop %v", bushelCrops, baleCrops)
	}

	// Get the units
	_, fromDenominator, err := splitCompoundUnit(fromUnit)
	if err != nil {
		return 0, fmt.Errorf("could not determine numerator in fromCompoundUnit: %w", err)
	}
	toNumerator, toDenominator, err := splitCompoundUnit(toUnit)
	if err != nil {
		return 0, fmt.Errorf("could not determine numerator in toCompoundUnit: %w", err)
	}

	fromAreaUnit, err := areaUnitByName(fromDenominator)
	if err != nil {
		return 0, fmt.Errorf("fromUnit %s denominator is not an AreaUnit: %w", fromUnit, err)
	}
	toAreaUnit, err := areaUnitByName(toDenominator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s denominator is not an AreaUnit: %w", toUnit, err)
	}

	// Mass -> Volume
	if IsMassAreaRatioUnit(fromUnit) {
		massRate, err := NewMassAreaRatioMeasureFromUnitString(value, fromUnit)
		if err != nil {
			return 0, fmt.Errorf("could not create MassAreaMeasurement Value: %w", err)
		}
		var cropVol VolumeMeasurement
		if isBushelCrop(crop) {
			cropVol, err = cropMassToBushels(Crop(crop), massRate.MassMeasurement)
			if err != nil {
				return 0, fmt.Errorf("could not convert crop MassMeasurement To bushels: %w", err)
			}
		}
		if isBaleCrop(crop) {
			cropVol, err = cropMassToBales(Crop(crop), massRate.MassMeasurement)
			if err != nil {
				return 0, fmt.Errorf("could not convert crop MassMeasurement To bales: %w", err)
			}
		}

		toVolumeUnit, err := volumeUnitByName(toNumerator)
		if err != nil {
			return 0, fmt.Errorf("toUnit %s numerator is not a VolumeUnit: %w", toUnit, err)
		}
		volRate := NewVolumeAreaMeasurement(cropVol.Value, cropVol.Unit, fromAreaUnit)
		toRate := volRate.To(toVolumeUnit, toAreaUnit)
		return toRate.Value(), nil
	}

	// Volume -> Mass
	vr, err := NewVolumeAreaMeasurementFromUnitString(value, fromUnit)
	if err != nil {
		return 0, fmt.Errorf("could not create VolumeAreaMeasurement Value: %w", err)
	}
	toMassUnit, err := massUnitByName(toNumerator)
	if err != nil {
		return 0, fmt.Errorf("toUnit %s numerator is not a MassUnit: %w", toUnit, err)
	}

	var cropMass MassMeasurement
	if isBushelCrop(crop) {
		vRate := vr.To(Bushel, fromAreaUnit) // keep original AreaUnit
		cropMass, err = cropBushelsToMass(Crop(crop), vRate.Value(), toMassUnit)
		if err != nil {
			return 0, fmt.Errorf("could not convert crop bushels To MassMeasurement: %w", err)
		}
	}
	if isBaleCrop(crop) {
		vRate := vr.To(Bale, fromAreaUnit) // keep original AreaUnit
		cropMass, err = cropBalesToMass(Crop(crop), vRate.Value(), toMassUnit)
		if err != nil {
			return 0, fmt.Errorf("could not convert crop bales To MassMeasurement: %w", err)
		}
	}
	massRate := NewMassAreaRatioMeasure(cropMass.Value, cropMass.Unit, fromAreaUnit)
	toRate := massRate.To(toMassUnit, toAreaUnit)
	return toRate.Value(), nil
}
