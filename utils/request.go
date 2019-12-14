package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/coinchasecom/coinchase-api-sdk-go/config"
)

func HttpGetRequest(strUrl string) string {
	httpClient := &http.Client{}

	request, err := http.NewRequest("GET", strUrl, nil)
	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	request.Header.Add("X-CCC-ACCESSKEY", config.AccessKey)

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

func HttpPostRequest(strUrl string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	requestBody, err := json.Marshal(mapParams)
	if err != nil {
		return err.Error()
	}

	request, err := http.NewRequest("POST", strUrl, bytes.NewBuffer(requestBody))
	if nil != err {
		return err.Error()
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-CCC-ACCESSKEY", config.AccessKey)

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

func ApiKeyGet(mapParams map[string]string, strRequestPath string) string {
	method := "GET"
	timestamp := time.Now().Unix()

	mapParams["timestamp"] = strconv.FormatInt(timestamp, 10)
	mapParams["signature"] = Sign(mapParams, method, config.Host, strRequestPath, config.SecretKey)

	params := url.Values{}
	for k, v := range mapParams {
		params.Set(k, v)
	}
	u := url.URL{Scheme: "https", Host: config.Host, Path: strRequestPath}
	u.RawQuery = params.Encode()

	return HttpGetRequest(u.String())
}

func ApiKeyPost(mapParams map[string]string, path string, urlValues url.Values) string {
	method := "POST"
	timestamp := time.Now().Unix()

	mapParams["timestamp"] = strconv.FormatInt(timestamp, 10)
	mapParams["signature"] = Sign(mapParams, method, config.Host, path, config.SecretKey)

	u := url.URL{Scheme: "https", Host: config.Host, Path: path}
	u.RawQuery = urlValues.Encode()

	return HttpPostRequest(u.String(), mapParams)
}
