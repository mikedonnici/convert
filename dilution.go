package convert

import (
	"fmt"
)

// DilutedProductApplication represents the application of a product (solute) that is diluted with a carrier (solvent)
// and then applied over an area. The objective is to work out how the rate that the product is applied over the area.
// For example, if 10g product X is dissolved in each litre of water (carrier) and then applied at 10L per hectare,
// how many grams of product X is applied per hectare.
type DilutedProductApplication struct {
	ProductAmount               float64 // eg 10
	ProductUnitLabel            string  // eg grams
	CarrierSolventUnitLabel     string  // eg litres
	CarrierApplicationAmount    float64 // eg 100
	CarrierApplicationUnitLabel string  // eg litres
	AreaUnitLabel               string  // eg hectares

	productUnit            Unit
	carrierSolventUnit     Unit
	carrierApplicationUnit Unit
	areaUnit               AreaUnit
}

// ApplicationRate returns the application rate of the diluted product, eg g1ha-1
func (d *DilutedProductApplication) ApplicationRate() (float64, string, error) {
	if err := d.UnitCheck(); err != nil {
		return 0, "", err
	}

	// Product in one unit of solvent
	p1 := d.ProductAmount

	// Convert the application amount to the solvent unit
	p2, err := ValueFromTo(d.CarrierApplicationAmount, d.CarrierApplicationUnitLabel, d.CarrierSolventUnitLabel)
	if err != nil {
		return 0, "", fmt.Errorf("failed to convert from %s to %s: %w", d.CarrierApplicationUnitLabel, d.CarrierSolventUnitLabel, err)
	}

	p3 := p1 * p2
	unit := RatioUnit{
		Numerator:   d.productUnit,
		Denominator: d.areaUnit,
	}
	return p3, unit.String(), nil
}

// UnitCheck checks that the units make sense.
func (d *DilutedProductApplication) UnitCheck() error {
	var err error

	d.areaUnit, err = areaUnitByName(d.AreaUnitLabel)
	if err != nil {
		return fmt.Errorf("invalid area unit: %s", d.AreaUnitLabel)
	}

	d.productUnit, err = UnitFromLabel(d.ProductUnitLabel)
	if err != nil {
		return fmt.Errorf("invalid product unit: %s", d.ProductUnitLabel)
	}
	if !IsMassUnit(d.productUnit.String()) && !IsVolumeUnit(d.productUnit.String()) {
		return fmt.Errorf("product unit %s is not a mass or volume unit", d.productUnit.String())
	}

	// CarrierSolventUnit and CarrierApplicationUnit both need to be the same type - ie, either mass or volume.
	// Eg: 10g/kg dilution spread at 50kg/ha makes sense, but 10g/l dilution spread at 10kg/ha cannot be resolved without
	// knowing density.
	if !IsMassUnit(d.CarrierSolventUnitLabel) && !IsVolumeUnit(d.CarrierSolventUnitLabel) {
		return fmt.Errorf("carrier (solvent) unit %s is not a mass or volume unit", d.CarrierSolventUnitLabel)
	}
	d.carrierSolventUnit, err = UnitFromLabel(d.CarrierSolventUnitLabel)
	if err != nil {
		return fmt.Errorf("invalid carrier (solvent) unit: %s", d.CarrierSolventUnitLabel)
	}

	if !IsMassUnit(d.CarrierApplicationUnitLabel) && !IsVolumeUnit(d.CarrierApplicationUnitLabel) {
		return fmt.Errorf("carrier (application) unit %s is not a mass or volume unit", d.CarrierApplicationUnitLabel)
	}
	d.carrierApplicationUnit, err = UnitFromLabel(d.CarrierApplicationUnitLabel)
	if err != nil {
		return fmt.Errorf("invalid carrier (application) unit: %s", d.CarrierApplicationUnitLabel)
	}

	// Final check is that the carrier (solvent) unit and the carrier (application) unit are the same type.
	if IsMassUnit(d.CarrierSolventUnitLabel) && IsVolumeUnit(d.CarrierApplicationUnitLabel) ||
		IsVolumeUnit(d.CarrierSolventUnitLabel) && IsMassUnit(d.CarrierApplicationUnitLabel) {
		return fmt.Errorf("carrier (solvent) unit %s and carrier (application) unit %s need to both be mass or both be volume", d.CarrierSolventUnitLabel, d.CarrierApplicationUnitLabel)
	}

	return nil
}
