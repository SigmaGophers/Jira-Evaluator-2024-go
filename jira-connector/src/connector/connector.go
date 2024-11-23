package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jira-connector/src/config"
	"github.com/jira-connector/src/dto"
)

type JIRAConnector struct {
	HttpClient *http.Client
	Config     *config.Config
}

func NewJIRAConnector(cfg *config.Config) *JIRAConnector {
	connector := new(JIRAConnector)
	connector.HttpClient = &http.Client{}
	connector.Config = cfg
	return connector
}

func (connector *JIRAConnector) GetProjectIssuesJSON(projectIdOrKey string) ([]dto.Issue, error) { // А что если projectIdOrKey неправильный?
	url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&maxResults=%d&expand=changelog", connector.Config.Connector.JiraUrl, projectIdOrKey, connector.Config.Connector.IssuesPerRequest)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := connector.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", response.Status)
	}

	return connector.parseIssuesResponse(response.Body)
}

func (connector *JIRAConnector) parseIssuesResponse(body io.ReadCloser) ([]dto.Issue, error) {
	var searchResponse struct {
		Issues []dto.Issue `json:"issues"`
	}

	err := json.NewDecoder(body).Decode(&searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.Issues, nil
}
