package random_data

import (
	"encoding/json"
	"io/ioutil"
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
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(file, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}

func LoadLanguagesFromFile(filePath string) (languages []LanguageType, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &languages); err != nil {
		return nil, err
	}
	return languages, nil
}

func LoadStatesFromFile(filePath string) (states []StateType, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &states); err != nil {
		return nil, err
	}
	return states, nil
}

func LoadFemaleNamesFromFile(filePath string) (femaleNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &femaleNames); err != nil {
		return nil, err
	}
	return femaleNames, nil
}

func LoadMaleNamesFromFile(filePath string) (maleNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &maleNames); err != nil {
		return nil, err
	}
	return maleNames, nil
}

func LoadLastNamesFromFile(filePath string) (lastNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &lastNames); err != nil {
		return nil, err
	}
	return lastNames, nil
}

func LoadEmailDomainsFromFile(filePath string) (emailDomains []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &emailDomains); err != nil {
		return nil, err
	}
	return emailDomains, nil
}

func LoadParagraphsFromFile(filePath string) (paragraphs []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &paragraphs); err != nil {
		return nil, err
	}
	return paragraphs, nil
}

// http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
type RandomDataCollection struct {
	Countries    []CountryType
	Languages    []LanguageType
	States       []StateType
	FemaleNames  []string
	MaleNames    []string
	LastNames    []string
	EmailDomains []string
	Paragraphs   []string
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
	TextsFile        string = "texts.json"
)

func InitCollection(dataPath string) (*RandomDataCollection, error) {
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
	paragraphs, err := LoadParagraphsFromFile(filepath.Join(dataPath, TextsFile))
	if err != nil {
		return nil, err
	}
	return &RandomDataCollection{
		Countries:    countries,
		Languages:    languages,
		States:       states,
		FemaleNames:  femaleNames,
		MaleNames:    maleNames,
		LastNames:    lastNames,
		EmailDomains: emailDomains,
		Paragraphs:   paragraphs,
	}, nil
}
