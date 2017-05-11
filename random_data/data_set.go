package random_data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
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

func GetDataDir() string {
	data_path := os.Getenv("MOCK_ASS_DATA_DIR")
	if data_path == "" {
		log.Fatal(color.RedString("Environment variable MOCK ASS_DATA_DIR is empty"))
	}
	return data_path
}

func LoadCountriesFromFile(file_path string) (countries []CountryType, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(file, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}

func LoadLanguagesFromFile(file_path string) (languages []LanguageType, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &languages); err != nil {
		return nil, err
	}
	return languages, nil
}

func LoadStatesFromFile(file_path string) (states []StateType, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &states); err != nil {
		return nil, err
	}
	return states, nil
}

func LoadFemaleNamesFromFile(file_path string) (femaleNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &femaleNames); err != nil {
		return nil, err
	}
	return femaleNames, nil
}

func LoadMaleNamesFromFile(file_path string) (maleNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &maleNames); err != nil {
		return nil, err
	}
	return maleNames, nil
}

func LoadLastNamesFromFile(file_path string) (lastNames []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &lastNames); err != nil {
		return nil, err
	}
	return lastNames, nil
}

func LoadEmailDomainsFromFile(file_path string) (emailDomains []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &emailDomains); err != nil {
		return nil, err
	}
	return emailDomains, nil
}

func LoadParagraphsFromFile(file_path string) (paragraphs []string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(file_path)
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

func InitCollection() (*RandomDataCollection, error) {
	countries, err := LoadCountriesFromFile(GetDataDir() + "/countries.json")
	if err != nil {
		return nil, err
	}
	languages, err := LoadLanguagesFromFile(GetDataDir() + "/languages.json")
	if err != nil {
		return nil, err
	}
	states, err := LoadStatesFromFile(GetDataDir() + "/usa_states.json")
	if err != nil {
		return nil, err
	}
	femaleNames, err := LoadFemaleNamesFromFile(GetDataDir() + "/female_names.json")
	if err != nil {
		return nil, err
	}
	maleNames, err := LoadMaleNamesFromFile(GetDataDir() + "/male_names.json")
	if err != nil {
		return nil, err
	}
	lastNames, err := LoadLastNamesFromFile(GetDataDir() + "/last_names.json")
	if err != nil {
		return nil, err
	}
	emailDomains, err := LoadEmailDomainsFromFile(GetDataDir() + "/email_domains.json")
	if err != nil {
		return nil, err
	}
	paragraphs, err := LoadParagraphsFromFile(GetDataDir() + "/texts.json")
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
