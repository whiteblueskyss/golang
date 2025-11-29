package main

import (
	"fmt"
	"io"
	"net/http"
)

const Myurl = "https://echo.free.beeceptor.com"

func httprequest() {
	fmt.Println("A web request")

	response, err := http.Get(Myurl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
