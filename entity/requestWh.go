package entity

type RequestWh struct {
	Id        string `json:"id"`
	Monday    Day    `json:"monday"`
	Tuesday   Day    `json:"tuesday"`
	Wednesday Day    `json:"wednesday"`
	Thursday  Day    `json:"thursday"`
	Friday    Day    `json:"friday"`
	Saturday  Day    `json:"saturday"`
	Sunday    Day    `json:"sunday"`
}

type Day struct {
	Ranges []string `json:"ranges"`
}
