package gql

import "encoding/json"

func getGraphqlBody(query, operationName string, variables map[string]any) (string, error) {
	varsJson, err := json.Marshal(variables)
	if err != nil {
		return "", err
	}
	result := map[string]interface{}{
		"query":         query,
		"operationName": operationName,
		"variables":     string(varsJson),
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultJson), err
}
