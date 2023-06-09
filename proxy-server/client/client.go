package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	// response, err := c.Get("http://localhost:1338/")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }
	response2, err := c.Get("http://localhost:1338/tv")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// defer response.Body.Close()
	defer response2.Body.Close()

	body, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		fmt.Println("Failed to Read Response body: ", err)
	}
	fmt.Printf("Body : %s\n", body)
}
