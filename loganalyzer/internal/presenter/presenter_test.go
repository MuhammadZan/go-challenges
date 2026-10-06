package presenter

import (
	"testing"
	"loganalyzer/internal/analyzer"
)

func TestPresentData(t *testing.T){
	logData := &LogsAnalysis{
		TotalLogs: 10,
		MailformedLogs: []string{"invalid"},
		Analysis: &analyzer.Analysis{
			TotalResponseSize: 1024,
			LogsByMethod: []analyzer.RequestByMethod{
				{Method: "GET", Count: 5},
			},
			LogsByCode: []analyzer.RequestByCode{
				{Code: 200, Count: 5},
			},
			LogsByUrl: []analyzer.RequestByUrl{
				{Url: "/index.html", Count: 5},
			},
			LogsByIP: []analyzer.RequestByIP{
				{IP: "192.168.1.1", Count: 5},
			},
		},
	}
	PresentLogsAnalysis(logData)
}