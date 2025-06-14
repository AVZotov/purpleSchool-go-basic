package main

import (
	"encoding/json"
	"fmt"
	"http_cli/api"
	"http_cli/api/response"
	"http_cli/config"
	"io"
)

type PostData struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId string `json:"userId"`
}

func main() {
	postData := PostData{
		Title:  "Post Title1",
		Body:   "Post Body1",
		UserId: "12345",
	}
	data, _ := json.Marshal(postData)
	configs := config.NewEnvConfig()
	client := api.NewClient(configs)
	resp, err := client.Delete("6846078b8a456b7966ab05fc")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body := resp.Body
	defer body.Close()
	data, _ = io.ReadAll(body)
	var resStr response.Response
	err = json.Unmarshal(data, &resStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resStr.GetMessage())
	fmt.Println(resStr.GetId())
}
