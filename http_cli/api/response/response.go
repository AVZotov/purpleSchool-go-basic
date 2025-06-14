package response

import "time"

type Metadata struct {
	CreatedAt       time.Time `json:"createdAt"`
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Private         bool      `json:"private"`
	VersionsDeleted int       `json:"versionsDeleted"`
}

type Response struct {
	Metadata Metadata `json:"metadata"`
	Message  string   `json:"message"`
}

func (r *Response) GetMessage() string {
	return r.Message
}
func (r *Response) GetId() string {
	return r.Metadata.Id
}
