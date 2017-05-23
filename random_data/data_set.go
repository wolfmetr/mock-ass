package random_data

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

type CountryType struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	CountryCode3 string   `json:"cca3"`
	CountryCode2 string   `json:"cca2"`
	Currency     []string `json:"currency"`
	CallingCode  []string `json:"callingCode"`
	Capital      string   `json:"capital"`
}

type LanguageType struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type StateType struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

func LoadCountriesFromFile(filePath string) (countries []CountryType, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	countries, err = LoadCountriesFromBytes(b)
	return
}

func LoadCountriesFromBytes(b []byte) (countries []CountryType, err error) {
	if err := json.Unmarshal(b, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}

func LoadLanguagesFromFile(filePath string) (languages []LanguageType, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	languages, err = LoadLanguagesFromBytes(b)
	return
}

func LoadLanguagesFromBytes(b []byte) (languages []LanguageType, err error) {
	if err := json.Unmarshal(b, &languages); err != nil {
		return nil, err
	}
	return languages, nil
}

func LoadStatesFromFile(filePath string) (states []StateType, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	states, err = LoadStatesFromBytes(b)
	return
}

func LoadStatesFromBytes(b []byte) (states []StateType, err error) {
	if err := json.Unmarshal(b, &states); err != nil {
		return nil, err
	}
	return states, nil
}

func LoadFemaleNamesFromFile(filePath string) (femaleNames []string, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	femaleNames, err = LoadFemaleNamesFromBytes(b)
	return
}

func LoadFemaleNamesFromBytes(b []byte) (femaleNames []string, err error) {
	if err := json.Unmarshal(b, &femaleNames); err != nil {
		return nil, err
	}
	return femaleNames, nil
}

func LoadMaleNamesFromFile(filePath string) (maleNames []string, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	maleNames, err = LoadMaleNamesFromBytes(b)
	return
}

func LoadMaleNamesFromBytes(b []byte) (maleNames []string, err error) {
	if err := json.Unmarshal(b, &maleNames); err != nil {
		return nil, err
	}
	return maleNames, nil
}

func LoadLastNamesFromFile(filePath string) (lastNames []string, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lastNames, err = LoadLastNamesFromBytes(b)
	return
}

func LoadLastNamesFromBytes(b []byte) (lastNames []string, err error) {
	if err := json.Unmarshal(b, &lastNames); err != nil {
		return nil, err
	}
	return lastNames, nil
}

func LoadEmailDomainsFromFile(filePath string) (emailDomains []string, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	emailDomains, err = LoadEmailDomainsFromBytes(b)
	return
}

func LoadEmailDomainsFromBytes(b []byte) (emailDomains []string, err error) {
	if err := json.Unmarshal(b, &emailDomains); err != nil {
		return nil, err
	}
	return emailDomains, nil
}

func LoadParagraphsFromFile(filePath string) (paragraphs []string, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	paragraphs, err = LoadParagraphsFromBytes(b)
	return
}

func LoadParagraphsFromBytes(b []byte) (paragraphs []string, err error) {
	if err := json.Unmarshal(b, &paragraphs); err != nil {
		return nil, err
	}
	return paragraphs, nil
}

// http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
type RandomDataCollection struct {
	countries    []CountryType
	languages    []LanguageType
	states       []StateType
	femaleNames  []string
	maleNames    []string
	lastNames    []string
	emailDomains []string
	paragraphs   []string
}

func (rdc *RandomDataCollection) Country(r *rand.Rand) *CountryType {
	return &rdc.countries[r.Intn(len(rdc.countries))]
}

func (rdc *RandomDataCollection) EmailDomain(r *rand.Rand) string {
	return rdc.emailDomains[r.Intn(len(rdc.emailDomains))]
}

func (rdc *RandomDataCollection) Language(r *rand.Rand) *LanguageType {
	return &rdc.languages[r.Intn(len(rdc.languages))]
}

func (rdc *RandomDataCollection) State(r *rand.Rand) *StateType {
	return &rdc.states[r.Intn(len(rdc.states))]
}

func (rdc *RandomDataCollection) MaleName(r *rand.Rand) string {
	return rdc.maleNames[r.Intn(len(rdc.maleNames))]
}

func (rdc *RandomDataCollection) FemaleName(r *rand.Rand) string {
	return rdc.femaleNames[r.Intn(len(rdc.femaleNames))]
}

func (rdc *RandomDataCollection) LastName(r *rand.Rand) string {
	return rdc.lastNames[r.Intn(len(rdc.lastNames))]
}

// Default file names; reset if you need.
var (
	CountriesFile    string = "countries.json"
	LanguagesFile    string = "languages.json"
	StatesFile       string = "usa_states.json"
	FemaleNamesFile  string = "female_names.json"
	MaleNamesFile    string = "male_names.json"
	LastNamesFile    string = "last_names.json"
	EmailDomainsFile string = "email_domains.json"
	ParagraphsFile   string = "texts.json"
)

func InitCollectionFromPath(dataPath string) (*RandomDataCollection, error) {
	countries, err := LoadCountriesFromFile(filepath.Join(dataPath, CountriesFile))
	if err != nil {
		return nil, err
	}
	languages, err := LoadLanguagesFromFile(filepath.Join(dataPath, LanguagesFile))
	if err != nil {
		return nil, err
	}
	states, err := LoadStatesFromFile(filepath.Join(dataPath, StatesFile))
	if err != nil {
		return nil, err
	}
	femaleNames, err := LoadFemaleNamesFromFile(filepath.Join(dataPath, FemaleNamesFile))
	if err != nil {
		return nil, err
	}
	maleNames, err := LoadMaleNamesFromFile(filepath.Join(dataPath, MaleNamesFile))
	if err != nil {
		return nil, err
	}
	lastNames, err := LoadLastNamesFromFile(filepath.Join(dataPath, LastNamesFile))
	if err != nil {
		return nil, err
	}
	emailDomains, err := LoadEmailDomainsFromFile(filepath.Join(dataPath, EmailDomainsFile))
	if err != nil {
		return nil, err
	}
	paragraphs, err := LoadParagraphsFromFile(filepath.Join(dataPath, ParagraphsFile))
	if err != nil {
		return nil, err
	}
	return &RandomDataCollection{
		countries:    countries,
		languages:    languages,
		states:       states,
		femaleNames:  femaleNames,
		maleNames:    maleNames,
		lastNames:    lastNames,
		emailDomains: emailDomains,
		paragraphs:   paragraphs,
	}, nil
}

func InitCollectionFromBytes(countriesBytes, languagesBytes, statesBytes, femaleNamesBytes, maleNamesBytes, lastNamesBytes, emailDomainsBytes, paragraphsBytes []byte) (*RandomDataCollection, error) {
	countries, err := LoadCountriesFromBytes(countriesBytes)
	if err != nil {
		return nil, err
	}
	languages, err := LoadLanguagesFromBytes(languagesBytes)
	if err != nil {
		return nil, err
	}
	states, err := LoadStatesFromBytes(statesBytes)
	if err != nil {
		return nil, err
	}
	femaleNames, err := LoadFemaleNamesFromBytes(femaleNamesBytes)
	if err != nil {
		return nil, err
	}
	maleNames, err := LoadMaleNamesFromBytes(maleNamesBytes)
	if err != nil {
		return nil, err
	}
	lastNames, err := LoadLastNamesFromBytes(lastNamesBytes)
	if err != nil {
		return nil, err
	}
	emailDomains, err := LoadEmailDomainsFromBytes(emailDomainsBytes)
	if err != nil {
		return nil, err
	}
	paragraphs, err := LoadParagraphsFromBytes(paragraphsBytes)
	if err != nil {
		return nil, err
	}
	return &RandomDataCollection{
		countries:    countries,
		languages:    languages,
		states:       states,
		femaleNames:  femaleNames,
		maleNames:    maleNames,
		lastNames:    lastNames,
		emailDomains: emailDomains,
		paragraphs:   paragraphs,
	}, nil
}
