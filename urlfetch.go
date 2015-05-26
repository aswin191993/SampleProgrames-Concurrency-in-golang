package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"https://github.com/aswin191993/",

}

type HttpResp struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string){
	end:=time.After(10000 * time.Millisecond)
	ch := make(chan *HttpResp)
	responses := []*HttpResp{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HttpResp{url, resp, err}
			}(url)
		}
	for{
		select{
			case r := <-ch:
				fmt.Printf("\n%s was fetched\n", r.url)
				responses = append(responses, r)
				fmt.Printf("status: %s ",r.response.Status)
				if len(responses) == len(urls) {
					return 
				}
			case <-end:
				fmt.Printf("\nSome urls are not featched...\n")
				return 
			default:
				fmt.Printf(".")
				time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	asyncHttpGets(urls)
}

