package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areaUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList  []string
		wantUnit AreaUnit
		wantErr  bool
	}{
		"square metres": {
			argList:  []string{"m2", "m^2", "m²"},
			wantUnit: SquareMetre,
			wantErr:  false,
		},
		"hectares": {
			argList:  []string{"ha"},
			wantUnit: Hectare,
			wantErr:  false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: AreaUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, gotErr := areaUnitFromString(arg)
				assert.Equal(t, c.wantErr, gotErr != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}

func Test_areaTo(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		arg  AreaMeasurement
		want AreaMeasurement
	}{
		"0 m2 to ha": {
			arg:  AreaMeasurement{0, SquareMetre},
			want: AreaMeasurement{0, Hectare},
		},
		"10000 m2 to ha": {
			arg:  AreaMeasurement{10000, SquareMetre},
			want: AreaMeasurement{1, Hectare},
		},
		"4046.86 m2 to ac": {
			arg:  AreaMeasurement{4046.86, SquareMetre},
			want: AreaMeasurement{1, Acre},
		},
		"0.092903 m2 to ft2": {
			arg:  AreaMeasurement{0.092903, SquareMetre},
			want: AreaMeasurement{1, SquareFoot},
		},
		"1 ac to ha": {
			arg:  AreaMeasurement{1, Acre},
			want: AreaMeasurement{0.404686, Hectare},
		},
		"1 ft2 To m2": {
			arg:  AreaMeasurement{1, SquareFoot},
			want: AreaMeasurement{0.092903, SquareMetre},
		},
		"1 sq mile to km2": {
			arg:  AreaMeasurement{1, SquareMile},
			want: AreaMeasurement{2.58999, SquareKilometre},
		},
		"1034 cm2 to sq yards": {
			arg:  AreaMeasurement{1034, SquareCentimetre},
			want: AreaMeasurement{0.123457, SquareYard},
		},
	}

	const tolerance = 0.001
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			got := c.arg.To(c.want.Unit)
			assert.InDelta(t, c.want.Value, got.Value, tolerance)
			assert.Equal(t, c.want.Unit, got.Unit)
		})
	}
}
