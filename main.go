package main

import (
	"io"
	"fmt"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/users"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(-1)
	}
	fmt.Println("Response:", resp)
	fmt.Println("Response.body:", resp.Body)

	// respBs := make([]byte, 32 * 1024)
	// resp.Body.Read(respBs)
	// fmt.Println(string(respBs))
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}