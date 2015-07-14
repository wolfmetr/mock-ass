# mock-ass
Mock HTTP responses with processed random data and a custom template

### Installation
```bash
$ go get github.com/wolfmetr/mock-ass
```

### Usage
```bash
$ make run
or
$ ./mock-ass [-no-color] [-port=8000] 
```

to initialize session send POST request to `http://localhost:8000/init` with media-type `application/x-www-form-urlencoded`
```
template: {
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
        }
session_ttl_min: 15 (optianal field)
content_type: application/json (optional field)
```
and get response like
```json
{
    "session": "ac8c81bf-75ae-42d4-90c1-de1523acddb7",
    "url": "/session/?s=ac8c81bf-75ae-42d4-90c1-de1523acddb7"
}
```

Then you can get your rendered template from GET `http://localhost:8000/session/?s=ac8c81bf-75ae-42d4-90c1-de1523acddb7`
Every request to `http://localhost:8000/session/?s=...` redirects request with 307 code to url like `http://localhost:8000/session/?s=...&h=...` where `h` is unique hash.
If you send GET request to `http://localhost:8000/session/?s=...&h=...` you'll get cached data, NOT random!

## Template functions
- `FirstName()` — random male/female firstname
- `FirstNameChain(key int)`
- `FirstNameMale()` — random male firstname
- `FirstNameMaleChain(key int)`
- `FirstNameFemale()` — random female firstname
- `FirstNameFemaleChain(key int)`
- `LastName()` — random lastname
- `LastNameChain(key int)`
- `FullName()` — random male/female fullname
- `FullNameChain(key int)`
- `FullNameMale()` — random male fullname
- `FullNameMaleChain(key int)`
- `FullNameFemale()` — random female fullname 
- `FullNameFemaleChain(key int)`
- `Email()` — random email
- `EmailChain(key int)`
- `FullCountry()` — random full country name
- `FullCountryChain(key int)`
- `TwoLetterCountry()` — random two-letter country code (ISO 3166-1 alpha-2)
- `TwoLetterCountryChain(key int)`
- `ThreeLetterCountry()` — random three-letter country code (ISO 3166-1 alpha-3)
- `ThreeLetterCountryChain(key int)`
- `City()` — random city string
- `CityChain(key int)`
- `StateUsaCode()` — random USA state code string
- `StateUsaCodeChain(key int)`
- `StateUsaName()` — random USA state name string
- `StateUsaNameChain(key int)`
- `Number(max_num int)` — random number from range 0 to `max_num`
- `Number(min_num, max_num int)` — random number from range `min_num` to `max_num`
- `NumberChain(key, [min_num,] max_num int)`
- `NumberString([min_num,] max_num int)` — random number string (see Number)
- `NumberStringChain(key, [min_num,] max_num int)`
- `Float([min_float,] max_float int)` — random float from range 0(or min_float) to `max_float`
- `FloatChain(key, [min_float,] max_float int)`
- `Decimal([min_float,] max_float int)` — see Float
- `DecimalChain(key, [min_float,] max_float int)`
- `Boolean()` — random boolean
- `BooleanChain(key int)`
- `BooleanString()` — random boolean string
- `BooleanStringChain(key int)`
- `Paragraph()` — random 'lorem ipsum'-like text
- `ParagraphChain(key int)`
- `IPv4()` — random IPv4 address
- `IPv4Chain(key int)`
- `Range(size int)` — array from 1 to `size`(including)

## Template variables
- hash — unique hash for request

## Contact

- Vladimir Savvateev
- [wolfmetrjob@gmail.com](mailto:wolfmetr@gmail.com)
- [http://twitter.com/Wolfmetrs](http://twitter.com/Wolfmetrs)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
