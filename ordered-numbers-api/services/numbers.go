package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type NumbersService struct {
	apiURL   string
	apiToken string
}

type OrderedNumbersResponse struct {
	OrderedNumbers []int `json:"ordered_numbers"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewNumbersService(apiURL, apiToken string) *NumbersService {
	return &NumbersService{
		apiURL:   apiURL,
		apiToken: apiToken,
	}
}

func (s *NumbersService) FetchOrderedNo() ([]int, error) {
	req, err := http.NewRequest("GET", s.apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode, resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %v", err)
		}
		return nil, fmt.Errorf("API error: %s", errorResponse.Error)
	}

	var data []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	var integers []int
	for _, item := range data {
		switch v:= item.(type) {
		case float64: 
		integers= append(integers, int(v))
		case string:
			if intVal, err:= strconv.Atoi((v)); err== nil {
				integers = append(integers, intVal)
			}
		}
	}
	sort.Slice(integers, func(i, j int) bool {
        a := integers[i]
        b := integers[j]
        if a%2 != 0 && b%2 == 0 {
            return true
        }
        if a%2 == 0 && b%2 != 0 {
            return false
        }
        return a < b
    })

    return integers, nil
}