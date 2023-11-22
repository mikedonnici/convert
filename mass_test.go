package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_massUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList  []string
		wantUnit MassUnit
		wantErr  bool
	}{
		"kilograms": {
			argList:  []string{"kg", "kilogram", "kilograms"},
			wantUnit: Kilogram,
			wantErr:  false,
		},
		"grams": {
			argList:  []string{"g", "gram", "grams"},
			wantUnit: Gram,
			wantErr:  false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: MassUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, gotErr := massUnitByName(arg)
				assert.Equal(t, c.wantErr, gotErr != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}

func TestMassMeasurementTo(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		arg  MassMeasurement
		want MassMeasurement
	}{
		"0 kg To kg":      {arg: MassMeasurement{0, Kilogram}, want: MassMeasurement{0, Kilogram}},
		"1 g To kg":       {arg: MassMeasurement{1, Gram}, want: MassMeasurement{0.001, Kilogram}},
		"1 dg To kg":      {arg: MassMeasurement{1, Decigram}, want: MassMeasurement{0.0001, Kilogram}},
		"1 mg To kg":      {arg: MassMeasurement{1, Milligram}, want: MassMeasurement{0.000001, Kilogram}},
		"1 t To kg":       {arg: MassMeasurement{1, Tonne}, want: MassMeasurement{1000, Kilogram}},
		"1 ton To kg":     {arg: MassMeasurement{1, Ton}, want: MassMeasurement{907.185, Kilogram}},
		"1 lb To kg":      {arg: MassMeasurement{1, Pound}, want: MassMeasurement{0.453592, Kilogram}},
		"1 ozm To kg":     {arg: MassMeasurement{1, OunceMass}, want: MassMeasurement{0.0283495, Kilogram}},
		"1 st To kg":      {arg: MassMeasurement{1, Stone}, want: MassMeasurement{6.35029, Kilogram}},
		"453.592 g To lb": {arg: MassMeasurement{453.592, Gram}, want: MassMeasurement{1, Pound}},
	}

	const tolerance = 0.000001
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := c.arg.To(c.want.Unit)
			assert.InDelta(t, c.want.Value, got.Value, tolerance)
			assert.Equal(t, c.want.Unit, got.Unit)
		})
	}
}
