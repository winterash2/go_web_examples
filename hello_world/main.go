package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("test")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Fprintln(w, r)
	})

	http.ListenAndServe(":10081", nil)
}

/*
Hello, you've requested: /
&{
	GET / 
	HTTP/1.1 1 1 
	map[
		Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*'/*;q=0.8,application/signed-exchange;v=b3;q=0.7] 
		Accept-Encoding:[gzip, deflate, br] 
		Accept-Language:[ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7] 
		Connection:[keep-alive] 
		Sec-Ch-Ua:["Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"] Sec-Ch-Ua-Mobile:[?0] 
		Sec-Ch-Ua-Platform:["macOS"] 
		Sec-Fetch-Dest:[document] 
		Sec-Fetch-Mode:[navigate] 
		Sec-Fetch-Site:[none] 
		Sec-Fetch-User:[?1] 
		Upgrade-Insecure-Requests:[1] 
		User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) 
		AppleWebKit/537.36 (KHTML, like Gecko) 
		Chrome/112.0.0.0 Safari/537.36]
	] {} <nil> 0 [] false localhost:10081 map[] map[] <nil> map[] [::1]:56133 / <nil> <nil> <nil> 0x14000090050}
*/