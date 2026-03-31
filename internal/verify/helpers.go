package verify

import (
	"encoding/json"
	"os"
)

func ReadVerifyData(filename string) (*VerifyData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var verifyData VerifyData
	err = json.Unmarshal(data, &verifyData)
	if err != nil {
		return nil, err
	}

	return &verifyData, nil
}