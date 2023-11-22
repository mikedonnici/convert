package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVolumeAreaMeasurementFromUnitString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		v       float64
		u       string
		want    VolumePerAreaMeasurement
		wantErr bool
	}{
		"1 l/ha":     {1, "l1ha-1", NewVolumeAreaMeasurement(1, Litre, Hectare), false},
		"1 l/m2":     {1, "l1m2-1", NewVolumeAreaMeasurement(1, Litre, SquareMetre), false},
		"1 ml/m2":    {1, "ml1[m2]-1", NewVolumeAreaMeasurement(1, Millilitre, SquareMetre), false},
		"1 pt/ft2":   {1, "pt1[ft2]-1", NewVolumeAreaMeasurement(1, Pint, SquareFoot), false},
		"1 floz/ft2": {1, "floz1ac-1", NewVolumeAreaMeasurement(1, FluidOunce, Acre), false},
		// not a volume Unit - error
		"error - 1 kg/m2 - not a volume unit": {1, "kg1[m2]-1", NewVolumeAreaMeasurement(0, VolumeUnit{}, AreaUnit{}), true},
	}

	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got, err := NewVolumeAreaMeasurementFromUnitString(c.v, c.u)
			assert.Equal(t, c.wantErr, err != nil)
			assert.Equal(t, c.want, got)
		})
	}
}

func Test_volumeAreaRatioUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList    []string
		wantUnit   VolumeAreaRatioUnit
		wantString string
		wantErr    bool
	}{
		"litres per square metre": {
			argList: []string{"l/m2", "l1[m2]-1"},
			wantUnit: VolumeAreaRatioUnit{
				Numerator:   Litre,
				Denominator: SquareMetre,
			},
			wantString: "l1[m2]-1",
			wantErr:    false,
		},
		"fluid ounces per square foot": {
			argList: []string{"floz/ft2", "floz1[ft2]-1"},
			wantUnit: VolumeAreaRatioUnit{
				Numerator:   FluidOunce,
				Denominator: SquareFoot,
			},
			wantString: "floz1[ft2]-1",
			wantErr:    false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: VolumeAreaRatioUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, err := volumeAreaRatioUnitByName(arg)
				assert.Equal(t, c.wantErr, err != nil)
				if !c.wantErr && err == nil {
					assert.Equal(t, c.wantUnit, gotUnit)
					assert.Equal(t, c.wantString, gotUnit.String())
				}
			}
		})
	}
}
