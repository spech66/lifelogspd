package helper

import "encoding/json"

// DataToJSON converts data to JSON
func DataToJSON(data interface{}) []byte {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(&data)
	if err != nil {
		panic(err)
	}

	return jsonData
}
