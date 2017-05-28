package random_data

import (
	"gopkg.in/flosch/pongo2.v3"
)

func Range(size int) []int {
	sl := make([]int, size)
	for i := range sl {
		sl[i] = i + 1
	}
	return sl
}

func Render(template string, hash string, collection *RandomDataCollection) (out string, err error) {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err
	}
	rd := NewRandomData(hash, collection)
	out, err = tpl.Execute(pongo2.Context{
		"FirstName":               rd.FirstName,
		"FirstNameChain":          rd.FirstNameChain,
		"FirstNameMale":           rd.FirstNameMale,
		"FirstNameMaleChain":      rd.FirstNameMaleChain,
		"FirstNameFemale":         rd.FirstNameFemale,
		"FirstNameFemaleChain":    rd.FirstNameFemaleChain,
		"LastName":                rd.LastName,
		"LastNameChain":           rd.LastNameChain,
		"FullName":                rd.FullName,
		"FullNameChain":           rd.FullNameChain,
		"FullNameMale":            rd.FullNameMale,
		"FullNameMaleChain":       rd.FullNameMaleChain,
		"FullNameFemale":          rd.FullNameFemale,
		"FullNameFemaleChain":     rd.FullNameFemaleChain,
		"Email":                   rd.Email,
		"EmailChain":              rd.EmailChain,
		"FullCountry":             rd.FullCountry,
		"FullCountryChain":        rd.FullCountryChain,
		"TwoLetterCountry":        rd.CountryCode2,
		"TwoLetterCountryChain":   rd.CountryCode2Chain,
		"ThreeLetterCountry":      rd.CountryCode3,
		"ThreeLetterCountryChain": rd.CountryCode3Chain,
		"City":               rd.City,
		"CityChain":          rd.CityChain,
		"StateUsaCode":       rd.StateUsaCode,
		"StateUsaCodeChain":  rd.StateUsaCodeChain,
		"StateUsaName":       rd.StateUsaName,
		"StateUsaNameChain":  rd.StateUsaNameChain,
		"Number":             rd.Number,
		"NumberChain":        rd.NumberChain,
		"NumberString":       rd.NumberString,
		"NumberStringChain":  rd.NumberStringChain,
		"Decimal":            rd.Float,
		"DecimalChain":       rd.FloatChain,
		"Float":              rd.Float,
		"FloatChain":         rd.FloatChain,
		"Boolean":            rd.Boolean,
		"BooleanChain":       rd.BooleanChain,
		"BooleanString":      rd.BooleanString,
		"BooleanStringChain": rd.BooleanStringChain,
		"Paragraph":          rd.Paragraph,
		"ParagraphChain":     rd.ParagraphChain,
		"IPv4":               rd.IPv4,
		"IPv4Chain":          rd.IPv4Chain,
		"Range":              Range,
		"hash":               hash,
	})
	if err != nil {
		return "", err
	}
	//out = strings.Replace(out, "\n", "\\n", -1)
	return out, nil
}
