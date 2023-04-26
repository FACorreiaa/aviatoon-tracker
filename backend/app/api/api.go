package api

import (
	"io/ioutil"
	"net/http"
)

func GetAPIData(url string) ([]byte, error) {
	httpClient := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//q := req.URL.Query()
	//q.Add("access_key", "YOUR_ACCESS_KEY")
	//req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
