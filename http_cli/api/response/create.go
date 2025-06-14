package response

import "time"

type Metadata struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
}

type CreateResponse struct {
	Metadata Metadata `json:"metadata"`
}
