package util

import "encoding/json"

func StructToMap(s interface{}) (map[string]interface{}, error) {
	struct2byte, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(struct2byte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
