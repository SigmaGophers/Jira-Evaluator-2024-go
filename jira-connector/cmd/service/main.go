package main

import (
	"fmt"

	"github.com/jira-connector/config"
	"github.com/jira-connector/src/connector"
)

func main() {
	cfg := config.Config{}
	err := cfg.Parse()
	if err != nil {
		fmt.Print(err)
	}

	jiraConnector := connector.NewJIRAConnector(&cfg)
	data, err := jiraConnector.GetProjectIssuesJSON("AAR")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(data[0].Fields.Project)
}
