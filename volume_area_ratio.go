package convert

import (
	"fmt"
)

// VolumePerArea is a compound unit of volume per area
type VolumePerArea string

// todo: probably remove these - unused but here for reference - too many combinations to list so we can rely on string
//
//	representations of compound units for the time being.
const (
	GallonsPerAcre        VolumePerArea = "gal1ac-1"
	FluidOuncesPerAcre    VolumePerArea = "floz1ac-1"
	QuartsPerAcre         VolumePerArea = "qt1ac-1"
	PintsPerAcre          VolumePerArea = "pt1ac-1"
	LitresPerHectare      VolumePerArea = "l1ha-1"
	KilolitresPerHectare  VolumePerArea = "kl1ha-1"
	DecilitresPerHectare  VolumePerArea = "dl1ha-1"
	MillilitresPerHectare VolumePerArea = "ml1ha-1"
	MicrolitresPerHectare VolumePerArea = "ul1ha-1"
	CubicMetresPerHectare VolumePerArea = "[m3]1ha-1"
	LitresPerAcre         VolumePerArea = "l1ac-1"

	// BushelsPerAcre is a harvest volume measurement for various grains and seeds
	BushelsPerAcre VolumePerArea = "bu1ac-1" // grain

	// BalesPerAcre is a harvest volume measurement for cotton
	BalesPerAcre VolumePerArea = "bale1ac-1" // cotton
)

// VolumeAreaRatioUnit is a divisive unit with a volume numerator and an area denominator
type VolumeAreaRatioUnit struct {
	Numerator   VolumeUnit
	Denominator AreaUnit
}

// String returns the string representation of the VolumeAreaRatioUnit unit.
func (u VolumeAreaRatioUnit) String() string {
	return RatioUnit{
		Numerator:   u.Numerator,
		Denominator: u.Denominator,
	}.String()
}

// volumeAreaRatioUnitByName attempts to derive a VolumeAreaRatioUnit from a string.
func volumeAreaRatioUnitByName(s string) (VolumeAreaRatioUnit, error) {
	n, d, err := splitCompoundUnit(s)
	if err != nil {
		return VolumeAreaRatioUnit{}, err
	}
	un, err := volumeUnitByName(n)
	if err != nil {
		return VolumeAreaRatioUnit{}, fmt.Errorf("numerator of compound unit %s (%s) is not a volume unit", s, n)
	}
	ud, err := areaUnitByName(d)
	if err != nil {
		return VolumeAreaRatioUnit{}, fmt.Errorf("denominator of compound unit %s (%s) is not an area unit", s, d)
	}
	return VolumeAreaRatioUnit{
		Numerator:   un,
		Denominator: ud,
	}, nil
}

// VolumeAreaRatioMeasurement represents a volume over AreaMeasurement Value.
type VolumeAreaRatioMeasurement struct {
	VolumeMeasurement
	AreaUnit AreaUnit
}

// NewVolumeAreaMeasurement returns a VolumeAreaRatioMeasurement with the specified field values.
func NewVolumeAreaMeasurement(v float64, vu VolumeUnit, au AreaUnit) VolumeAreaRatioMeasurement {
	return VolumeAreaRatioMeasurement{
		VolumeMeasurement: VolumeMeasurement{
			Value: v,
			Unit:  vu,
		},
		AreaUnit: au,
	}
}

// NewVolumeAreaMeasurementFromUnitString returns a VolumeAreaRatioMeasurement initialised with a Value.
// It attempts To work out the correct AreaUnit from the complexUnit string.
func NewVolumeAreaMeasurementFromUnitString(v float64, compoundUnit string) (VolumeAreaRatioMeasurement, error) {
	n, d, err := splitCompoundUnit(compoundUnit)
	if err != nil {
		return VolumeAreaRatioMeasurement{}, err
	}
	un, err := volumeUnitByName(n)
	if err != nil {
		return VolumeAreaRatioMeasurement{}, fmt.Errorf("numerator of compound unit %s (%s) is not a volume VolumeUnit", compoundUnit, n)
	}
	ud, err := areaUnitByName(d)
	if err != nil {
		return VolumeAreaRatioMeasurement{}, fmt.Errorf("denominator of compound unit %s (%s) is not an AreaUnit", compoundUnit, d)
	}
	return NewVolumeAreaMeasurement(v, un, ud), nil
}

func (vr VolumeAreaRatioMeasurement) To(toVolumeUnit VolumeUnit, toAreaUnit AreaUnit) VolumeAreaRatioMeasurement {
	toVolume := vr.VolumeMeasurement.To(toVolumeUnit)
	toArea := AreaMeasurement{Value: 1, Unit: vr.AreaUnit}.To(toAreaUnit)
	toVolume.Value = toVolume.Value / toArea.Value
	return VolumeAreaRatioMeasurement{
		VolumeMeasurement: toVolume,
		AreaUnit:          toAreaUnit,
	}
}

// Value returns the value of the VolumeAreaRatioMeasurement
func (vr VolumeAreaRatioMeasurement) Value() float64 {
	return vr.VolumeMeasurement.Value
}

// Unit returns the unit of the VolumeAreaRatioMeasurement
func (vr VolumeAreaRatioMeasurement) Unit() (string, error) {
	return joinCompoundUnit(vr.VolumeMeasurement.Unit.String(), vr.AreaUnit.String())
}

// volumeAreaUnitByName returns the first volume/area unit that is a case-sensitive match for s, or an error if no match is found.
// func volumeAreaUnitByName(s string) (VolumePerArea, error) {
// 	n, d := splitCompoundUnit(s)
//
// 	// Numerator should be a volumeUnit
// 	vu, err := volumeUnitByName(n)
// 	if err != nil {
// 		return VolumePerArea(""), fmt.Errorf("numerator of volume per area compound unit %s (%s) is not a volume unit", s, n)
// 	}
//
// 	// denominator should be an area unit
// 	au, err := areaUnitByName(d)
// 	if err != nil {
// 		return VolumePerArea(""), fmt.Errorf("denominator of volume per area compound unit %s (%s) is not an area unit", s, d)
// 	}
//
// 	for _, u := range volumeAreaUnits {
// 		if u.Matches(s) {
// 			return u, nil
// 		}
// 	}
// 	return VolumePerArea(""), fmt.Errorf("no volume/area unit found for %s", s)
// }
