package main

import (
	"log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatalf("An Error Occured: %v", e)
	}
}

func bomber(phone string) string {
	type M map[string]interface{}
	var headers map[string]interface{}

	JSONData := parseJSON()
	var mapData []map[string]interface{}

	for _, result := range JSONData {
		data := result["data"].(map[string]interface{})
		url := result["URL"].(string)
		mask := result["mask"].(string)
		requestBodyType := result["request_body"].(string)
		method := result["method"].(string)

		maskedPhone := phoneMask(phone, mask)

		if _, ok := result["headers"]; ok {
			headers = result["headers"].(map[string]interface{})
		}

		if _, ok := result["10digits"]; ok {
			maskedPhone = maskedPhone[1:]
		}

		xx := M{"data": data, "URL": url, "request_body": requestBodyType, "phone": maskedPhone, "headers": headers, "method": method}

		mapData = append(mapData, xx)
	}

	sendPOST(mapData)

	return "OK"

}
