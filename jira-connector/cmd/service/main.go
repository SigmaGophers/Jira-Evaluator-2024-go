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
	for _, issue := range data.Issues {
		fmt.Printf("Issue ID: %s, Key: %s\n", issue.ID, issue.Key)
	}

	fmt.Print(data.IssuesCount, len(data.Issues))
	fmt.Printf("\nTime: %v\n", duration) // Time: 2.7817425s Time: 2.7046177s Time: 3.0152594s Time: 2.645088s
}
