package parser

import (
	"bufio"
	"fmt"
	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/presenter"
	"strconv"
	"strings"
)


func ProccessFile(scanner *bufio.Scanner) (*presenter.LogsAnalysis, error) {
	az := analyzer.NewAnalyzer()
	logCount := 0
	mailformedLogs := []string{}

	for scanner.Scan() {
		logline := scanner.Text()
		logEntery, err := cleanAndExtractFields(logline)
		if err != nil {
			mailformedLogs = append(mailformedLogs, logline)
		}
		if logEntery == nil {
			mailformedLogs = append(mailformedLogs, logline)
		} else {
			logCount++
			az.Process(logEntery)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log file: ", err)
		return &presenter.LogsAnalysis{}, err
	}
	az.Sort()
	logAnalysis := presenter.LogsAnalysis{
		TotalLogs:      logCount,
		MailformedLogs: mailformedLogs,
		Analysis:       az,
	}
	
	return &logAnalysis, nil
}

func cleanAndExtractFields(logline string) (*analyzer.LogFields, error) {

	// Clean Log extraction will work assuming a fixed log format
	// Input Log Line: 192.168.1.105 - - [01/Jan/2026:08:02:48 +0000] "POST /api/v1/products HTTP/1.1" 200 1289
	cleanedLogLine := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(logline, "- -", ""), "[", ""), "]", ""), "\"", "")
	fields := strings.Fields(cleanedLogLine)
	if len(fields) >= 8 {
		statusCode, err := strconv.Atoi(fields[6])
		if err != nil {

			return nil, fmt.Errorf("Mailformed Log line (invalid status code): %s, err: %s", logline, err)
		}
		responseSize, err := strconv.Atoi(fields[7])
		if err != nil {

			return nil, fmt.Errorf("Mailformed Log line (invalid response size): %s, err: %s", logline, err)
		}
		validIP := validateIP(fields[0])
		validStatusCode := validateStatusCode(statusCode)
		if !validIP || !validStatusCode {
			return nil, fmt.Errorf("Mailformed Log line (invalid IP or status code): %s, err: %s", logline, err)
		}
		return &analyzer.LogFields{
			IpAddress:    fields[0],
			TimeStamp:    fields[1],
			Method:       fields[3],
			URL:          fields[4],
			StatusCode:   statusCode,
			ResponseSize: responseSize,
			HTTPprotocol: fields[5],
		}, nil
	}
	return nil, fmt.Errorf("Mailformed Log line (missing some fields): %s", logline)
}

func validateIP(ip string) bool {
	ipAddressFields := strings.Split(ip, ".")
	if len(ipAddressFields) != 4 {
		return false
	}
	for _, field := range ipAddressFields {
		octet, err := strconv.Atoi(field)
		if err != nil || octet < 0 || octet > 255 {
			return false
		}
	}
	return true
}

func validateStatusCode(statusCode int) bool {
	if statusCode >= 100 && statusCode <= 599 {
		return true
	}
	return false
}
