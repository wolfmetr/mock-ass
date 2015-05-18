package random_data

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/golang/glog"
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
	State string `json:state`
	Code  string `json:"code"`
}

// http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
var Countries []CountryType
var Languages []LanguageType
var States []StateType
var FemaleNames []string
var MaleNames []string
var LastNames []string
var EmailDomains []string
var Paragraphs []string

func LoadCountriesFromFile(file_path string) {
	glog.Info(color.BlueString("Load countries from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &Countries); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadCountries() {
	LoadCountriesFromFile("./data/countries.json")
}

func LoadLanguagesFromFile(file_path string) {
	glog.Info(color.BlueString("Load languages from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &Languages); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadLanguages() {
	LoadLanguagesFromFile("./data/languages.json")
}

func LoadStatesFromFile(file_path string) {
	glog.Info(color.BlueString("Load states from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &States); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadStates() {
	LoadStatesFromFile("./data/usa_states.json")
}

func LoadFemaleNamesFromFile(file_path string) {
	glog.Info(color.BlueString("Load female names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &FemaleNames); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadFemaleNames() {
	LoadFemaleNamesFromFile("./data/female_names.json")
}

func LoadMaleNamesFromFile(file_path string) {
	glog.Info(color.BlueString("Load male names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &MaleNames); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadMaleNames() {
	LoadMaleNamesFromFile("./data/male_names.json")
}

func LoadLastNamesFromFile(file_path string) {
	glog.Info(color.BlueString("Load last names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &LastNames); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadLastNames() {
	LoadLastNamesFromFile("./data/last_names.json")
}

func LoadEmailDomainsFromFile(file_path string) {
	glog.Info(color.BlueString("Load email domains from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &EmailDomains); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadEmailDomains() {
	LoadEmailDomainsFromFile("./data/email_domains.json")
}

func LoadParagraphsFromFile(file_path string) {
	glog.Info(color.BlueString("Load paragraphs from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		glog.Error(color.RedString("File error %v\n ", err))
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &Paragraphs); err != nil {
		glog.Error(color.RedString("Unmarshal error %v\n ", err))
		os.Exit(1)
	}
}

func LoadParagraphs() {
	LoadParagraphsFromFile("./data/texts.json")
}

func InitWithDefaults() {
	LoadCountries()
	LoadStates()
	LoadLanguages()
	LoadLastNames()
	LoadFemaleNames()
	LoadMaleNames()
	LoadEmailDomains()
}
