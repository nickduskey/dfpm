package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Npms struct {
	Version string
}

func (n *Npms) Search(query string) ([]NpmsSearchResult, error) {
	formattedQuery := strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("https://api.npms.io/v2/search?=%s", formattedQuery)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var npmsRes NpmsSearchResponse
	err = json.Unmarshal(bodyBytes, &npmsRes)
	if err != nil {
		return nil, err
	}
	return npmsRes.Results, nil
}

func (n *Npms) GetPackageDetails(packageName string) (NpmsPackageResponse, error) {
	url := fmt.Sprintf("https://api.npms.io/v2/package/%s", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return NpmsPackageResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NpmsPackageResponse{}, err
	}

	var npmsRes NpmsPackageResponse
	err = json.Unmarshal(bodyBytes, &npmsRes)
	if err != nil {
		return npmsRes, err
	}
	return npmsRes, nil
}
