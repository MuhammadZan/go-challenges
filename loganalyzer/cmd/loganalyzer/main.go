package main

import (
	"bufio"
	"flag"
	"log"
	"loganalyzer/internal/parser"
	"loganalyzer/internal/presenter"
	"os"
)


func main() {
	fileAdd := flag.String("file", "server.log", "File to analyze logs")
	flag.Parse()
	file, err := os.OpenFile(*fileAdd, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	analysis,err:= parser.ProccessFile(scanner)
	if err != nil {
		log.Fatal(err)
	}
	presenter.PresentLogsAnalysis(analysis)
}
