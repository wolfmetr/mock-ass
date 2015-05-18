package generator

import (
	"strings"

	"github.com/wolfmetr/mock-ass/random_data"

	"gopkg.in/flosch/pongo2.v3"
)

func Range(size int) []int {
	sl := make([]int, size, size)
	for i := range sl {
		sl[i] = i + 1
	}
	return sl
}

func GenerateByTemplate(template string, hash string) (out string, err error) {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err
	}
	rd := random_data.NewRandomData(hash)
	out, err = tpl.Execute(pongo2.Context{
		"FirstName": rd.FirstName,

		"FirstNameChain":   rd.FirstNameChain,
		"FirstNameMale":    rd.FirstNameMale,
		"FirstNameFemale":  rd.FirstNameFemale,
		"LastName":         rd.LastName,
		"LastNameChain":    rd.LastNameChain,
		"FullName":         rd.FullName,
		"FullNameChain":    rd.FullNameChain,
		"FullNameMale":     rd.FullNameMale,
		"FullNameFemale":   rd.FullNameFemale,
		"Email":            rd.Email,
		"FullCountry":      rd.FullCountry,
		"TwoCharCountry":   rd.CountryCode2,
		"ThreeCharCountry": rd.CountryCode3,
		"City":             rd.City,
		"StateUsaCode":     rd.StateUsaCode,
		"StateUsaName":     rd.StateUsaName,
		"Number":           rd.Number,
		"NumberString":     rd.NumberString,
		"Decimal":          rd.Float,
		"Float":            rd.Float,
		"Boolean":          rd.BooleanString,
		"Paragraph":        rd.Paragraph,
		"IPv4":             rd.IPv4,
		"Range":            Range,
		"hash":             hash,
	})
	if err != nil {
		return "", err
	}
	out = strings.Replace(out, "\n", "\\n", -1)
	return out, nil
}
