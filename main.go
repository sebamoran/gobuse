package main

import (
	"fmt"
	"github.com/sebamoran/gobuse/getabuse"
	"github.com/gocolly/colly/v2"
)

func main(){
	c := colly.NewCollector()

	channel := make(chan int)
	
	var addr []string
	var texto string
	c.OnHTML(".col-md-3 > a[href]", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		//fmt.Println(e.Text)
		if len(e.Text) > 4 {
			addr = append(addr,e.Text)
			fmt.Println("dd	->  \n"+e.Text)
			fmt.Println(len(addr))
			go getabuse.Get_Abuse_Daily(e.Text, &texto, channel)
			channel <- len(addr)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.abuseipdb.com/")

	//fmt.Println(addr)
	fmt.Println(texto)

}
