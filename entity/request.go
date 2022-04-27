package entity

type Request struct {
	Type      string `json:"type"`
	Query     Query  `json:"query"`
	Size      int    `json:"size"`
	ContextId string `json:"context_id"`
}

type Query struct {
	AnyEquals []AnyEquals `json:"any_equals"`
	Equals    []Equals    `json:"equals"`
}

type AnyEquals struct {
	Path   string   `json:"path"`
	Values []string `json:"values"`
}

type Equals struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}
