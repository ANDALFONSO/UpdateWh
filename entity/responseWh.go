package entity

type ResponseWh struct {
	Values []Values `json:"values"`
}

type Values struct {
	Id string `json:"id"`
}
