package heredotcom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/krylphi/travels-naming/internal/util"
)

// can be swapped to
//import jsoniter "github.com/json-iterator/go"
//var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	rootUrl            = "https://revgeocode.search.hereapi.com"
	version            = "v1"
	revGeoCodeEndpoint = "revgeocode"
)

const (
	apiKeyEnv = "HERECOM_API_KEY"
)

type API interface {
	RevGeoCode(lat, lon float64) (*RevGeoCodeResponse, error)
}

type api struct {
	client http.Client
	apiKey string
}

func NewAPI(apiKey string) API {
	return &api{
		client: http.Client{},
		apiKey: apiKey,
	}
}

func (a *api) RevGeoCode(lat, lon float64) (*RevGeoCodeResponse, error) {
	req, err := a.buildRevGeoCodeRequest(lat, lon)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer util.CloseOrLog(resp.Body)
	if resp.StatusCode == http.StatusOK {
		return parseResponse[RevGeoCodeResponse](resp.Body)
	}
	return nil, fmt.Errorf("response code: %d = %s", resp.StatusCode, resp.Status)
}

func (a *api) buildRevGeoCodeRequest(lat, lon float64) (*http.Request, error) {
	uri := util.Concat(
		rootUrl, "/", version, "/", revGeoCodeEndpoint,
		"?at=", floatToStr(lat), ",", floatToStr(lon),
		"&apiKey=", a.apiKey,
	)

	return http.NewRequest(http.MethodGet, uri, http.NoBody)
}

func parseResponse[T any](b io.Reader) (*T, error) {
	input, err := io.ReadAll(b)
	if err != nil {
		return nil, err
	}
	var data T
	err = json.Unmarshal(input, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
