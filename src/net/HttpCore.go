package net

import (
	"net/http"
	"reptile/src/util"
	"io/ioutil"
	"io"
)

var (
	log       = util.Log{}
	TAG       = "HTTP_CORE"
	client    = &http.Client{}
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36"
)

//发送GET请求，返回string
func GetRequest(url string) (string, error) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", UserAgent)
	resp, error := client.Do(request)
	if error != nil {
		log.I(TAG, error.Error())
		return "", error
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.I(TAG, "StatusCode="+string(resp.StatusCode))
		return "", nil
	}
	bytes, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.I(TAG, error.Error())
		return "", error
	}
	responseStr := string(bytes)
	return responseStr, nil
}

//发送GET请求，返回Body
func GetRequestForReader(url string) (io.Reader, error) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", UserAgent)
	resp, error := client.Do(request)
	if error != nil {
		log.I(TAG, error.Error())
		return nil, error
	}
	if resp.StatusCode != 200 {
		log.I(TAG, "StatusCode="+string(resp.StatusCode))
		return nil, nil
	}
	return resp.Body, nil
}
