package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_timeUnitByName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argList  []string
		wantUnit TimeUnit
		wantErr  bool
	}{
		"second": {
			argList:  []string{"s", "sec", "secs", "second", "seconds"},
			wantUnit: Second,
			wantErr:  false,
		},
		"minute": {
			argList:  []string{"m", "min", "mins", "minute", "minutes"},
			wantUnit: Minute,
			wantErr:  false,
		},
		"hour": {
			argList:  []string{"h", "hrs", "hour", "hours"},
			wantUnit: Hour,
			wantErr:  false,
		},
		"day": {
			argList:  []string{"d", "day", "dAyS"},
			wantUnit: Day,
			wantErr:  false,
		},
		"week": {
			argList:  []string{"wks", "WEEK", "weeks"},
			wantUnit: Week,
			wantErr:  false,
		},
		"month": {
			argList:  []string{"Month", "months"},
			wantUnit: Month,
			wantErr:  false,
		},
		"year": {
			argList:  []string{"y", "year", "years"},
			wantUnit: Year,
			wantErr:  false,
		},
		"no match": {
			argList:  []string{"a", "b", "c"},
			wantUnit: TimeUnit{},
			wantErr:  true,
		},
	}

	for name, c := range cases {
		name, c := name, c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, arg := range c.argList {
				gotUnit, err := timeUnitFromString(arg)
				assert.Equal(t, c.wantErr, err != nil)
				assert.Equal(t, c.wantUnit, gotUnit)
			}
		})
	}
}
