package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPResp asdasd asd
type HTTPResp struct {
	Resp *http.Response
	Err  error
}

func sendPOST(mapData []map[string]interface{}) []*HTTPResp {
	var req *http.Request
	var err error
	ch := make(chan *HTTPResp)
	responses := []*HTTPResp{}

	for _, data := range mapData {
		if data["method"] == "POST" {
			if data["request_body"] == "JSON" {
				datajs, err := json.Marshal(data["data"])
				if err != nil {
					log.Println(err)
				}

				respData := strings.ReplaceAll(string(datajs), "{NUMBER}", data["phone"].(string))
				req, err = http.NewRequest("POST", data["URL"].(string), bytes.NewBuffer([]byte(respData)))
				req.Header.Set("Content-Type", "application/json")

			} else {
				formData := url.Values{}
				for k, v := range data["data"].(map[string]interface{}) {
					if v == "{NUMBER}" {
						formData.Add(string(k), data["phone"].(string))
					} else {
						formData.Add(string(k), v.(string))
					}
				}
				req, err = http.NewRequest("POST", data["URL"].(string), strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}
			checkErr(err)

			headers := data["headers"]
			for k, v := range headers.(map[string]interface{}) {
				req.Header.Set(string(k), v.(string))
			}
		} else if data["method"] == "GET" {
			req, err = http.NewRequest("GET", data["URL"].(string), nil)
			if err != nil {
				log.Print(err)
			}
			q := req.URL.Query()
			for k, v := range data["data"].(map[string]interface{}) {
				if v == "{NUMBER}" {
					q.Add(string(k), data["phone"].(string))
				} else {
					q.Add(string(k), v.(string))
				}
			}
			req.URL.RawQuery = q.Encode()
		}

		go func(req *http.Request) {
			client := &http.Client{}
			resp, err := client.Do(req)
			checkErr(err)
			ch <- &HTTPResp{resp, err}
		}(req)
	}

	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == 50 {
				break
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
		return responses
	}

}
