package dto

type IssuesResponse struct {
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Project     Project   `json:"project"`
	Creator     User      `json:"creator"`
	Summary     string    `json:"summary"`
	Changelog   Changelog `json:"changelog"`
	Description string    `json:"description"`
}

type Project struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Key         string `json:"key"`
}

type Changelog struct {
	MaxResults string    `json:"maxResults"`
	Total      string    `json:"total"`
	Histories  []History `json:"histories"`
}

type History struct {
	ID      string `json:"id"`
	Author  User   `json:"author"`
	Created string `json:"created"`
	Items   []Item `json:"items"`
}

type Item struct {
	Field      string `json:"field"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}
