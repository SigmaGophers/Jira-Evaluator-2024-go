package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jira-connector/config"
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

func (connector *JIRAConnector) GetProjectDataJSON(idOrKey string) (Project, error) {
	url := fmt.Sprintf("%s/rest/api/2/project/%s", connector.Config.Connector.JiraUrl, idOrKey)
	req, err := http.NewRequest("GET", url, nil) // example: https://issues.apache.org/jira/rest/api/2/project/AAR
	if err != nil {
		return Project{}, err
	}

	response, err := connector.HttpClient.Do(req)
	if err != nil {
		return Project{}, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Project{}, fmt.Errorf("error: %s", response.Status)
	}

	return connector.parseProjectResponse(response.Body)
}

func (connector *JIRAConnector) parseProjectResponse(body io.ReadCloser) (Project, error) {
	var project Project
	err := json.NewDecoder(body).Decode(&project)
	if err != nil {
		return Project{}, err
	}
	return project, nil
}
