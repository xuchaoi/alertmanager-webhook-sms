package util

import (
	"crypto/tls"
	"io"
	"k8s.io/klog"
	"net/http"
	"time"
)

func InsecureHttpsGet(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	res, err := client.Get(url)
	return res, err
}

func HttpGet(url string) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Get(url)
	return res, err
}

func HttpPost(url string, body io.Reader) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Post(url, "application/json", body)
	return res, err
}

func SMSHttpPost(url string, body io.Reader) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-Channel-Id", "SDNGWF")
	t := time.Now()
	request.Header.Add("X-Trans-Id",t.Format("20060102150405"))
	klog.Infof("[Debug] header: %v", request.Header)
	res, err := client.Do(request)
	return res, err
}