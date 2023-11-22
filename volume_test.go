package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_volumeUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList  []string
		wantUnit VolumeUnit
		wantErr  bool
	}{
		"litre": {
			argList:  []string{"l", "litre", "liter", "litres", "liters"},
			wantUnit: Litre,
			wantErr:  false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: VolumeUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, err := volumeUnitByName(arg)
				assert.Equal(t, c.wantErr, err != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}

func Test_volumeTo(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg  VolumeMeasurement
		want VolumeMeasurement
	}{
		"0 litre to litre": {
			arg:  VolumeMeasurement{0, Litre},
			want: VolumeMeasurement{0, Litre},
		},
		"1 millilitre to mitre": {
			arg:  VolumeMeasurement{1, Millilitre},
			want: VolumeMeasurement{0.001, Litre},
		},
		"1 decilitre to mitre": {
			arg:  VolumeMeasurement{1, Decilitre},
			want: VolumeMeasurement{0.1, Litre},
		},
		"1 kilolitre to mitre": {
			arg:  VolumeMeasurement{1, Kilolitre},
			want: VolumeMeasurement{1000, Litre},
		},
		"1 gallon to mitre": {
			arg:  VolumeMeasurement{1, Gallon},
			want: VolumeMeasurement{3.78541, Litre},
		},
		"1 fluid ounce to litre": {
			arg:  VolumeMeasurement{1, FluidOunce},
			want: VolumeMeasurement{0.0284131, Litre},
		},
		"1 quart to litre": {
			arg:  VolumeMeasurement{1, Quart},
			want: VolumeMeasurement{0.946353, Litre},
		},
		"1 pint to litre": {
			arg:  VolumeMeasurement{1, Pint},
			want: VolumeMeasurement{0.473176, Litre},
		},
		"1 cubic metre to litre": {
			arg:  VolumeMeasurement{1, CubicMetre},
			want: VolumeMeasurement{1000, Litre},
		},
		// Although crop-specific, there are is a standard volume for a bushel: https://en.wikipedia.org/wiki/Bushel
		"1 bushel to litre": {
			arg:  VolumeMeasurement{1, Bushel},
			want: VolumeMeasurement{35.2391, Litre},
		},
		// Cotton is a bit hazy but have Us standard of .48 cubic metres: https://en.wikipedia.org/wiki/Cotton_bale
		"1 bale to litre": {
			arg:  VolumeMeasurement{1, Bale},
			want: VolumeMeasurement{480, Litre},
		},
		"1 acre inch to gallon": {
			arg:  VolumeMeasurement{1, AcreInch},
			want: VolumeMeasurement{27154.3, Gallon},
		},
		"1 acre foot to megalitre": {
			arg:  VolumeMeasurement{1, AcreFoot},
			want: VolumeMeasurement{1.23348, Megalitre},
		},
		"1 cubic yard to cubic metre": {
			arg:  VolumeMeasurement{1, CubicYard},
			want: VolumeMeasurement{0.764555, CubicMetre},
		},
		"1 cubic foot to litre": {
			arg:  VolumeMeasurement{1, CubicFoot},
			want: VolumeMeasurement{28.3168, Litre},
		},
		"1 cubic inch to litre": {
			arg:  VolumeMeasurement{1, CubicInch},
			want: VolumeMeasurement{0.0163871, Litre},
		},
	}

	const tolerance = 0.01
	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := c.arg.To(c.want.Unit)
			assert.InDelta(t, c.want.Value, got.Value, tolerance)
			assert.Equal(t, c.want.Unit, got.Unit)
		})
	}
}
