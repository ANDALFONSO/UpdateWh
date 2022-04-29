package entity

type ResponseWh struct {
	Values []Values `json:"values"`
}

type Values struct {
	Id        string `json:"id"`
	Monday    Day    `json:"monday"`
	Tuesday   Day    `json:"tuesday"`
	Wednesday Day    `json:"wednesday"`
	Thursday  Day    `json:"thursday"`
	Friday    Day    `json:"friday"`
	Saturday  Day    `json:"saturday"`
	Sunday    Day    `json:"sunday"`
}
