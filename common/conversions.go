package common

import (
	"encoding/json"
)


func Struct2Json(s interface{}) (string, error) {
	if s == nil {
		return "", nil
	} else {
		if e, err := json.Marshal(s); err != nil {
			return "", err
		} else {
			return string(e), nil
		}
	}
}

func Map2Struct(mapData map[string]interface{}, object interface{}) error {
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonStr, object); err != nil {
		return err
	}

	return nil
}
