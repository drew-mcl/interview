package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

func isIndentedLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	return unicode.IsSpace(rune(line[0]))
}

func countErrorLogs(filename string) (totalERROR int, messages []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, nil, fmt.Errorf("error reading log file: %v", err)
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if isErrorLine(line) {
			totalERROR++

			level := ""
			if strings.Contains(line, "ERROR") {
				level = "ERROR"
			} else if strings.Contains(line, " ERR ") || strings.HasPrefix(line, "ERR ") || strings.HasSuffix(line, " ERR") {
				level = "ERR"
			} else {
				continue
			}

			message := ""
			if idx := strings.LastIndex(line, "-"); idx != -1 {
				message = strings.TrimSpace(line[idx+1:])
			} else {
				message = strings.TrimSpace(line)
			}

			stackTraceLines := []string{}
			for i+1 < len(lines) && isIndentedLine(lines[i+1]) {
				i++
				stackTraceLines = append(stackTraceLines, strings.TrimSpace(lines[i]))
			}

			formattedMessage := fmt.Sprintf("%s: %s", level, message)
			if len(stackTraceLines) > 0 {
				formattedMessage += "\n  Stack Trace:"
				for _, traceLine := range stackTraceLines {
					formattedMessage += "\n    " + traceLine
				}
			}
			messages = append(messages, formattedMessage)
		}
	}

	return totalERROR, messages, nil
}
