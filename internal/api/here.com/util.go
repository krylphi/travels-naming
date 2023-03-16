package heredotcom

import (
	"log"
	"os"
	"strconv"
)

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func GetApiKey() (string, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return "", errMissingApiKey
	}
	err := os.Unsetenv(apiKeyEnv)
	if err != nil {
		log.Printf("failed to unset [%s] environment variable", apiKeyEnv)
	}
	return apiKey, nil
}
