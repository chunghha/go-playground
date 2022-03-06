package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type StatesResponse struct {
	States []State `json:"states"`
}

type State struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GegSStates func to return a list of states.
func GetStates() (StatesResponse, error) {
	jsonFile, err := os.Open("./apps/api/states.json")
	if err != nil {
		return StatesResponse{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var stateResponse StatesResponse
	json.Unmarshal(byteValue, &stateResponse)

	return stateResponse, nil
}
