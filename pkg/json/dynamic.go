package json

import "encoding/json"

func JsonGenerator(jsonData []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
