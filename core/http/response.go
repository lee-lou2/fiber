package http

import "encoding/json"

func ResponseToJson(response *string) *map[string]interface{} {
	/*
		Request 결과 값 json 형태로 변환
	*/
	responseBody := make(map[string]interface{})

	err := json.Unmarshal([]byte(string(*response)), &responseBody)

	if err != nil {
		panic(err)
	}

	return &responseBody
}
