package main

import (
	"fmt"
	"time"

	"github.com/jira-connector/src/config"
	"github.com/jira-connector/src/connector"
)

func main() {
	cfg := config.Config{}
	err := cfg.Parse()
	if err != nil {
		fmt.Print(err)
	}

	jiraConnector := connector.NewJIRAConnector(&cfg)
	start := time.Now()
	data, err := jiraConnector.GetProjectIssuesJSON("AAR")
	duration := time.Since(start)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(data.Issues[0].Fields.Project)
	fmt.Printf("\nTime: %v\n", duration) // Time: 746.696ms Time: 726.4737ms Time: 726.4737ms Time: 731.9331ms
}
