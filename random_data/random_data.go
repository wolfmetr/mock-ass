package random_data

import (
	"fmt"
	"hash/crc64"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	AnyGender = iota
	Male
	Female
)

const (
	CountryNameFormat = iota
	CountryCode2Format
	CountryCode3Format
)

const (
	StateUsaCodeFormat = iota
	StateUsaNameFormat
)

func stringToInt64(str string) int64 {
	buf := []byte(str)
	hash_tbl := crc64.MakeTable(crc64.ISO)
	res := int64(crc64.Checksum(buf, hash_tbl))
	return res
}

type RandomData struct {
	hash       string
	hashInt64  int64
	collection *RandomDataCollection
}

func NewRandomData(hash string, collection *RandomDataCollection) *RandomData {
	hashInt64 := stringToInt64(hash)
	return &RandomData{
		hash:       hash,
		hashInt64:  hashInt64,
		collection: collection,
	}
}

func (rd *RandomData) getFirstName(gender int, src int64) string {
	r := rand.New(rand.NewSource(src))
	switch gender {
	case Male:
		return rd.collection.MaleNames[r.Intn(len(rd.collection.MaleNames))]
	case Female:
		return rd.collection.FemaleNames[r.Intn(len(rd.collection.FemaleNames))]
	default:
		if rd.getBoolean(src) {
			return rd.collection.MaleNames[r.Intn(len(rd.collection.MaleNames))]
		} else {
			return rd.collection.FemaleNames[r.Intn(len(rd.collection.FemaleNames))]
		}
	}
}

func (rd *RandomData) getLastName(src int64) string {
	r := rand.New(rand.NewSource(src))
	return rd.collection.LastNames[r.Intn(len(rd.collection.LastNames))]
}

func (rd *RandomData) getEmail(src int64) string {
	r := rand.New(rand.NewSource(src))
	return fmt.Sprintf("%s.%s.example@%s",
		strings.ToLower(rd.FirstName()),
		strings.ToLower(rd.LastName()),
		rd.collection.EmailDomains[r.Intn(len(rd.collection.EmailDomains))])
}

func (rd *RandomData) getCity(src int64) string {
	r := rand.New(rand.NewSource(src))
	return rd.collection.Countries[r.Intn(len(rd.collection.Countries))].Capital
}

func (rd *RandomData) getCountry(formatCountry int, src int64) string {
	r := rand.New(rand.NewSource(src))
	switch formatCountry {
	case CountryNameFormat:
		return rd.collection.Countries[r.Intn(len(rd.collection.Countries))].Name.Official
	case CountryCode2Format:
		return rd.collection.Countries[r.Intn(len(rd.collection.Countries))].CountryCode2
	case CountryCode3Format:
		return rd.collection.Countries[r.Intn(len(rd.collection.Countries))].CountryCode3
	}
	return rd.collection.Countries[r.Intn(len(rd.collection.Countries))].Name.Official
}

func (rd *RandomData) getStateUsa(stateFormat int, src int64) string {
	r := rand.New(rand.NewSource(src))
	switch stateFormat {
	case StateUsaCodeFormat:
		return rd.collection.States[r.Intn(len(rd.collection.States))].Code
	case StateUsaNameFormat:
		return rd.collection.States[r.Intn(len(rd.collection.States))].State
	}
	return rd.collection.States[r.Intn(len(rd.collection.States))].State
}

func (rd *RandomData) getBoolean(src int64) bool {
	r := rand.New(rand.NewSource(src))
	if r.Intn(2)%2 > 0 {
		return true
	}
	return false
}

func (rd *RandomData) getNumber(src int64, numberRange ...int) int {
	r := rand.New(rand.NewSource(src))
	if len(numberRange) > 1 {
		return r.Intn(numberRange[1]-numberRange[0]) + numberRange[0]
	} else {
		return r.Intn(numberRange[0])
	}
}

func (rd *RandomData) getFloat(src int64, numberRange ...int) (result float64) {
	r := rand.New(rand.NewSource(src))

	if len(numberRange) > 1 {
		result = r.Float64()*float64(numberRange[1]-numberRange[0]) + float64(numberRange[0])
	} else {
		result = r.Float64() * float64(numberRange[0])
	}

	if len(numberRange) > 2 {
		precision := math.Pow(10, float64(numberRange[2]))
		result = float64(int64(result*precision)) / precision
	}
	return
}

func (rd *RandomData) getIPv4(src int64) string {
	r := rand.New(rand.NewSource(src))
	return fmt.Sprintf("%d.%d.%d.%d", r.Intn(256), r.Intn(256), r.Intn(256), r.Intn(256))
}

func (rd *RandomData) getParagraph(src int64) string {
	r := rand.New(rand.NewSource(src))
	return rd.collection.Paragraphs[r.Intn(len(rd.collection.Paragraphs))]
}

func (rd *RandomData) FirstName() string {
	src := time.Now().UnixNano()
	return rd.getFirstName(AnyGender, src)
}

func (rd *RandomData) FirstNameChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getFirstName(AnyGender, src)
}

func (rd *RandomData) FirstNameMale() string {
	src := time.Now().UnixNano()
	return rd.getFirstName(Male, src)
}

func (rd *RandomData) FirstNameMaleChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getFirstName(Male, src)
}

func (rd *RandomData) FirstNameFemale() string {
	src := time.Now().UnixNano()
	return rd.getFirstName(Female, src)
}

func (rd *RandomData) FirstNameFemaleChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getFirstName(Female, src)
}

func (rd *RandomData) LastName() string {
	src := time.Now().UnixNano()
	return rd.getLastName(src)
}

func (rd *RandomData) LastNameChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getLastName(src)
}

func (rd *RandomData) FullName() string {
	return rd.FirstName() + " " + rd.LastName()
}
func (rd *RandomData) FullNameChain(key int) string {
	return rd.FirstNameChain(key) + " " + rd.LastNameChain(key)
}

func (rd *RandomData) FullNameMale() string {
	return rd.FirstNameMale() + " " + rd.LastName()
}

func (rd *RandomData) FullNameMaleChain(key int) string {
	return rd.FirstNameMaleChain(key) + " " + rd.LastNameChain(key)
}

func (rd *RandomData) FullNameFemale() string {
	return rd.FirstNameFemale() + " " + rd.LastName()
}

func (rd *RandomData) FullNameFemaleChain(key int) string {
	return rd.FirstNameFemaleChain(key) + " " + rd.LastNameChain(key)
}

func (rd *RandomData) Email() string {
	src := time.Now().UnixNano()
	return rd.getEmail(src)
}

func (rd *RandomData) EmailChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getEmail(src)
}

func (rd *RandomData) City() string {
	src := time.Now().UnixNano()
	return rd.getCity(src)
}

func (rd *RandomData) CityChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getCity(src)
}

func (rd *RandomData) FullCountry() string {
	src := time.Now().UnixNano()
	return rd.getCountry(CountryNameFormat, src)
}

func (rd *RandomData) FullCountryChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getCountry(CountryNameFormat, src)
}

func (rd *RandomData) CountryCode2() string {
	src := time.Now().UnixNano()
	return rd.getCountry(CountryCode2Format, src)
}

func (rd *RandomData) CountryCode2Chain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getCountry(CountryCode2Format, src)
}

func (rd *RandomData) CountryCode3() string {
	src := time.Now().UnixNano()
	return rd.getCountry(CountryCode3Format, src)
}

func (rd *RandomData) CountryCode3Chain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getCountry(CountryCode3Format, src)
}

func (rd *RandomData) StateUsaCode() string {
	src := time.Now().UnixNano()
	return rd.getStateUsa(StateUsaCodeFormat, src)
}

func (rd *RandomData) StateUsaCodeChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getStateUsa(StateUsaCodeFormat, src)
}

func (rd *RandomData) StateUsaName() string {
	src := time.Now().UnixNano()
	return rd.getStateUsa(StateUsaNameFormat, src)
}

func (rd *RandomData) StateUsaNameChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getStateUsa(StateUsaNameFormat, src)
}

func (rd *RandomData) Boolean() bool {
	src := time.Now().UnixNano()
	return rd.getBoolean(src)
}

func (rd *RandomData) BooleanChain(key int) bool {
	src := int64(key) + rd.hashInt64
	return rd.getBoolean(src)
}

func (rd *RandomData) BooleanString() string {
	if rd.Boolean() {
		return "true"
	}
	return "false"
}

func (rd *RandomData) BooleanStringChain(key int) string {
	if rd.BooleanChain(key) {
		return "true"
	}
	return "false"
}

func (rd *RandomData) Number(numberRange ...int) int {
	src := time.Now().UnixNano()
	return rd.getNumber(src, numberRange...)
}

func (rd *RandomData) NumberChain(key int, numberRange ...int) int {
	src := int64(key) + rd.hashInt64
	return rd.getNumber(src, numberRange...)
}

func (rd *RandomData) NumberString(numberRange ...int) string {
	return strconv.Itoa(rd.Number(numberRange...))
}

func (rd *RandomData) NumberStringChain(key int, numberRange ...int) string {
	return strconv.Itoa(rd.NumberChain(key, numberRange...))
}

func (rd *RandomData) Float(numberRange ...int) float64 {
	src := time.Now().UnixNano()
	return rd.getFloat(src, numberRange...)
}

func (rd *RandomData) FloatChain(key int, numberRange ...int) float64 {
	src := int64(key) + rd.hashInt64
	return rd.getFloat(src, numberRange...)
}

func (rd *RandomData) IPv4() string {
	src := time.Now().UnixNano()
	return rd.getIPv4(src)
}

func (rd *RandomData) IPv4Chain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getIPv4(src)
}

func (rd *RandomData) Paragraph() string {
	src := time.Now().UnixNano()
	return rd.getParagraph(src)
}

func (rd *RandomData) ParagraphChain(key int) string {
	src := int64(key) + rd.hashInt64
	return rd.getParagraph(src)
}
