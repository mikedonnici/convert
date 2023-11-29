package convert

import (
	"fmt"
	"strings"
)

type Time string

const (
	SecondStandard Time = "s"
	MinuteStandard Time = "min"
	HourStandard   Time = "h"
	DayStandard    Time = "d"
	WeekStandard   Time = "wk"
	MonthStandard  Time = "mo"
	YearStandard   Time = "yr"
)

// String returns the string representation of the time unit
func (t Time) String() string {
	return string(t)
}

// TimeUnit represents a time unit.
type TimeUnit struct {
	unit       Time
	full       string
	fancy      string
	aliases    []string
	conversion float64
}

// String returns the string representation of the base time unit.
func (u TimeUnit) String() string {
	return u.unit.String()
}

// Matches returns true if s matches the time unit.
func (u TimeUnit) Matches(s string) bool {
	if strings.EqualFold(u.String(), s) ||
		strings.EqualFold(u.fancy, s) ||
		strings.EqualFold(u.full, s) {
		return true
	}
	for _, alias := range u.aliases {
		if strings.EqualFold(alias, s) {
			return true
		}
	}
	return false
}

// timeUnits is a list of all supported time units.
var timeUnits = []TimeUnit{
	Second,
	Minute,
	Hour,
	Day,
	Week,
	Month,
	Year,
}

var Second = TimeUnit{
	unit:  SecondStandard,
	full:  "second",
	fancy: string(SecondStandard),
	aliases: []string{
		"seconds",
		"sec",
		"secs",
	},
	conversion: 1,
}

var Minute = TimeUnit{
	unit:  MinuteStandard,
	full:  "minute",
	fancy: string(MinuteStandard),
	aliases: []string{
		"minutes",
		"mins",
		"m",
	},
	conversion: 60,
}

var Hour = TimeUnit{
	unit:  HourStandard,
	full:  "hour",
	fancy: string(HourStandard),
	aliases: []string{
		"hours",
		"hr",
		"hrs",
	},
	conversion: 3600,
}

var Day = TimeUnit{
	unit:  DayStandard,
	full:  "day",
	fancy: string(DayStandard),
	aliases: []string{
		"days",
	},
	conversion: 86400,
}

var Week = TimeUnit{
	unit:  WeekStandard,
	full:  "week",
	fancy: string(WeekStandard),
	aliases: []string{
		"weeks",
		"wks",
	},
	conversion: 604800,
}

var Month = TimeUnit{
	unit:  MonthStandard,
	full:  "month",
	fancy: string(MonthStandard),
	aliases: []string{
		"months",
	},
	conversion: 2628000,
}

var Year = TimeUnit{
	unit:  YearStandard,
	full:  "year",
	fancy: string(YearStandard),
	aliases: []string{
		"years",
		"y",
		"yrs",
	},
	conversion: 31536000,
}

// timeUnitFromString returns the first time unit that matches s.
func timeUnitFromString(s string) (TimeUnit, error) {
	for _, u := range timeUnits {
		if u.Matches(s) {
			return u, nil
		}
	}
	return TimeUnit{}, fmt.Errorf("no time unit found for %s", s)
}

// TimeMeasurement represents a time measurement.
type TimeMeasurement struct {
	Value float64
	Unit  TimeUnit
}

// To converts a time measurement To the specified unit.
func (m TimeMeasurement) To(unit TimeUnit) TimeMeasurement {
	if m.Value != 0 {
		m.Value = (m.Value * m.Unit.conversion) / unit.conversion
	}
	m.Unit = unit
	return m
}
