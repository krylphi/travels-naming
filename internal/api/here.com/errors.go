package heredotcom

import "errors"

var (
	errMissingApiKey = errors.New("HERECOM_API_KEY environment variable must be set")
)
