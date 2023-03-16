package heredotcom

type (
	RevGeoCodeResponse struct {
		Items []RevGeoCodeResponseItem `json:"items"`
	}

	RevGeoCodeResponseItem struct {
		Title                  string       `json:"title"`
		Id                     string       `json:"id"`
		PoliticalView          string       `json:"politicalView"`
		ResultType             string       `json:"resultType"`
		HouseNumberType        string       `json:"houseNumberType"`
		AddressBlockType       string       `json:"addressBlockType"`
		LocalityType           string       `json:"localityType"`
		AdministrativeAreaType string       `json:"administrativeAreaType"`
		Address                Address      `json:"address"`
		Position               Position     `json:"position"`
		Access                 []Access     `json:"access"`
		Distance               int          `json:"distance"`
		MapView                MapView      `json:"mapView"`
		Categories             []Category   `json:"categories"`
		FoodTypes              []FoodType   `json:"foodTypes"`
		HouseNumberFallback    bool         `json:"houseNumberFallback"`
		TimeZone               TimeZone     `json:"timeZone"`
		StreetInfo             []StreetInfo `json:"streetInfo"`
		CountryInfo            CountryInfo  `json:"countryInfo"`
	}

	Address struct {
		Label       string   `json:"label"`
		CountryCode string   `json:"countryCode"`
		CountryName string   `json:"countryName"`
		StateCode   string   `json:"stateCode"`
		State       string   `json:"state"`
		CountyCode  string   `json:"countyCode"`
		County      string   `json:"county"`
		City        string   `json:"city"`
		District    string   `json:"district"`
		SubDistrict string   `json:"subdistrict"`
		Street      string   `json:"street"`
		Streets     []string `json:"streets"`
		Block       string   `json:"block"`
		SubBlock    string   `json:"subblock"`
		PostalCode  string   `json:"postalCode"`
		HouseNumber string   `json:"houseNumber"`
		Building    string   `json:"building"`
	}

	Position struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	Access struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	MapView struct {
		West  int `json:"west"`
		South int `json:"south"`
		East  int `json:"east"`
		North int `json:"north"`
	}

	Category struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Primary bool   `json:"primary"`
	}
	FoodType struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Primary bool   `json:"primary"`
	}
	TimeZone struct {
		Name      string `json:"name"`
		UtcOffset string `json:"utcOffset"`
	}
	StreetInfo struct {
		BaseName           string `json:"baseName"`
		StreetType         string `json:"streetType"`
		StreetTypePrecedes bool   `json:"streetTypePrecedes"`
		StreetTypeAttached bool   `json:"streetTypeAttached"`
		Prefix             string `json:"prefix"`
		Suffix             string `json:"suffix"`
		Direction          string `json:"direction"`
		Language           string `json:"language"`
	}
	CountryInfo struct {
		Alpha2 string `json:"alpha2"`
		Alpha3 string `json:"alpha3"`
	}
)
