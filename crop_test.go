package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bushelsToGrams(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		crop    Crop
		bushels float64
		want    MassMeasurement
	}{
		"alfalfa bushels to grams": {
			crop:    "alfalfa",
			bushels: 1.0,
			want:    MassMeasurement{27215, Gram},
		},

		// {name: "barley", crop: "barley", bushels: 1.0, want: MassMeasurement{21772, Gram}},
		// {name: "corn", crop: "corn", bushels: 1.0, want: MassMeasurement{25400, Gram}},
		// {name: "flax", crop: "flax", bushels: 1.0, want: MassMeasurement{25401, Gram}},
		// {name: "lucerne", crop: "lucerne", bushels: 1.0, want: MassMeasurement{27215, Gram}},
		// {name: "maize", crop: "maize", bushels: 1.0, want: MassMeasurement{25400, Gram}},
		// {name: "millet", crop: "millet", bushels: 1.0, want: MassMeasurement{22679, Gram}},
		// {name: "oats", crop: "oats", bushels: 1.0, want: MassMeasurement{14515, Gram}},
		// {name: "rye", crop: "rye", bushels: 1.0, want: MassMeasurement{25401, Gram}},
		// {name: "sorghum", crop: "sorghum", bushels: 1.0, want: MassMeasurement{25400, Gram}},
		// {name: "soybean", crop: "soybean", bushels: 1.0, want: MassMeasurement{27216, Gram}},
		// {name: "soybeans", crop: "soybeans", bushels: 1.0, want: MassMeasurement{27216, Gram}},
		// {name: "spelt", crop: "spelt", bushels: 1.0, want: MassMeasurement{18144, Gram}},
		// {name: "wheat", crop: "wheat", bushels: 1.0, want: MassMeasurement{27216, Gram}},
	}

	const tolerance = 1
	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := bushelsToGrams(c.bushels, c.crop)
			assert.NoError(t, err)
			assert.InDelta(t, c.want.Value, got.Value, tolerance)
			assert.Equal(t, c.want.Unit, got.Unit)
		})
	}
}

func Test_balesToGrams(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		crop  Crop
		bales float64
		want  MassMeasurement
	}{
		{name: "cotton", crop: "cotton", bales: 1.0, want: MassMeasurement{226800, Gram}},
	}

	const tolerance = 0.001
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got, err := balesToGrams(c.bales, c.crop)
			assert.NoError(t, err)
			assert.InDelta(t, c.want.Value, got.Value, tolerance)
			assert.Equal(t, c.want.Unit, got.Unit)
		})
	}
}
