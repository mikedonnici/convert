package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMassAreaMeasurementFromUnitString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argValue float64
		argUnit  string
		want     MassPerAreaMeasurement
		wantErr  bool
	}{
		"1 kg/ha": {
			argValue: 1,
			argUnit:  "kg1ha-1",
			want: MassPerAreaMeasurement{
				MassMeasurement: MassMeasurement{
					Value: 1,
					Unit:  Kilogram,
				},
				unitArea: Hectare,
			},
			wantErr: false,
		},
		"1 g/m2": {
			argValue: 1,
			argUnit:  "g1[m2]-1",
			want: MassPerAreaMeasurement{
				MassMeasurement: MassMeasurement{
					Value: 1,
					Unit:  Gram,
				},
				unitArea: SquareMetre,
			},
			wantErr: false,
		},
		"invalid mass numerator [m3]1[m2]-1": {
			argValue: 1,
			argUnit:  "[m3]1[m2]-1",
			want:     MassPerAreaMeasurement{},
			wantErr:  true,
		},
		"invalid area denominator kg[m3]-1": {
			argValue: 1,
			argUnit:  "kg[m3]-1",
			want:     MassPerAreaMeasurement{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := NewMassPerAreaMeasurementFromUnitString(c.argValue, c.argUnit)
			assert.Equal(t, c.wantErr, err != nil)
			assert.Equal(t, c.want, got)
		})
	}
}

func Test_massRateTo(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg        MassPerAreaMeasurement
		toMassUnit MassUnit
		toAreaUnit AreaUnit
		want       MassPerAreaMeasurement
	}{
		"0 kg/ha to lb/ac": {
			arg:        NewMassPerAreaMeasurement(0, Kilogram, Hectare),
			toMassUnit: Pound,
			toAreaUnit: Acre,
			want:       NewMassPerAreaMeasurement(0, Pound, Acre),
		},
		"100 kg/ha to lb/ac": {
			arg:        NewMassPerAreaMeasurement(100, Kilogram, Hectare),
			toMassUnit: Pound,
			toAreaUnit: Acre,
			want:       NewMassPerAreaMeasurement(89.2179, Pound, Acre),
		},
		"100 lb/ac to kg/ha": {
			arg:        NewMassPerAreaMeasurement(89.2179, Pound, Acre),
			toMassUnit: Kilogram,
			toAreaUnit: Hectare,
			want:       NewMassPerAreaMeasurement(100, Kilogram, Hectare),
		},
	}

	const tolerance = 0.001
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := c.arg.To(c.toMassUnit, c.toAreaUnit)
			assert.InDelta(t, c.want.Value(), got.Value(), tolerance, ".To().Value")
			wantUnit, err := c.want.Unit()
			assert.NoErrorf(t, err, "c.want.Unit() err = %s", err)
			gotUnit, err := got.Unit()
			assert.NoErrorf(t, err, ".To().Unit() err = %s", err)
			assert.Equal(t, wantUnit, gotUnit, ".To().Unit = %s, want %s", gotUnit, wantUnit)
		})
	}
}
