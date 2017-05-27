package random_data

import (
	"fmt"
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

const errLoadFmt = "error on load %s: %v"

func InitCollectionFromPath(dataPath string) (*RandomDataCollection, error) {
	countries, err := LoadCountriesFromFile(filepath.Join(dataPath, CountriesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, CountriesFile, err)
	}
	languages, err := LoadLanguagesFromFile(filepath.Join(dataPath, LanguagesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, LanguagesFile, err)
	}
	states, err := LoadStatesFromFile(filepath.Join(dataPath, StatesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, StatesFile, err)
	}
	femaleNames, err := LoadFemaleNamesFromFile(filepath.Join(dataPath, FemaleNamesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, FemaleNamesFile, err)
	}
	maleNames, err := LoadMaleNamesFromFile(filepath.Join(dataPath, MaleNamesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, MaleNamesFile, err)
	}
	lastNames, err := LoadLastNamesFromFile(filepath.Join(dataPath, LastNamesFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, LastNamesFile, err)
	}
	emailDomains, err := LoadEmailDomainsFromFile(filepath.Join(dataPath, EmailDomainsFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, EmailDomainsFile, err)
	}
	paragraphs, err := LoadParagraphsFromFile(filepath.Join(dataPath, ParagraphsFile))
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, ParagraphsFile, err)
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
		return nil, fmt.Errorf(errLoadFmt, "countries", err)
	}
	languages, err := LoadLanguagesFromBytes(languagesBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "languages", err)
	}
	states, err := LoadStatesFromBytes(statesBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "states", err)
	}
	femaleNames, err := LoadFemaleNamesFromBytes(femaleNamesBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "femaleNames", err)
	}
	maleNames, err := LoadMaleNamesFromBytes(maleNamesBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "maleNames", err)
	}
	lastNames, err := LoadLastNamesFromBytes(lastNamesBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "lastNames", err)
	}
	emailDomains, err := LoadEmailDomainsFromBytes(emailDomainsBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "emailDomains", err)
	}
	paragraphs, err := LoadParagraphsFromBytes(paragraphsBytes)
	if err != nil {
		return nil, fmt.Errorf(errLoadFmt, "paragraphs", err)
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
