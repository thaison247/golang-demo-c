package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Result struct {
	Code    int                                 `json:"code"`
	Message string                              `json:"message"`
	Data    map[string][]map[string]interface{} `json:"data"`
}

func GetAPI(url string, headers map[string]string) (int, map[string]interface{}) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}
	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Call API
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("API response %d", resp.StatusCode)
		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Print(scanner.Text())
		}
		return resp.StatusCode, nil
	}
	defer resp.Body.Close()

	// Parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil
	}

	var jResult map[string]interface{}
	json.Unmarshal(body, &jResult)
	return http.StatusOK, jResult
}

func PostAPI(url string, headers map[string]string, bodyRequest map[string]interface{}) (int, Result) {
	client := &http.Client{}
	var result Result
	newbodyRequest, err := json.Marshal(bodyRequest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(newbodyRequest))
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}
	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Call API
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		fmt.Println("body: ", buf.String())
		log.Errorf("API response %d", resp.StatusCode)
		return resp.StatusCode, result
	}
	defer resp.Body.Close()

	// Parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, result
	}

	result = Result{}
	json.Unmarshal(body, &result)
	return http.StatusOK, result
}
