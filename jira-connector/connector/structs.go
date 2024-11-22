package connector

type Project struct {
	Self string `json:"self"` // А надо сохранять это поле?
	ID   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}
