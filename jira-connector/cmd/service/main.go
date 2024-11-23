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

	fmt.Print(data[0].Fields.Project)
	fmt.Printf("\nTime: %v\n", duration) // Time: 710.7165ms Time: 751.1132ms Time: 781.1767ms Time: 717.7052ms
}
