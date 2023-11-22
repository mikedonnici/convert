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
