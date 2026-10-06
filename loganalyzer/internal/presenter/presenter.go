package presenter

import (
	"fmt"
	"loganalyzer/internal/analyzer"
)
type LogsAnalysis struct {
	TotalLogs int
	MailformedLogs []string
	Analysis *analyzer.Analysis
}

func PresentLogsAnalysis(logAnalysis *LogsAnalysis){
	fmt.Println("Total Logs Analyzed:",logAnalysis.TotalLogs)
	if logAnalysis.TotalLogs > 0 {
		fmt.Println("Average response size:",logAnalysis.Analysis.TotalResponseSize/logAnalysis.TotalLogs)
	}else{
		fmt.Println("Average response size:", 0)

	}
	fmt.Println("\nTop requests by method:")
	for _, l := range logAnalysis.Analysis.LogsByMethod {
		fmt.Println(l.Method,l.Count)
	}
	fmt.Println("\nTop 10 requests by status code:")
	if len(logAnalysis.Analysis.LogsByCode) > 10 {
		for _, l := range logAnalysis.Analysis.LogsByCode[:10] {
			fmt.Println(l.Code,l.Count)
		}
	}else{
		for _, l := range logAnalysis.Analysis.LogsByCode {
			fmt.Println(l.Code,l.Count)
		}
	}
	fmt.Println("\nTop 10 requests by URL:")
	if len(logAnalysis.Analysis.LogsByUrl) > 10 {
		for _, l := range logAnalysis.Analysis.LogsByUrl[:10] {
			fmt.Println(l.Url,l.Count)
		}
	}else{
		for _, l := range logAnalysis.Analysis.LogsByUrl {
			fmt.Println(l.Url,l.Count)
		}
	}
	fmt.Println("\nTop 10 requests by IP address:")
	if len(logAnalysis.Analysis.LogsByIP) > 10 {
		for _, l := range logAnalysis.Analysis.LogsByIP[:10] {
			fmt.Println(l.IP,l.Count)
		}
	}else{
		for _, l := range logAnalysis.Analysis.LogsByIP {
			fmt.Println(l.IP,l.Count)
		}
	}
	if len(logAnalysis.MailformedLogs) > 0 {
		fmt.Println("\nMailformed Logs:")
		for _, l := range logAnalysis.MailformedLogs {
			fmt.Println(l)
		}
	}
}