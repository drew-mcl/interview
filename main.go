package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// LogLevel defines a log level with an associated weight
type LogLevel struct {
	Name   string
	Weight int
}

// LogEntry represents a single log entry with relevant fields.
type LogEntry struct {
	Timestamp  string
	Thread     string
	Level      string
	Component  string
	Message    string
	StackTrace []string
}

// generateTradingLogFile generates a trading application log file with normalized timestamps.
func generateTradingLogFile(filename string, numLines int) {
	// Define log levels with weights
	logLevels := []LogLevel{
		{"INFO", 40},  // 40% probability
		{"DEBUG", 30}, // 30% probability
		{"WARN", 20},  // 20% probability
		{"ERROR", 7},  // 7% probability
		{"ERR", 3},    // 3% probability
	}

	components := []string{
		"com.tradingapp.order.OrderService",
		"com.tradingapp.marketdata.MarketDataHandler",
		"com.tradingapp.execution.ExecutionEngine",
		"com.tradingapp.utils.DatabaseConnector",
		"com.tradingapp.auth.AuthenticationService",
		"com.tradingapp.alert.AlertManager",
		"com.tradingapp.config.ConfigLoader",
		"com.tradingapp.cache.CacheManager",
		"com.tradingapp.network.NetworkManager",
		"com.tradingapp.analytics.AnalyticsProcessor",
	}

	messages := map[string][]string{
		"INFO": {
			"Order received for execution",
			"Market data update processed",
			"User logged in successfully",
			"Heartbeat received from exchange",
			"Configuration loaded successfully",
			"Cache initialized",
			"New trading session started",
			"Historical data loaded",
			"Connection established with exchange",
			"Order book updated",
			"Price feed subscribed",
			"Trade confirmation received",
			"Balance updated",
			"User preferences saved",
			"Session token refreshed",
			"Trade settled successfully",
			"Position updated in portfolio",
			"Market summary generated",
			"Risk parameters evaluated",
			"Trade history archived",
			"Portfolio rebalanced",
			"Compliance check passed",
			"Liquidity provision optimized",
			"Asset allocation adjusted",
			"Margin requirements updated",
			"Trade limit reset",
			"Order status synchronized",
			"External API call successful",
			"Scheduled maintenance completed",
			"User notification sent",
			"Exchange rate updated",
		},
		"DEBUG": {
			"Order validation passed",
			"Processing trade execution",
			"Fetching market data snapshot",
			"Cache refreshed",
			"Session token generated",
			"Debugging authentication flow",
			"Calculating trade metrics",
			"Parsing configuration file",
			"Serializing order object",
			"Deserializing market data",
			"Thread started",
			"Memory usage checked",
			"Latency measurement taken",
			"Message queue size: 42",
			"Temporary file created",
			"Garbage collection initiated",
			"Lock acquired on resource",
			"Thread synchronization complete",
			"Retrying failed operation",
			"Loading user preferences",
			"Updating cache entries",
			"Monitoring system health",
			"Analyzing trade patterns",
			"Executing background job",
			"Refreshing API tokens",
			"Validating input parameters",
			"Encrypting sensitive data",
			"Decompressing data stream",
			"Optimizing query performance",
			"Profiling memory usage",
			"Tracking user session",
			"Resolving DNS query",
			"Handling socket connection",
			"Dispatching event handler",
			"Updating UI components",
			"Parsing JSON response",
			"Compressing log data",
			"Managing thread pool",
			"Synchronizing database state",
			"Processing batch job",
			"Validating transaction integrity",
			"Initializing module dependencies",
			"Capturing screenshot for debugging",
			"Loading external libraries",
			"Setting up test environment",
			"Executing unit tests",
			"Aggregating metrics data",
			"Rotating log files",
			"Scaling application instances",
			"Balancing load across servers",
			"Configuring network settings",
			"Establishing secure connection",
			"Decrypting received data",
			"Serializing response payload",
			"Deserializing request payload",
		},
		"WARN": {
			"Latency spike detected",
			"Market data delayed",
			"Order quantity exceeds threshold",
			"Price deviation detected",
			"Partial fill received",
			"High memory usage",
			"Disk space running low",
			"Unrecognized message type",
			"Retrying connection to exchange",
			"Deprecated API used",
			"Configuration parameter missing, using default",
			"User session about to expire",
			"Thread pool exhausted",
			"Cache miss occurred",
			"Failed to send heartbeat",
			"Slow response from database",
			"Potential deadlock detected",
			"Unexpected user input received",
			"Service latency above threshold",
			"Failed to retrieve user data",
			"Insufficient permissions for operation",
			"High CPU usage detected",
			"Unusual trading volume observed",
			"Failed to write to log file",
			"Memory leak detected in module",
			"External service response delayed",
			"Invalid trade signal received",
			"Connection timeout with broker",
			"Data inconsistency found",
			"Failed to acquire necessary locks",
			"Unsupported protocol version",
			"Resource utilization exceeds limits",
		},
		"ERROR": {
			"Failed to connect to database",
			"NullPointerException in OrderService",
			"ArrayIndexOutOfBoundsException in MarketDataHandler",
			"User authentication failed",
			"Order rejection from exchange",
			"Timeout while waiting for response",
			"Failed to parse configuration",
			"Transaction rollback due to error",
			"Network unreachable",
			"IOException in NetworkManager",
			"Failed to load security keys",
			"Order validation error",
			"Unable to update balance",
			"Failed to initialize cache",
			"Data corruption detected",
			"Service unavailable",
			"Failed to serialize object",
			"Unhandled exception in thread",
			"Error committing transaction",
			"Failed to start listener",
			"Disk read failure",
			"Memory allocation failed",
			"Unexpected shutdown of service",
			"Permission denied when accessing resource",
			"Failed to acquire database lock",
			"Corrupted log file detected",
			"Dependency resolution failed",
			"Failed to allocate buffer memory",
			"Error writing to output stream",
			"Unhandled error in request handler",
			"Failed to terminate process",
			"Error initializing network interface",
		},
		"ERR": { // Added messages for "ERR" level
			"Failed to process order",
			"Error retrieving market data",
			"Unexpected null value encountered",
			"Error in trade execution module",
			"Failed to authenticate user",
			"Error writing to database",
			"Invalid configuration detected",
			"Error in network communication",
			"Failed to start service",
			"Error during shutdown process",
			"Resource not found",
			"Dependency injection failed",
			"Failed to allocate memory",
			"Invalid user input format",
			"Error loading external resource",
			"Failed to initialize module",
			"Configuration validation failed",
			"Error parsing user request",
			"Failed to retrieve asset details",
			"Error updating trade status",
			"Service crash detected",
			"Failed to bind to port",
			"Error in data serialization",
			"Failed to delete temporary files",
			"Invalid response from external API",
			"Error handling client request",
			"Failed to restart service",
			"Error during data migration",
			"Failed to release resources",
			"Error in authentication middleware",
			"Failed to log user activity",
			"Error processing webhook",
		},
	}

	timestampFormat := "2006-01-02 15:04:05.000" // Normalized timestamp format
	separators := []string{" ", "\t"}
	threadNames := []string{
		"main",
		"OrderThread-1",
		"ExecThread-2",
		"MarketDataThread-3",
		"AuthThread-4",
		"CacheThread-5",
		"AlertThread-6",
		"AnalyticsThread-7",
		"NetworkThread-8",
		"DBThread-9",
		"TradeProcessor-10",
		"RiskManager-11",
		"DataIngestion-12",
		"NotificationService-13",
		"ReportGenerator-14",
		"BackupThread-15",
	}

	stackTraceVariations := []int{3, 4, 5} // Different stack trace lengths

	currentTime := time.Now()

	// Open the file for writing
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())

	// Define after how many lines the config dump should appear
	configDumpLine := 10
	configDumpInserted := false

	for i := 0; i < numLines; i++ {
		// Random time increment
		delta := time.Duration(rand.Intn(500)+1) * time.Millisecond
		currentTime = currentTime.Add(delta)

		// Generate normalized timestamp
		timestamp := currentTime.Format(timestampFormat)

		// Select a log level based on weights
		level := selectLogLevel(logLevels)

		// Choose a component/class name
		component := components[rand.Intn(len(components))]

		// Choose a thread name
		thread := threadNames[rand.Intn(len(threadNames))]

		// Choose a separator
		separator := separators[rand.Intn(len(separators))]

		// Choose a message based on the level
		messageOptions := messages[level]
		message := messageOptions[rand.Intn(len(messageOptions))]

		// Introduce ' ... ' pattern with a certain probability for ERR level
		includeErrPattern := false
		if level == "ERR" {
			// 50% chance to include ' ... ' instead of component
			if rand.Float64() < 0.5 {
				includeErrPattern = true
			}
		}

		var line string

		// Insert config dump once after initial log lines
		if !configDumpInserted && i == configDumpLine {
			configDump := generateConfigDump()
			line = fmt.Sprintf("%s\n", configDump)
			configDumpInserted = true
		} else {
			// Determine if a stack trace should be included
			includeStackTrace := false
			if level == "ERROR" || level == "ERR" {
				// Reduce stack trace frequency to 15%
				if rand.Float64() < 0.15 {
					includeStackTrace = true
				}
			}

			if includeErrPattern {
				// Format with ' ... ' instead of component
				line = fmt.Sprintf("%s%s[%s] %s ... - %s\n", timestamp, separator, thread, level, message)
			} else {
				// Occasionally omit the component for non-ERR and non-ERROR levels (e.g., 5% chance)
				omitComponent := false
				if level != "ERR" && level != "ERROR" {
					if rand.Float64() < 0.05 {
						omitComponent = true
					}
				}

				if omitComponent {
					line = fmt.Sprintf("%s%s[%s] %s - %s\n", timestamp, separator, thread, level, message)
				} else {
					// Normal log line
					line = fmt.Sprintf("%s%s[%s] %s %s - %s\n", timestamp, separator, thread, level, component, message)
				}
			}

			// Append stack trace if needed
			if includeStackTrace {
				// Vary the size of the stack trace
				numTraceLines := stackTraceVariations[rand.Intn(len(stackTraceVariations))]
				stackTrace := generateStackTrace(component, numTraceLines)
				line += stackTrace + "\n"
			}
		}

		// Write the line to the file
		_, err := f.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// selectLogLevel selects a log level based on assigned weights
func selectLogLevel(logLevels []LogLevel) string {
	totalWeight := 0
	for _, level := range logLevels {
		totalWeight += level.Weight
	}

	randNum := rand.Intn(totalWeight) + 1
	cumulativeWeight := 0
	for _, level := range logLevels {
		cumulativeWeight += level.Weight
		if randNum <= cumulativeWeight {
			return level.Name
		}
	}

	// Fallback (shouldn't occur if weights are correct)
	return "INFO"
}

// generateStackTrace generates a stack trace for a given component with a specified number of lines
func generateStackTrace(component string, numLines int) string {
	baseStackTrace := []string{
		fmt.Sprintf("	at %s.methodA(%s.java:56)", component, getSimpleComponentName(component)),
		fmt.Sprintf("	at %s.methodB(%s.java:78)", component, getSimpleComponentName(component)),
		"	at com.tradingapp.utils.HelperClass.methodC(HelperClass.java:102)",
		"	at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)",
		"	at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)",
		"	at java.base/java.lang.Thread.run(Thread.java:834)",
	}

	// Shuffle the base stack trace to add more variations
	rand.Shuffle(len(baseStackTrace), func(i, j int) {
		baseStackTrace[i], baseStackTrace[j] = baseStackTrace[j], baseStackTrace[i]
	})

	// Select the desired number of stack trace lines
	if numLines > len(baseStackTrace) {
		numLines = len(baseStackTrace)
	}
	selectedTrace := baseStackTrace[:numLines]
	return strings.Join(selectedTrace, "\n")
}

// getSimpleComponentName extracts the simple name of the component from its full package path
func getSimpleComponentName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return fullName
}

// generateConfigDump generates a configuration dump string
func generateConfigDump() string {
	config := map[string]string{
		"max_connections":      "100",
		"timeout_seconds":      "30",
		"enable_logging":       "true",
		"log_level":            "DEBUG",
		"database_url":         "jdbc:mysql://localhost:3306/tradingdb",
		"cache_size":           "1024",
		"retry_attempts":       "5",
		"api_key":              "abcd1234efgh5678",
		"secret_key":           "wxyz9876tsrq5432",
		"exchange_endpoint":    "wss://exchange.example.com/socket",
		"allowed_ip_addresses": "192.168.1.1,192.168.1.2",
		"feature_flag_new_ui":  "false",
		"maintenance_mode":     "false",
		"backup_schedule":      "02:00 AM daily",
		"security_protocol":    "TLS1.2",
		"session_timeout":      "45",
		"data_retention_days":  "365",
		"max_trade_volume":     "1000000",
		"min_order_size":       "10",
		"currency_supported":   "USD, EUR, GBP, JPY",
		"notification_emails":  "admin@tradingapp.com,support@tradingapp.com",
	}

	var builder strings.Builder
	builder.WriteString("Current Configuration:\n")
	for k, v := range config {
		builder.WriteString(fmt.Sprintf("%s = %s\n", k, v))
	}
	return builder.String()
}

func main() {
	generateTradingLogFile("trading_application.log", 1000)
}
