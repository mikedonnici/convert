package convert

import (
	"fmt"
	"testing"
)

func TestIsMassRateUnit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		arg  string
		want bool
	}{
		{"kg1ha-1", true},
		{"kg 1 ha -1", true},
		{"lb1ac-1", true},
		{"pt1[ft2]-1", false},
		{"floz1[ft2]-1", false},
		{"ozm1ac-1", true},
		{"kg/ha", true},
		{"lb/ac", true},
		{"pt/ft2", false},
		{"AAA/BBB", false},
	}
	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := IsMassAreaRatioUnit(c.arg)
			if got != c.want {
				t.Errorf("IsMassRateUnit() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestIsVolumeRateUnit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		arg  string
		want bool
	}{
		{"kg1ha-1", false},
		{"kg 1 ha -1", false},
		{"lb1ac-1", false},
		{"pt1[ft2]-1", true},
		{"floz1[ft2]-1", true},
		{"ozm1ac-1", false},
		{"kg/ha", false},
		{"lb/ac", false},
		{"pt/ft2", true},
		{"AAA/BBB", false},
	}
	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := IsVolumeAreaRatioUnit(c.arg)
			if got != c.want {
				t.Errorf("IsMassRateUnit() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_splitCompoundUnit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		arg             string
		wantNumerator   string
		wantDenominator string
		wantError       bool
	}{
		// Exponent format - spaces should be ok
		{"kg1ha-1", "kg", "ha", false},
		{"kg 1 ha -1", "kg", "ha", false},
		{"lb1ac-1", "lb", "ac", false},
		{"pt1[ft2]-1", "pt", "ft2", false},
		{"floz1[ft2]-1", "floz", "ft2", false},
		{"ozm1ac-1", "ozm", "ac", false},

		// Slash format
		{"kg/ha", "kg", "ha", false},
		{"kg  /  ha", "kg", "ha", false},
		{"lb/ac", "lb", "ac", false},
		{"pt/ft2", "pt", "ft2", false},
		{"floz/ft2", "floz", "ft2", false},
		{"ozm/ac", "ozm", "ac", false},

		// Case-insensitive
		{"Kg/Ha", "kg", "ha", false},
		{"kG  /  hA", "kg", "ha", false},
		{"LB/ac", "lb", "ac", false},
		{"pt/FT2", "pt", "ft2", false},
		{"FlOz/fT2", "floz", "ft2", false},
		{"OZM/AC", "ozm", "ac", false},

		// Missing square brackets - this works but maybe should be an error?
		{"m3 1 ha-1", "m3", "ha", false},
		{"lb1ft2-1", "lb", "ft2", false},

		// bung or unknown units - error
		{"", "", "", true},
		{"xx1yy-1", "", "", true},
		{"[kg2]1ha-1", "", "", true},
		{"kg1[ha3]-1", "", "", true},
		{"aa/bb", "", "", true},
		{"aa//bb", "", "", true},
		{"lb//ac", "", "", true},
	}

	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			gotNumerator, gotDenominator, err := splitCompoundUnit(c.arg)
			if c.wantError && err == nil {
				t.Fatalf("splitCompoundUnit() err = nil, want error")
			}
			if err != nil && !c.wantError {
				t.Fatalf("splitCompoundUnit() err = %s", err)
			}
			if gotNumerator != c.wantNumerator {
				t.Fatalf("splitCompoundUnit() numerator = %s, want %s", gotNumerator, c.wantNumerator)
			}
			if gotDenominator != c.wantDenominator {
				t.Fatalf("splitCompoundUnit() denominator = %s, want %s", gotDenominator, c.wantDenominator)
			}
		})
	}
}

func Test_joinCompoundUnit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		numerator   string
		denominator string
		want        string
		wantError   bool
	}{
		{"kg", "ha", "kg1ha-1", false},
		{"lb", "ac", "lb1ac-1", false},
		{"pt", "ft2", "pt1[ft2]-1", false},
		{"floz", "ft2", "floz1[ft2]-1", false},
		{"ozm", "ac", "ozm1ac-1", false},

		// bung or unknown units - error
		{"", "", "", true},
		{"xx", "yy", "", true},
		{"kg2", "ha", "", true},
		{"kg", "ha3", "", true},

		// Valid but square brackets have nto been removed - error
		{"[m3]", "ha", "", true},
		{"lb", "[ft2]", "", true},
	}

	for i, c := range cases {
		c := c
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got, err := joinCompoundUnit(c.numerator, c.denominator)
			if c.wantError && err == nil {
				t.Fatal("joinCompoundUnit() err = nil, want error")
			}
			if err != nil && !c.wantError {
				t.Fatalf("joinCompoundUnit() err = %s", err)
			}
			if got != c.want {
				t.Fatalf("joinCompoundUnit() = %s, want %s", got, c.want)
			}
		})
	}
}
