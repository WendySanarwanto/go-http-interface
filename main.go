package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/users"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(-1)
	}
	fmt.Println("Response:", resp)
	fmt.Println("Response.body:", resp.Body)
	respBs := make([]byte, 99999)
	resp.Body.Read(respBs)
	fmt.Println(string(respBs))
}