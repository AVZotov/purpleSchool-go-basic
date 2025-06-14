package main

import (
	"encoding/json"
	"fmt"
	"http_cli/api"
	"http_cli/api/response"
	"http_cli/config"
	"http_cli/storage"
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
	ls := storage.NewBinList()
	configs := config.NewEnvConfig()
	client := api.NewClient(configs, ls)
	resp, err := client.Create("test1", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var createResponse response.CreateResponse
	data, _ = io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &createResponse)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	fmt.Println(createResponse.Metadata.CreatedAt)
	fmt.Println(createResponse.Metadata.Id)
	fmt.Println(createResponse.Metadata.Name)
	fmt.Println(createResponse.Metadata.Private)
}
