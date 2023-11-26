package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDilutionRate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		dpa       DilutedProductApplication
		wantValue float64 // eg 1
		wantUnit  string  // eg l1ha-1
		wantErr   bool
	}{
		"all zeros": {
			dpa: DilutedProductApplication{
				ProductAmount:               0,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    0,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 0,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"all ones": {
			dpa: DilutedProductApplication{
				ProductAmount:               1,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    1,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 1,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"1g per l applied at 100l per ha": {
			dpa: DilutedProductApplication{
				ProductAmount:               1,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    100,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 100,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"15g per l applied at 150l per ha": {
			dpa: DilutedProductApplication{
				ProductAmount:               15,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    150,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 2250,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"15ml per l applied at 150l per ha": {
			dpa: DilutedProductApplication{
				ProductAmount:               15,
				ProductUnitLabel:            "ml",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    150,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 2250,
			wantUnit:  "ml1ha-1",
			wantErr:   false,
		},
		"250g per kg applied at 1.5t per ha": {
			dpa: DilutedProductApplication{
				ProductAmount:               250,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "kg",
				CarrierApplicationAmount:    1.5,
				CarrierApplicationUnitLabel: "t",
				AreaUnitLabel:               "ha",
			},
			wantValue: 375000,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"250g per hectolitre applied at 150l per ha": {
			dpa: DilutedProductApplication{
				ProductAmount:               250,
				ProductUnitLabel:            "g",
				CarrierSolventUnitLabel:     "100l",
				CarrierApplicationAmount:    150,
				CarrierApplicationUnitLabel: "l",
				AreaUnitLabel:               "ha",
			},
			wantValue: 375,
			wantUnit:  "g1ha-1",
			wantErr:   false,
		},
		"cannot convert solvent litres to application kg": {
			dpa: DilutedProductApplication{
				ProductAmount:               1,
				ProductUnitLabel:            "kg",
				CarrierSolventUnitLabel:     "l",
				CarrierApplicationAmount:    100,
				CarrierApplicationUnitLabel: "kg",
				AreaUnitLabel:               "ha",
			},
			wantValue: 0,
			wantUnit:  "",
			wantErr:   true,
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			v, u, err := c.dpa.ApplicationRate()
			assert.Equal(t, c.wantErr, err != nil, "ApplicationRate() err = %s", err)
			assert.Equal(t, c.wantValue, v)
			assert.Equal(t, c.wantUnit, u)
		})
	}
}
