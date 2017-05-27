package random_data

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestWorkingDirectory(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(wd)

	path := filepath.Join(wd, "testdata")
	c, err := InitCollectionFromPath(path)
	if err != nil {
		t.Fatalf("error on InitCollectionFromPath: %v", err)
	}
	if c == nil {
		t.Errorf("cannot init collection from path %s: %v", path, err)
	}

	var src int64 = 1
	r := rand.New(rand.NewSource(src))

	countriesSet := map[string]bool{
		"London":          true,
		"Moscow":          true,
		"Washington D.C.": true,
		"Paris":           true,
	}
	country := c.Country(r).Capital
	if !countriesSet[country] {
		t.Errorf("Capital: actual %s not found", country)
	}

	femaleNamesSet := map[string]bool{
		"Olivia":  true,
		"Grace":   true,
		"Jessica": true,
	}
	femaleName := c.FemaleName(r)
	if !femaleNamesSet[femaleName] {
		t.Errorf("FemaleName: actual %s not found", femaleName)
	}

	maleNamesSet := map[string]bool{
		"Jack":   true,
		"Thomas": true,
		"Joshua": true,
	}
	maleName := c.MaleName(r)
	if !maleNamesSet[maleName] {
		t.Errorf("MaleName: actual %s not found", maleName)
	}

	lastNamesSet := map[string]bool{
		"Smith":    true,
		"Johnson":  true,
		"Williams": true,
	}
	lastName := c.LastName(r)
	if !lastNamesSet[lastName] {
		t.Errorf("LastName: actual %s not found", lastName)
	}

	languagesSet := map[string]bool{
		"English": true,
		"Russian": true,
		"French":  true,
	}
	language := c.Language(r)
	if !languagesSet[language.Language] {
		t.Errorf("Language: actual %s not found", language)
	}

	emailDomainsSet := map[string]bool{
		"hotmail.com":  true,
		"facebook.com": true,
		"gmail.com":    true,
	}
	emailDomain := c.EmailDomain(r)
	if !emailDomainsSet[emailDomain] {
		t.Errorf("EmailDomain: actual %s not found", emailDomain)
	}

	statesSet := map[string]bool{
		"CALIFORNIA": true,
		"ALASKA":     true,
		"COLORADO":   true,
	}
	state := c.State(r).State
	if !statesSet[state] {
		t.Errorf("State: actual %s not found", state)
	}
}
