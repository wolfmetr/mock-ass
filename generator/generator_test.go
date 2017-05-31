package generator

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
)

func initTestCollection(t *testing.T) *RandomDataCollection {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata")

	t.Logf("init testdata from %s", path)
	collection, err := InitCollectionFromPath(path)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}
	return collection
}

var testTemplateJson = `{
    "first_name": "{{ FirstName() }}",
    "last_name": "{{ LastName() }}",
    "full_name": "{{ FullName() }}",
    "city": "{{ City() }}",
    "country": "{{ FullCountry() }}",
    "country2": "{{ TwoLetterCountry() }}",
    "country3": "{{ ThreeLetterCountry() }}",
    "isActive": {{ BooleanString() }},
    "float": {{ Float(12) }},
    "float2": {{ Float(10, 15) }},
    "float3": {{ Float(10, 15, 2) }},
    "float4": {{ Float(10, 15, 3) }},
    "ip_v4": "{{ IPv4() }}",
    "person": [
        {% for x in Range(5) %}
        {
            "first_name": "{{ FirstNameChain(forloop.Counter0) }}",
            "last_name": "{{ LastNameChain(forloop.Counter0) }}",
            "full_name": "{{ FullNameChain(forloop.Counter0) }}",
            "age": {{ Number(10, 100) }},
            "email": "{{ Email() }}"
        }{% if not forloop.Last %}, {% endif %}
        {% endfor %}
    ]
}`

type testTplJson struct {
	EmptyField string `json:"empty_field"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

func TestRenderJson(t *testing.T) {
	collection := initTestCollection(t)

	out, err := Render(testTemplateJson, "tet hash", collection)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}

	t.Log(out)

	var parsedTpl testTplJson
	err = json.Unmarshal([]byte(out), &parsedTpl)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}

	if parsedTpl.FirstName == "" {
		t.Error("FirstName is empty")
	}

	if parsedTpl.LastName == "" {
		t.Error("LastName is empty")
	}

	if parsedTpl.EmptyField != "" {
		t.Error("EmptyField is not empty")
	}
}

var testTemplateXml = `
	<Person>
		<FullName>{{ FullName() }}</FullName>
		<Company>Example Inc.</Company>
		<Email where="home">
			<Addr>{{ Email() }}</Addr>
		</Email>
		<Email where='work'>
			<Addr>{{ Email() }}</Addr>
		</Email>
		<Group>
			<Value>Friends</Value>
			<Value>Squash</Value>
		</Group>
		<City>{{ City() }}</City>
		<State>{{ StateUsaName() }}</State>
	</Person>
`

type testTplXmlEmail struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type testTplXmlAddress struct {
	City, State string
}
type testTplXml struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []testTplXmlEmail
	Groups  []string `xml:"Group>Value"`
	testTplXmlAddress
}

func TestRenderXml(t *testing.T) {
	collection := initTestCollection(t)

	out, err := Render(testTemplateXml, "tet hash", collection)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}

	t.Log(out)

	var parsedTpl testTplXml
	err = xml.Unmarshal([]byte(out), &parsedTpl)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}

	if parsedTpl.XMLName.Local != "Person" {
		t.Error("XMLName is not Person")
	}

	if parsedTpl.Name == "" {
		t.Error("Name is empty")
	}

	if parsedTpl.City == "" {
		t.Error("City is empty")
	}

	if parsedTpl.State == "" {
		t.Error("State is empty")
	}

	for _, email := range parsedTpl.Email {
		if email.Addr == "" {
			t.Error("email is empty")
		}
	}
}
