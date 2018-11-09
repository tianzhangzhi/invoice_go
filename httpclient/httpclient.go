package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	//"log"
	//"net"
	"net/http"
	"os"
	//"time"
)

/*
var (
	httpClient *http.Client
)

func init() {
	httpClient = createHTTPClient()
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdelConnTimeout     int = 90
)

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdelConnTimeout) * time.Second,
		},
		Timeout: 20 * time.Second,
	}
	return client
}
*/
func main() {

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token", nil)
	if err != nil {
		fmt.Println("req error")
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("grant_type", "client_credential")
	q.Add("appid", "wx10ba04e7f559ba31")
	q.Add("secret", "a894883b0f05d27358bdf3dde85113be")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp error")
	} else {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("resp read error")
		}
		fmt.Println("resp body: %s", string(body))
	}
}
