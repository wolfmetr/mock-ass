package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var testTpl = `{
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

func TestRender(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata")

	t.Logf("init testdata from %s", path)
	collection, err := InitCollectionFromPath(path)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}
	out, err := Render(testTpl, "tet hash", collection)
	if err != nil {
		t.Fatalf("Got err %+v", err)
	}

	fmt.Println(out)
}
