package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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

func (con *JIRAConnector) GetProjectIssuesJSON(projectIdOrKey string) (dto.IssuesResponse, error) {
	baseURL := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&maxResults=0&expand=changelog", con.Cfg.Connector.JiraUrl, projectIdOrKey)

	initialResp, err := con.doIssuesRequest(baseURL)
	if err != nil {
		return dto.IssuesResponse{}, err
	}

	issuesChan := make(chan []dto.Issue)
	errorsChan := make(chan error)
	for i := 0; i < con.Cfg.Connector.GoroutinesCount; i++ {
		startAt := i * con.Cfg.Connector.MaxIssuesInRequest
		url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&startAt=%d&maxResults=%d&expand=changelog", con.Cfg.Connector.JiraUrl, projectIdOrKey, startAt, con.Cfg.Connector.MaxIssuesInRequest)
		fmt.Println(url)
		go func(url string) {
			resp, err := con.doIssuesRequest(url)
			if err != nil {
				errorsChan <- err
				return
			}
			issuesChan <- resp.Issues
		}(url)
	}

	doneGoroutines := 0
	for {
		select {
		case issues, ok := <-issuesChan:
			if ok {
				initialResp.Issues = append(initialResp.Issues, issues...)
				doneGoroutines++
			}
		case err := <-errorsChan:
			return dto.IssuesResponse{}, err
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout waiting for issues.")
			return dto.IssuesResponse{}, fmt.Errorf("timeout waiting for issues from JIRA API")
		}

		if doneGoroutines >= con.Cfg.Connector.GoroutinesCount {
			break
		}
	}

	return initialResp, nil
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
