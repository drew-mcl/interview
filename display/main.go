package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run log_counter.go <logfile>")
		return
	}

	logFile := os.Args[1]
	totalErrors, messages, err := countErrorLogs(logFile)
	if err != nil {
		fmt.Printf("Error processing log file: %v\n", err)
		return
	}

	fmt.Println("===== Error Log Counts =====")
	fmt.Printf("Total ERROR logs: %d\n", totalErrors)
	fmt.Println()
	fmt.Println("===== Error Messages =====")
	for _, msg := range messages {
		fmt.Println(msg)
	}
}

func isErrorLine(line string) bool {
	return strings.Contains(line, "ERROR") || strings.Contains(line, "ERR")
}

func countErrorLogs(filename string) (totalERROR int, messages []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isErrorLine(line) {
			totalERROR++

			str := strings.Split(line, "-")
			fmt.Printf(str[1])

			message := ""
			if idx := strings.LastIndex(line, "-"); idx != -1 {
				message = strings.TrimSpace(line[idx+1:])
			}

			formattedMessage := fmt.Sprintf("ERROR: ", message)
			messages = append(messages, formattedMessage)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, fmt.Errorf("error reading log file: %v", err)
	}

	return totalERROR, messages, nil
}
