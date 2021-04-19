package getabuse

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"strconv"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Get_Last_Abuse(addr string, texto *string, c chan int) {
	params := url.Values{}
	// params.Add("ipAddress%3D118.25.6.39", ``)
	params.Add("maxAgeInDays", `90`)
	params.Add("verbose", ``)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/check?ipAddress="+addr, body)
	if err != nil {
		// handle err
		fmt.Println(err)
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

	body_res, _ := io.ReadAll(resp.Body)

	//fmt.Println(string(bodys))
	if !strings.Contains("error", string(body_res)){
		//fmt.Println("non contiene errori")
		ip := gjson.Get(string(body_res),"data.ipAddress")
		country := gjson.Get(string(body_res),"data.countryCode")
		score := gjson.Get(string(body_res),"data.abuseConfidenceScore")
		domain := gjson.Get(string(body_res),"data.domain")
		ab_ioc,_ := sjson.Set("","IoC.ip",ip.String())
		ab_ioc,_ = sjson.Set(ab_ioc,"IoC.country",country.String())
		ab_ioc,_ = sjson.Set(ab_ioc,"IoC.score",score.String())
		ab_ioc,_ = sjson.Set(ab_ioc,"IoC.domain",domain.String())
		*texto += ab_ioc
		//*texto += string(body_res)
	}
	fmt.Println(<-c)
	defer resp.Body.Close()
}

func Get_Score_Abuse(score int){
	params := url.Values{}

        body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/blacklist", body)
        if err != nil {
                // handle err
                fmt.Println(err)
        }
	req.Header.Set("confidenceMinimum", strconv.Itoa(score))
        req.Header.Set("Key", os.ExpandEnv("80132f8b7acf47c00cf107ee7a33c965a3cb3de45cba7750df4fcf44cc66301795d65c993aa4c5bd"))
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

        body_res, _ := io.ReadAll(resp.Body)

        fmt.Println(string(body_res))
        if !strings.Contains("error", string(body_res)){
                //fmt.Println("non contiene errori")
                ip := gjson.Get(string(body_res),"data.ipAddress")
                country := gjson.Get(string(body_res),"data.countryCode")
                score := gjson.Get(string(body_res),"data.abuseConfidenceScore")
                domain := gjson.Get(string(body_res),"data.domain")
                ab_ioc,_ := sjson.Set("","IoC.ip",ip.String())
                ab_ioc,_ = sjson.Set(ab_ioc,"IoC.country",country.String())
                ab_ioc,_ = sjson.Set(ab_ioc,"IoC.score",score.String())
                ab_ioc,_ = sjson.Set(ab_ioc,"IoC.domain",domain.String())
                //*texto += string(body_res)
        }
        defer resp.Body.Close()	

}
