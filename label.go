package convert

//
// import "fmt"
//
// var labelsWithSuperscript = map[any]string{
// 	SquareCentimetre: "cm²",
// 	CubicCentimetre:  "cm³",
// 	SquareMetre:      "m²",
// 	CubicMetre:       "m³",
// 	SquareKilometre:  "km²",
// 	SquareInch:       "in²",
// 	CubicInch:        "in³",
// 	SquareFoot:       "ft²",
// 	CubicFoot:        "ft³",
// 	CubicYard:        "yd³",
// 	SquareMile:       "mi²",
// }
//
// var labelsFullWord = map[string]string{
// 	"mm":       "millimetre",
// 	"cm":       "centimetre",
// 	"m2":       "square metre",
// 	"km":       "kilometre",
// 	"in":       "inch",
// 	"ft":       "foot",
// 	"yd":       "yard",
// 	"mi":       "mile",
// 	"ac":       "acre",
// 	"ha":       "hectare",
// 	"mg":       "milligram",
// 	"g":        "gram",
// 	"kg":       "kilogram",
// 	"t":        "tonne",
// 	"lb":       "pound",
// 	"oz":       "ounce",
// 	"st":       "stone",
// 	"s":        "second",
// 	"min":      "minute",
// 	"h":        "hour",
// 	"d":        "day",
// 	"wk":       "week",
// 	"mo":       "month",
// 	"yr":       "year",
// 	"ml":       "millilitre",
// 	"l":        "litre",
// 	"m3":       "cubic metre",
// 	"km3":      "cubic kilometre",
// 	"ft3":      "cubic foot",
// 	"yd3":      "cubic yard",
// 	"ac ft":    "acre foot",
// 	"ac in":    "acre inch",
// 	"US gal":   "US gallon",
// 	"US qt":    "US quart",
// 	"US pt":    "US pint",
// 	"US fl oz": "US fluid ounce",
// 	"mm3":      "cubic millimetre",
// 	"cm3":      "cubic centimetre",
// 	"dm3":      "cubic decimetre",
// 	"Ml":       "megalitre",
// 	"hl":       "hectolitre",
// 	"dal":      "decalitre",
// 	"dl":       "decilitre",
// 	"cl":       "centilitre",
// 	"mm2":      "square millimetre",
// 	"cm2":      "square centimetre",
// 	"km2":      "square kilometre",
// 	"in2":      "square inch",
// 	"ft2":      "square foot",
// 	"yd2":      "square yard",
// 	"mi2":      "square mile",
// 	"in3":      "cubic inch",
// }
//
// type UnitLabel string
//
// // NewUnitLabel returns a UnitLabel for a unit enum
// // The arg could be an actual unit enum, a string representing a unit enum, or a straight-up string label
// func NewUnitLabel(enum any) (UnitLabel, error) {
// 	switch enum.(type) {
// 	case AreaUnit:
// 		return UnitLabel(areaLabels[enum]), nil
// 	case LineUnit:
// 		return UnitLabel(linearLabels[enum]), nil
// 	case MassUnit:
// 		return UnitLabel(massLabels[enum]), nil
// 	case TimeUnit:
// 		return UnitLabel(timeLabels[enum]), nil
// 	case VolumeUnit:
// 		return UnitLabel(volumeLabels[enum]), nil
// 	case string:
// 		return newUnitLabelFromString(enum.(string))
// 	default:
// 		return "??", fmt.Errorf("unknown unit enum %v", enum)
// 	}
// }
//
// // newUnitLabelFromString returns a UnitLabel for a string representing a unit enum
// func newUnitLabelFromString(s string) (UnitLabel, error) {
// 	if s == "" {
// 		return "", fmt.Errorf("empty string")
// 	}
// 	if enum := standardLabelFromMaps(s); enum != nil {
// 		return UnitLabel(s), nil
// 	}
// 	return "", fmt.Errorf("unknown unit label %s", s)
// }
//
// // String returns the standard string label for a unit enum
// func (u UnitLabel) String() string {
// 	return string(u)
// }
//
// func (u UnitLabel) WithSuperscript() string {
// 	if s, ok := labelsWithSuperscript[u.String()]; ok {
// 		return s
// 	}
// 	return u.String()
// }
//
// func (u UnitLabel) FullWord() string {
// 	if s, ok := labelsFullWord[u.String()]; ok {
// 		return s
// 	}
// 	return u.String()
// }
//
