package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/regrowag/ses/go/pkg/convert"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		argValue    float64
		argFromUnit string
		argToUnit   string
		wantValue   float64
		wantError   bool
	}{
		"zero value": {
			argValue:    0,
			argFromUnit: "kg",
			argToUnit:   "g",
			wantValue:   0,
			wantError:   false,
		},
		"kg to kg": {
			argValue:    1,
			argFromUnit: "kg",
			argToUnit:   "kg",
			wantValue:   1,
		},

		"cm to m": {
			argValue:    100,
			argFromUnit: "cm",
			argToUnit:   "m",
			wantValue:   1,
			wantError:   false,
		},
		"kg to g": {
			argValue:    1,
			argFromUnit: "kg",
			argToUnit:   "g",
			wantValue:   1000,
			wantError:   false,
		},
		"m2 to cm2": {
			argValue:    1,
			argFromUnit: "m2",
			argToUnit:   "cm2",
			wantValue:   10000,
			wantError:   false,
		},
		"m2 to ha": {
			argValue:    10000,
			argFromUnit: "m2",
			argToUnit:   "ha",
			wantValue:   1,
			wantError:   false,
		},
		"m2 to ac": {
			argValue:    4046.86,
			argFromUnit: "m2",
			argToUnit:   "ac",
			wantValue:   1,
			wantError:   false,
		},
		"m2 to ft2": {
			argValue:    0.092903,
			argFromUnit: "m2",
			argToUnit:   "ft2",
			wantValue:   1,
			wantError:   false,
		},
		"ac to ha": {
			argValue:    1,
			argFromUnit: "ac",
			argToUnit:   "ha",
			wantValue:   0.404686,
			wantError:   false,
		},
		"kg1ha-1 to kg1ac-1": {
			argValue:    1,
			argFromUnit: "kg1ha-1",
			argToUnit:   "kg1ac-1",
			wantValue:   0.404686,
			wantError:   false,
		},
		"kg1ha-1 to lb1ac-1": {
			argValue:    1,
			argFromUnit: "kg1ha-1",
			argToUnit:   "lb1ac-1",
			wantValue:   0.892180,
			wantError:   false,
		},
		"t1ha-1 to ton1ac-1": {
			argValue:    1,
			argFromUnit: "t1ha-1",
			argToUnit:   "ton1ac-1",
			wantValue:   0.44609,
			wantError:   false,
		},
		"gal1ac-1 to l1ha-1": {
			argValue:    1,
			argFromUnit: "gal1ac-1",
			argToUnit:   "l1ha-1",
			wantValue:   9.35396,
			wantError:   false,
		},
	}

	const tolerance = 0.0001
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotValue, gotError := convert.ValueFromTo(c.argValue, c.argFromUnit, c.argToUnit)
			assert.InDelta(t, c.wantValue, gotValue, tolerance)
			assert.Equal(t, c.wantError, gotError != nil)
			t.Logf("gotValue: %v, gotError: %v", gotValue, gotError)
		})
	}
}

// func TestRate(t *testing.T) {
// 	t.Parallel()
//
// 	cases := []struct {
// 		argValue    float64
// 		argFromUnit string
// 		argToUnit   string
// 		wantValue   float64
// 		wantError   bool
// 	}{
// 		{0, "kg1ha-1", "lb1ac-1", 0, false},
// 		{1, "kg1ha-1", "kg1ha-1", 1, false},
// 		{1, "kg1ha-1", "lb1ac-1", 0.892179, false},
// 		{1, "ton1ac-1", "kg1ha-1", 2241.7010, false},
// 		{0.892179, "lb1ac-1", "kg1ha-1", 1, false},
// 		{100, "gal1ac-1", "l1ha-1", 935.396, false},
//
// 		// Missing, incorrect, mismatched units - error
// 		{1, "", "floz1ac-1", 0, true},
// 		{1, "kg1ha-1", "", 0, true},
// 		{1, "ha1kg-1", "floz1ac-1", 0, true},
// 		{1, "kg1[m3]-1", "floz1ac-1", 0, true}, // non-AreaMeasurement denominator
// 		{1, "kg1ha-1", "floz1lb-1", 0, true},   // non-AreaMeasurement denominator
// 		{1, "kg1ha-1", "floz1ac-1", 0, true},   // MassMeasurement To volume
// 	}
//
// 	const tolerance = 0.5
// 	for i, c := range cases {
// 		c := c
// 		t.Run(fmt.Sprint(i), func(t *testing.T) {
// 			t.Parallel()
// 			got, err := unit.ConvertRateMeasurement(c.argValue, c.argFromUnit, c.argToUnit)
// 			assert.Equal(t, c.wantError, err != nil)
// 			assert.InDelta(t, c.wantValue, got, tolerance)
// 		})
// 	}
// }
//
// func TestCropRate(t *testing.T) {
// 	t.Parallel()
//
// 	cases := []struct {
// 		name        string
// 		crop        string
// 		argValue    float64
// 		argFromUnit string
// 		argToUnit   string
// 		wantValue   float64
// 		wantError   bool
// 	}{
// 		// vol -> MassMeasurement conversions
// 		{"alfalfa-bu-lb", "Alfalfa", 1, "bu1ac-1", "lb1ac-1", 60.0, false},
// 		{"barley-bu-kg", "Barley", 1, "bu1ac-1", "kg1ac-1", 21.7724, false},
// 		{"flax-bu-lb", "Flax", 1, "bu1ac-1", "lb1ac-1", 56.00, false},
// 		{"millet-bu-kg", "Millet", 26.8, "bu1ac-1", "lb1ac-1", 1340, false},
// 		{"rye-bu-lb", "Rye", 1, "bu1ac-1", "lb1ac-1", 56.0, false},
// 		{"soybeans-bu-t", "Soybeans", 36.744, "bu1ac-1", "t1ac-1", 1, false},
// 		{"spelt-bu-kg", "Spelt", 26.5, "bu1ac-1", "kg1ac-1", 480.808, false},
//
// 		// Mass -> volume conversions
// 		{"corn-t-bu", "Corn", 1, "t1ac-1", "bu1ac-1", 39.3680, false},
// 		{"lucerne-ton-bu", "Alfalfa", 2.999, "ton1ac-1", "bu1ac-1", 100.0, false},
// 		{"maize-t-bu", "Maize", 1, "t1ac-1", "bu1ac-1", 39.3680, false},
// 		{"oats-t-bu", "Oats", 1, "ton1ac-1", "bu1ac-1", 62.5, false},
// 		{"sorghum-lb-bu", "Sorghum", 646, "lb1ac-1", "bu1ac-1", 11.536, false},
// 		{"soybean-t-bu", "Soybean", 1, "t1ac-1", "bu1ac-1", 36.744, false},
// 		{"wheat-t-bu", "Wheat", 1, "t1ac-1", "bu1ac-1", 36.7437, false},
// 		{"wheat-ton-bu", "Wheat", 1, "ton1ac-1", "bu1ac-1", 33.3334, false},
//
// 		// 10 bales of cotton
// 		{"cotton-bale-t-ac", "Cotton", 10, "bale1ac-1", "t1ac-1", 2.268, false},
// 		{"cotton-t-bale-ac", "Cotton", 2.268, "t1ac-1", "bale1ac-1", 10, false},
// 		// Change of AreaMeasurement Unit (todo: add a few more of these)
// 		{"cotton-bale-t-ha", "Cotton", 10, "bale1ac-1", "t1ha-1", 5.604, false},
//
// 		// errors
// 		{"no crop", "", 1, "bu1ha-1", "kg1ha-1", 0, true},
// 	}
//
// 	const tolerance = 1
// 	for _, c := range cases {
// 		c := c
// 		t.Run(c.name, func(t *testing.T) {
// 			t.Parallel()
// 			got, err := unit.CropRate(c.crop, c.argValue, c.argFromUnit, c.argToUnit)
// 			assert.Equal(t, c.wantError, err != nil)
// 			assert.InDelta(t, c.wantValue, got, tolerance)
// 		})
// 	}
// }
//
// func TestConvertArea(t *testing.T) {
// 	t.Parallel()
//
// 	cases := []struct {
// 		argValue    float64
// 		argFromUnit string
// 		argToUnit   string
// 		wantValue   float64
// 		wantError   bool
// 	}{
// 		// Can parse simple or compound units To get an AreaMeasurement conversion
// 		{1, "ha", "ha", 1, false},
// 		{1, "ha", "ac", 2.47105, false},
// 		{1, "kg1ha-1", "ac", 2.47105, false},
// 		{1, "ha", "gal1ac-1", 2.47105, false},
// 		{1, "kg1ha-1", "gal1ac-1", 2.47105, false},
// 		{1, "lb1ac-1", "ha", 0.404686, false},
// 	}
//
// 	const tolerance = 0.001
// 	for i, c := range cases {
// 		c := c
// 		t.Run(fmt.Sprint(i), func(t *testing.T) {
// 			t.Parallel()
// 			got, err := unit.convertAreaMeasurement(c.argValue, c.argFromUnit, c.argToUnit)
// 			assert.Equal(t, c.wantError, err != nil)
// 			assert.InDelta(t, c.wantValue, got, tolerance)
// 		})
// 	}
// }
//
// func Test_checkConversionUnits(t *testing.T) {
// 	t.Parallel()
//
// 	cases := []struct {
// 		name      string
// 		argConversion conversionType
// 		argFrom   string
// 		argTo     string
// 		wantError bool
// 	}{
// 		{"same unit", "ha", "ha", false},
// 	}
// }
