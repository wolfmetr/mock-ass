package generator

import (
	"encoding/json"
	"io/ioutil"
)

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
