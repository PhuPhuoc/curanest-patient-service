package common

import (
	"encoding/json"
)

func ExtractErrorFromResponse(response map[string]interface{}) error {
	// rawErr, ok := response["error"].(map[string]interface{})
	// if !ok {
	// 	resp := NewInternalServerError().WithReason("error response is not in expected format")
	// 	return resp
	// }

	var appErr AppError
	jsonData, err := json.Marshal(response["error"])
	if err != nil {
		resp := NewInternalServerError().WithReason("failed to marshal error response: " + err.Error())
		return resp
	}

	if err := json.Unmarshal(jsonData, &appErr); err != nil {
		resp := NewInternalServerError().WithReason("failed to unmarshal error response: " + err.Error())
		return resp
	}

	return &appErr
}

func ExtractDataFromResponse[T any](response map[string]interface{}) (*T, error) {
	rawData, ok := response["data"].(map[string]interface{})
	if !ok {
		resp := NewInternalServerError().WithReason("data response is not in expected format")
		return nil, resp
	}

	var resp T
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		resp := NewInternalServerError().WithReason("failed to marshal error response: " + err.Error())
		return nil, resp
	}

	if err := json.Unmarshal(jsonData, &resp); err != nil {
		resp := NewInternalServerError().WithReason("failed to unmarshal error response: " + err.Error())
		return nil, resp
	}

	return &resp, nil
}
