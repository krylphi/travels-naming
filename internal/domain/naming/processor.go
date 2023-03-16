package naming

import (
	"strconv"
	"time"

	heredotcom "github.com/krylphi/travels-naming/internal/api/here.com"
)

type Processor struct {
	api heredotcom.API
}

func NewProcessor(api heredotcom.API) *Processor {
	return &Processor{
		api: api,
	}
}

func (p *Processor) Process(date, lat, lon string) (string, error) {
	eventDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		eventDate, err = time.Parse(time.DateTime, date)
		if err != nil {
			return "", err
		}
	}

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return "", err
	}

	lonF, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return "", err
	}

	geoData, err := p.api.RevGeoCode(latF, lonF)
	if err != nil {
		return "", err
	}

	if len(geoData.Items) == 0 {
		return "", errNoGeoData
	}

	location := geoData.Items[0].Address.City
	if location == "" {
		location = geoData.Items[0].Address.County
	}

	epithet := EpithetFromTime(eventDate, location)

	return epithet, nil
}
