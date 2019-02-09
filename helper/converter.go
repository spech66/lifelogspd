package helper

import "encoding/json"

// StringListToJSON converts string array to JSON
func StringListToJSON(data *[]string) []byte {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(&data)
	if err != nil {
		panic(err)
	}

	return jsonData
}

// DataToJson converts data to JSON
func DataToJSON(data *interface{}) []byte {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(&data)
	if err != nil {
		panic(err)
	}

	return jsonData
}
