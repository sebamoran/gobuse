package getabuse

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Get_Abuse_Daily(addr string, texto *string, c chan int) {
	params := url.Values{}
	// params.Add("ipAddress%3D118.25.6.39", ``)
	params.Add("maxAgeInDays", `90`)
	params.Add("verbose", ``)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/check?ipAddress="+addr, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Key", os.ExpandEnv("9c43d10e8583f9d33b8b70b95d5593990b1f01b0ecd80e2bc9bb4c989622719752f08765dd9f2da3"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle err
		fmt.Println(err)
	}

	bodys, _ := io.ReadAll(resp.Body)

	//fmt.Println(string(bodys))
	if strings.Contains("error", string(bodys)){
		*texto += string(bodys)
	}
	fmt.Println(<-c)
	defer resp.Body.Close()
}
