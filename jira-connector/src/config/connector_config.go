package config

type ConnectorConfig struct {
	HttpPort           int    `yaml:"http_port"`
	JiraUrl            string `yaml:"jira_url"`
	MaxIssuesInRequest int    `yaml:"issues_per_request"`
	GoroutinesCount    int    `yaml:"number_of_goroutines"`
	MaxRetryWait       int    `yaml:"max_retry_wait"`
	InitialRetryWait   int    `yaml:"initial_retry_wait"`
}
