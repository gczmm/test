package util

import "encoding/json"

func Encode(data interface{}) string  {
	resp, _ := json.Marshal(data)
	return string(resp)
}

func Decode(source []byte, target interface{}) error  {
	err := json.Unmarshal(source, target)
	if err != nil {
		return err
	}
	return nil
}
