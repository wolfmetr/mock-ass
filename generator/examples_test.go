package generator_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wolfmetr/mock-ass/generator"
)

func ExampleRender() {
	type testTplJson struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		EmptyField string `json:"empty_field"`
	}

	// *Chain fields is not actually random. Its has once random seed
	var testTemplateJson = `{
	    "first_name": "{{ FirstNameChain(1) }}",
	    "last_name": "{{ LastNameChain(1) }}",
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

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata")

	collection, err := generator.InitCollectionFromPath(path)
	if err != nil {
		panic(err)
	}

	out, err := generator.Render(testTemplateJson, "my random source hash", collection)
	if err != nil {
		panic(err)
	}

	var parsedTpl testTplJson
	err = json.Unmarshal([]byte(out), &parsedTpl)
	if err != nil {
		panic(err)
	}

	// fmt.Println(out) // true random render
	fmt.Printf("parsedTpl: %+v", parsedTpl)

	// Output:
	// parsedTpl: {FirstName:Grace LastName:Johnson EmptyField:}
}
