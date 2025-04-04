package todo

type Todo struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Done bool   `json:"done,omitempty"`
}
