package utils

import "errors"

func InterfaceToMap(i interface{}) (map[string]interface{}, error) {
	if m, ok := i.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("value is not of map type")
}
