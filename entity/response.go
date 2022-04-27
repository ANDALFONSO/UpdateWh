package entity

type Response struct {
	Total     int         `json:"total"`
	ContextId string      `json:"context_id"`
	Documents []Documents `json:"documents"`
}

type Documents struct {
	Services []Services `json:"services"`
}

type Services struct {
	Id int `json:"id"`
}
