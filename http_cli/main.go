package main

import (
	"encoding/json"
	"fmt"
	"http_cli/api"
	"http_cli/config"
	"http_cli/storage"
	"io"
	"time"
)

type Metadata struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
}

type CreateResponse struct {
	Metadata Metadata `json:"metadata"`
}

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
	response, err := client.Create("test1", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//m := map[string]interface{}{
	//	"metadata": map[string]any{"createdAt": "2025-06-13T23:45:40.936Z", "id": "684cb8248a456b7966adbdab", "name": "test1", "private": true},
	//	"record":   map[string]any{"body": "Post Body1", "title": "Post Title1", "userId": "12345"},
	//}
	//fmt.Println(m)
	var createResponse CreateResponse
	data, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(data, &createResponse)
	if err != nil {
		fmt.Println(err.Error())
	}
	response.Body.Close()

	fmt.Println(createResponse.Metadata.CreatedAt)
	fmt.Println(createResponse.Metadata.Id)
	fmt.Println(createResponse.Metadata.Name)
	fmt.Println(createResponse.Metadata.Private)
}
