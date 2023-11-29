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
		want     MassAreaRatioMeasure
		wantErr  bool
	}{
		"1 kg/ha": {
			argValue: 1,
			argUnit:  "kg1ha-1",
			want: MassAreaRatioMeasure{
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
			want: MassAreaRatioMeasure{
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
			want:     MassAreaRatioMeasure{},
			wantErr:  true,
		},
		"invalid area denominator kg[m3]-1": {
			argValue: 1,
			argUnit:  "kg[m3]-1",
			want:     MassAreaRatioMeasure{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := NewMassAreaRatioMeasureFromUnitString(c.argValue, c.argUnit)
			assert.Equal(t, c.wantErr, err != nil)
			assert.Equal(t, c.want, got)
		})
	}
}

func Test_massRateTo(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg        MassAreaRatioMeasure
		toMassUnit MassUnit
		toAreaUnit AreaUnit
		want       MassAreaRatioMeasure
	}{
		"0 kg/ha to lb/ac": {
			arg:        NewMassAreaRatioMeasure(0, Kilogram, Hectare),
			toMassUnit: Pound,
			toAreaUnit: Acre,
			want:       NewMassAreaRatioMeasure(0, Pound, Acre),
		},
		"100 kg/ha to lb/ac": {
			arg:        NewMassAreaRatioMeasure(100, Kilogram, Hectare),
			toMassUnit: Pound,
			toAreaUnit: Acre,
			want:       NewMassAreaRatioMeasure(89.2179, Pound, Acre),
		},
		"100 lb/ac to kg/ha": {
			arg:        NewMassAreaRatioMeasure(89.2179, Pound, Acre),
			toMassUnit: Kilogram,
			toAreaUnit: Hectare,
			want:       NewMassAreaRatioMeasure(100, Kilogram, Hectare),
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

func Test_massRateUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList    []string
		wantUnit   MassAreaRatioUnit
		wantString string
		wantErr    bool
	}{
		"kilograms per hectare": {
			argList: []string{"kg/ha", "kg1ha-1"},
			wantUnit: MassAreaRatioUnit{
				Numerator:   Kilogram,
				Denominator: Hectare,
			},
			wantString: "kg1ha-1",
			wantErr:    false,
		},
		"pounds per acre": {
			argList: []string{"lb/ac", "lb1ac-1"},
			wantUnit: MassAreaRatioUnit{
				Numerator:   Pound,
				Denominator: Acre,
			},
			wantString: "lb1ac-1",
			wantErr:    false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: MassAreaRatioUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, err := massAreaRatioUnitFromString(arg)
				assert.Equal(t, c.wantErr, err != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}
