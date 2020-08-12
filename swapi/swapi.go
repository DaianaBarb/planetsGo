package swapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type searchResponse struct {
	Results []planetResponse `json:"results"`
}

type planetResponse struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}
type SWAPI struct {
	APIURL string
}

func (s SWAPI) CountPlanetAppearancesOnMovies(ctx context.Context, planetName string) (int, error) {
	var swapi SWAPI
	swapi.APIURL = "https://swapi.dev/api/planets/?search="

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, swapi.APIURL+planetName, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var searchResponse searchResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return 0, err
	}

	if len(searchResponse.Results) == 0 {
		return 0, nil
	}

	for _, p := range searchResponse.Results {
		if strings.EqualFold(p.Name, planetName) {
			return len(p.Films), nil
		}
	}

	return 0, nil
}
