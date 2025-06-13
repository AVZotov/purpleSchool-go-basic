package api

const (
	root        = "https://api.jsonbin.io/v3"
	rout        = "/b"
	contentType = "application/json"
)

type Configs interface {
	GetMasterKey() string
}

type Api struct {
	Key string
}

func NewApi(configs Configs) *Api {
	return &Api{
		Key: configs.GetMasterKey(),
	}
}

func (api *Api) Create() {

}

func (api *Api) Read() {

}

func (api *Api) Update() {

}

func (api *Api) Delete() {

}
