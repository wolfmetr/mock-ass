#!/bin/bash

make run 2> server.log & tail -f server.log | while read LOGLINE
do
    found=$(echo "${LOGLINE}" | grep -m 1 -q "Start server")
     if echo $LOGLINE | grep -m 1 -q "Start server"; then
        pkill -P $$ tail
     fi
done
rm server.log


trap "killall -v mock-ass" EXIT
host="http://localhost:8000"

SESSINIT=$(curl -s -X POST \
  $( printf "$host/init/?content_type=application/json&session_ttl_min=13" ) \
  -H 'content-type: text/plain' \
  -d '{
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
    }')

if [ -z "$SESSINIT" ]; then
    echo "Response is empty"
    exit 1
fi

echo "Init response = $SESSINIT"
session_path=$(echo "$SESSINIT" | grep -o '"url":"[^"]*' | cut -s -f 4 -d '"')
echo "Parsed session_path = $session_path"

REQ=$( printf "$host$session_path" )
echo "Get rendered template url: $REQ"
RESULT=$(curl -s -L $REQ)
if [ -z "$RESULT" ]; then
    echo "Response is empty"
    exit 1
fi
echo "Rendered template: \n$RESULT"

ip=$(echo "$RESULT" | grep -o '"ip_v4": "[^"]*' | cut -s -f 4 -d '"')

# dummy validation
if [[ $ip =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "success: ip '$ip'"
else
  echo "fail: ip '$ip'"
  exit 1
fi
