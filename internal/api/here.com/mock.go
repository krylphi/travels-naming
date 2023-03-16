package heredotcom

import (
	"bytes"
)

type MockAPI struct {
}

func (m MockAPI) RevGeoCode(_, _ float64) (*RevGeoCodeResponse, error) {
	// simulated response taken from https://www.here.com/learn/blog/reverse-geocoding-a-location-using-golang
	data := `{
  "items": [
    {
      "title": "Creativity",
      "id": "here:pds:place:356jx7ps-6b18784695c70f5001df83da965f2570",
      "resultType": "place",
      "address": {
        "label": "Creativity, Shivajinagar, Bengaluru 560001, India",
        "countryCode": "IND",
        "countryName": "India",
        "stateCode": "KA",
        "state": "Karnataka",
        "county": "Bengaluru",
        "city": "Bengaluru",
        "district": "Shivajinagar",
        "postalCode": "560001"
      },
      "position": {
        "lat": 12.98023,
        "lng": 77.60094
      },
      "access": [
        {
          "lat": 12.98022,
          "lng": 77.60093
        }
      ],
      "distance": 11,
      "categories": [
        {
          "id": "800-8200-0174",
          "name": "Школа",
          "primary": true
        }
      ]
    }
  ]
}`
	reader := &bytes.Buffer{}
	reader.WriteString(data)

	return parseResponse[RevGeoCodeResponse](reader)
}
