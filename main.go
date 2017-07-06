package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"crypto/tls"
	"time"
	"net"
	"github.com/mark-church/temp/types"
)

type Login struct {
	url			string
	username	string
	pass		string
}

func args() *Login {
	if len(os.Args) != 4 {
		panic("Incorrect arguments")
	}

	return &Login{
		url:		os.Args[1],
		username:	os.Args[2],
		pass:		os.Args[3],
	}
}

func main() {

	login := args()

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 10 * time.Second,
		}).Dial,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		MaxIdleConnsPerHost:   5,
	}

	client := &http.Client{ Transport: tr}

	req, err := http.NewRequest("GET", login.url, nil)
	req.SetBasicAuth(login.username, login.pass)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	header := resp.Header
	for key, value := range header {
		fmt.Println(key, ",", value)
	}

	fmt.Printf("%s", body)
}