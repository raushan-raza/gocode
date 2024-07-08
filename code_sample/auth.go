package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "https://apiv2.shiprocket.in/v1/external/auth/login"
	method := "POST"

	payload := strings.NewReader(`{
    "email": "shudhanshu.shekhar.test@shiprocket.com",
    "password": "Test@123"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
