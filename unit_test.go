package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAreaUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsAreaUnit("ha"))
	assert.True(t, IsAreaUnit("hectare"))
	assert.True(t, IsAreaUnit("hectares"))
	assert.True(t, IsAreaUnit("m2"))
	assert.False(t, IsAreaUnit("m"))
	assert.False(t, IsAreaUnit("m3"))
	assert.False(t, IsAreaUnit("l"))
}

func TestIsLineUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsLineUnit("mm"))
	assert.True(t, IsLineUnit("millimetre"))
	assert.True(t, IsLineUnit("cm"))
	assert.True(t, IsLineUnit("in"))
	assert.True(t, IsLineUnit("ft"))
	assert.True(t, IsLineUnit("m"))
	assert.True(t, IsLineUnit("yd"))
	assert.True(t, IsLineUnit("km"))
	assert.True(t, IsLineUnit("mi"))
	assert.False(t, IsLineUnit("m2"))
	assert.False(t, IsLineUnit("m3"))
	assert.False(t, IsLineUnit("l"))
}

func TestIsMassUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsMassUnit("mg"))
	assert.True(t, IsMassUnit("g"))
	assert.True(t, IsMassUnit("kg"))
	assert.True(t, IsMassUnit("t"))
	assert.True(t, IsMassUnit("lb"))
	assert.True(t, IsMassUnit("oz"))
	assert.True(t, IsMassUnit("st"))
	assert.False(t, IsMassUnit("m2"))
	assert.False(t, IsMassUnit("m3"))
	assert.False(t, IsMassUnit("l"))
}

func TestIsTimeUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsTimeUnit("s"))
	assert.True(t, IsTimeUnit("min"))
	assert.True(t, IsTimeUnit("h"))
	assert.True(t, IsTimeUnit("d"))
	assert.True(t, IsTimeUnit("wk"))
	assert.True(t, IsTimeUnit("mo"))
	assert.True(t, IsTimeUnit("yr"))
	assert.False(t, IsTimeUnit("m2"))
	assert.False(t, IsTimeUnit("m3"))
	assert.False(t, IsTimeUnit("l"))
}

func TestIsMassAreaRationUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsMassAreaRatioUnit("kg/ha"))
	assert.True(t, IsMassAreaRatioUnit("kg1ha-1"))
	assert.True(t, IsMassAreaRatioUnit("t/m2"))
	assert.True(t, IsMassAreaRatioUnit("lb1ac-1"))
	assert.False(t, IsMassAreaRatioUnit("l1ha-1"))
	assert.False(t, IsMassAreaRatioUnit("kg/m3"))
	assert.False(t, IsMassAreaRatioUnit("kg/l"))
}

func TestIsVolumeAreaRationUnit(t *testing.T) {
	t.Parallel()
	assert.True(t, IsVolumeAreaRatioUnit("l/m2"))
	assert.True(t, IsVolumeAreaRatioUnit("l1[m2]-1"))
	assert.True(t, IsVolumeAreaRatioUnit("floz1[ft2]-1"))
	assert.True(t, IsVolumeAreaRatioUnit("pt/ft2"))
	assert.False(t, IsVolumeAreaRatioUnit("l/m"))
	assert.False(t, IsVolumeAreaRatioUnit("l1[m3]-1"))
	assert.False(t, IsVolumeAreaRatioUnit("[m2]1/l"))
}

func TestStandardUnit(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg  string
		want any
	}{
		"m2": {
			arg:  "m2",
			want: SquareMetre,
		},
		"m²": {
			arg:  "m²",
			want: SquareMetre,
		},
		"m^2": {
			arg:  "m^2",
			want: SquareMetre,
		},
		"square metre": {
			arg:  "square metre",
			want: SquareMetre,
		},
		"square metres": {
			arg:  "square metres",
			want: SquareMetre,
		},
		"square meter": {
			arg:  "square meter",
			want: SquareMetre,
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := UnitFromLabel(c.arg)
			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

// Tests the conversion of ratio units to their string representation.
func TestRatioUnit_String(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg  Unit
		want string
	}{
		"litres per hectare": {
			arg: VolumeAreaRatioUnit{
				Numerator:   Litre,
				Denominator: Hectare,
			},
			want: "l1ha-1",
		},
		"kilograms per hecatre": {
			arg: MassAreaRatioUnit{
				Numerator:   Kilogram,
				Denominator: Hectare,
			},
			want: "kg1ha-1",
		},
		"litres per square metre": {
			arg: VolumeAreaRatioUnit{
				Numerator:   Litre,
				Denominator: SquareMetre,
			},
			want: "l1[m2]-1",
		},
		"cubic metres per square metre": {
			arg: VolumeAreaRatioUnit{
				Numerator:   CubicMetre,
				Denominator: SquareMetre,
			},
			want: "[m3]1[m2]-1",
		},
		"tons per square mile": {
			arg: MassAreaRatioUnit{
				Numerator:   Ton,
				Denominator: SquareMile,
			},
			want: "ton1[mi2]-1",
		},
		"tonnes per square kilometre": {
			arg: MassAreaRatioUnit{
				Numerator:   Tonne,
				Denominator: SquareKilometre,
			},
			want: "t1[km2]-1",
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := c.arg.String()
			assert.Equal(t, c.want, got)
		})
	}
}

func TestStandardLabel(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList []string
		want    string
		wantErr bool
	}{
		"square metre": {
			argList: []string{"m2", "m²", "m^2", "square metre", "square metres", "square meter"},
			want:    "m2",
		},
		"hectare": {
			argList: []string{"ha", "hectare", "hectares"},
			want:    "ha",
		},
		"millimetre": {
			argList: []string{"mm", "millimetre", "millimetres"},
			want:    "mm",
		},
		"kilograms per hectare": {
			argList: []string{"kg/ha", "kg1ha-1", "kilograms / hectare", "kilograms per hectare"},
			want:    "kg1ha-1",
		},
		"litres per square metre": {
			argList: []string{"l/m2", "l1[m2]-1", "litres per square metre"},
			want:    "l1[m2]-1",
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				got, err := StandardLabel(arg)
				assert.Equal(t, c.wantErr, err != nil)
				assert.Equal(t, c.want, got)
			}
		})
	}
}
