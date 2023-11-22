package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lineUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList  []string
		wantUnit LineUnit
		wantErr  bool
	}{
		"metres": {
			argList:  []string{"m", "metre", "meter", "metres", "meters"},
			wantUnit: Metre,
			wantErr:  false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: LineUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, err := lineUnitByName(arg)
				assert.Equal(t, c.wantErr, err != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}

func TestLineMeasurementTo(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg  LineMeasurement
		want LineMeasurement
	}{
		"1 inch to centimetre": {
			arg:  LineMeasurement{1, Inch},
			want: LineMeasurement{2.54, Centimetre},
		},
		"1 metre to foot": {
			arg:  LineMeasurement{1, Metre},
			want: LineMeasurement{3.28084, Foot},
		},
		"3.28084 foot to metre": {
			arg:  LineMeasurement{3.28084, Foot},
			want: LineMeasurement{1, Metre},
		},
		"8 mile to kilometre": {
			arg:  LineMeasurement{8, Mile},
			want: LineMeasurement{12.8748, Kilometre},
		},
	}
	const tolerance = 0.0001
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := c.arg.To(c.want.Unit)
			assert.InDelta(t, got.Value, c.want.Value, tolerance)
			assert.Equal(t, got.Unit, c.want.Unit)
		})
	}
}
