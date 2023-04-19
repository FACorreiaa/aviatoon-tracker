package api

import (
	"fmt"
	"os"
)

func BaseUrl(path string) string {
	url := fmt.Sprintf("http://api.aviationstack.com/v1/%s?access_key=%s", path, os.Getenv("API_CONNECTION"))
	return url
}
