package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeometry_AsWKT(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		coords  [][]float64
		xy      bool
		wantWKT string
		wantErr bool
	}{
		{
			name:    "empty",
			coords:  [][]float64{},
			xy:      false,
			wantWKT: "",
			wantErr: true,
		},
		{
			name: "simple polygon",
			coords: [][]float64{
				{0, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
			xy:      false,
			wantWKT: "POLYGON ((0 0, 0 1, 1 1, 1 0))",
			wantErr: false,
		},
		{
			name: "xy polygon",
			coords: [][]float64{
				{-26.432, 151.080},
				{-26.431, 151.068},
				{-26.436, 151.066},
				{-26.436, 151.069},
				{-26.439, 151.068},
				{-26.442, 151.076},
				{-26.432, 151.080},
			},
			xy:      false,
			wantWKT: "POLYGON ((151.08 -26.432, 151.068 -26.431, 151.066 -26.436, 151.069 -26.436, 151.068 -26.439, 151.076 -26.442, 151.08 -26.432))",
			wantErr: false,
		},
		{
			name: "yx polygon",
			coords: [][]float64{
				{151.080, -26.432},
				{151.068, -26.431},
				{151.066, -26.436},
				{151.069, -26.436},
				{151.068, -26.439},
				{151.076, -26.442},
				{151.080, -26.432},
				{151.080, -26.432},
			},
			xy:      true,
			wantWKT: "POLYGON ((151.08 -26.432, 151.068 -26.431, 151.066 -26.436, 151.069 -26.436, 151.068 -26.439, 151.076 -26.442, 151.08 -26.432, 151.08 -26.432))",
			wantErr: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotWKT, err := NewGeometry(Polygon).WithCoords(c.coords, c.xy).ToWKT()
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.wantWKT, gotWKT)
			}
		})
	}
}
