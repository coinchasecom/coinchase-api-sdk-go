package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
)

// Sign sign the request data
func Sign(mapParams map[string]string, method, hostname, path, secretKey string) string {
	mapCloned := make(map[string]string)
	for key, value := range mapParams {
		mapCloned[key] = url.QueryEscape(value)
	}

	strParams := Map2UrlQueryBySort(mapCloned)

	strPayload := method + "\n" + hostname + "\n" + path + "\n" + strParams
	return ComputeHmac256(strPayload, secretKey)
}

// Map2UrlQueryBySort format map params to a string
func Map2UrlQueryBySort(mapParams map[string]string) string {
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var strParams string
	for _, key := range keys {
		strParams += key + "=" + mapParams[key] + "&"
	}

	// remove "&" at the end of line
	if len(strParams) > 0 {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// ComputeHmac256 compute HMAC SHA256
func ComputeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
