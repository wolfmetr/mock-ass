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

// http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
var Countries []CountryType
var Languages []LanguageType
var States []StateType
var FemaleNames []string
var MaleNames []string
var LastNames []string
var EmailDomains []string
var Paragraphs []string

func GetDataDir() string {
	data_path := os.Getenv("MOCK_ASS_DATA_DIR")
	if data_path == "" {
		log.Fatal(color.RedString("Environment variable MOCK ASS_DATA_DIR is empty"))
	}
	return data_path
}

func LoadCountriesFromFile(file_path string) {
	log.Println(color.BlueString("Load countries from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &Countries); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadCountries() {
	LoadCountriesFromFile(GetDataDir() + "/countries.json")
}

func LoadLanguagesFromFile(file_path string) {
	log.Println(color.BlueString("Load languages from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &Languages); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadLanguages() {
	LoadLanguagesFromFile(GetDataDir() + "/languages.json")
}

func LoadStatesFromFile(file_path string) {
	log.Println(color.BlueString("Load states from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &States); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadStates() {
	LoadStatesFromFile(GetDataDir() + "/usa_states.json")
}

func LoadFemaleNamesFromFile(file_path string) {
	log.Println(color.BlueString("Load female names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &FemaleNames); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadFemaleNames() {
	LoadFemaleNamesFromFile(GetDataDir() + "/female_names.json")
}

func LoadMaleNamesFromFile(file_path string) {
	log.Println(color.BlueString("Load male names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &MaleNames); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadMaleNames() {
	LoadMaleNamesFromFile(GetDataDir() + "/male_names.json")
}

func LoadLastNamesFromFile(file_path string) {
	log.Println(color.BlueString("Load last names from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &LastNames); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadLastNames() {
	LoadLastNamesFromFile(GetDataDir() + "/last_names.json")
}

func LoadEmailDomainsFromFile(file_path string) {
	log.Println(color.BlueString("Load email domains from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &EmailDomains); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadEmailDomains() {
	LoadEmailDomainsFromFile(GetDataDir() + "/email_domains.json")
}

func LoadParagraphsFromFile(file_path string) {
	log.Println(color.BlueString("Load paragraphs from file %s", file_path))
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(color.RedString("File error %v\n ", err))
	}
	if err := json.Unmarshal(file, &Paragraphs); err != nil {
		log.Fatal(color.RedString("Unmarshal error %v\n ", err))
	}
}

func LoadParagraphs() {
	LoadParagraphsFromFile(GetDataDir() + "/texts.json")
}

func InitWithDefaults() {
	LoadCountries()
	LoadStates()
	LoadLanguages()
	LoadLastNames()
	LoadFemaleNames()
	LoadMaleNames()
	LoadEmailDomains()
	LoadParagraphs()
}
