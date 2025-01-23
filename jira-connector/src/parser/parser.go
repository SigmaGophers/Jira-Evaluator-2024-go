package parser

import (
	"encoding/json"
	"io"

	"github.com/jira-connector/src/dto"
)

// Мб пока хватит и просто функцию?
// Может стоит будет добавить логирование или что то в этом духе
type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseIssuesResponse(body io.ReadCloser) (dto.IssuesResponse, error) {
	var response dto.IssuesResponse
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	return response, nil
}
