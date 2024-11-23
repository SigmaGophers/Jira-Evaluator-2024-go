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
	Cfg        *config.Config
}

func NewJIRAConnector(cfg *config.Config) *JIRAConnector {
	connector := new(JIRAConnector)
	connector.HttpClient = &http.Client{}
	connector.Cfg = cfg
	return connector
}

func (con *JIRAConnector) GetProjectIssuesJSON(projectIdOrKey string) (dto.IssuesResponse, error) { // А что если projectIdOrKey неправильный? Добавить дефолтные значения
	url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&maxResults=%d&expand=changelog", con.Cfg.Connector.JiraUrl, projectIdOrKey, con.Cfg.Connector.MaxIssuesInRequest)
	resp, err := con.doIssuesRequest(url)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	return resp, nil
}

func (con *JIRAConnector) doIssuesRequest(url string) (dto.IssuesResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	response, err := con.HttpClient.Do(req)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return dto.IssuesResponse{}, fmt.Errorf("error: %s", response.Status)
	}

	return con.parseIssuesResponse(response.Body)
}

func (con *JIRAConnector) parseIssuesResponse(body io.ReadCloser) (dto.IssuesResponse, error) {
	var response dto.IssuesResponse

	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	return response, nil
}
