package log

// Defines if logging should be bound to the console output or not.
const (
	CONSOLE_ON  = true
	CONSOLE_OFF = false
)

// Defines if anonymization is enabled
const (
	ANONYMOUS_TRUE  = true
	ANONYMOUS_FALSE = false
)

// Log types
const (
	TYPE_COMPLIANCE byte = iota
	TYPE_LOG
	TYPE_METRIC
	TYPE_TRANS_RATE
	TYPE_DEBUG
	TYPE_TRACE
	TYPE_ALL
)

// Log sub-types
const (
	SUBTYPE_METRIC_DURATION  = "DURATION"
	SUBTYPE_METRIC_TIMESTAMP = "TIMESTAMP"
	SUBTYPE_METRIC_COUNTER   = "COUNTER"
	SUBTYPE_METRIC_VALUE     = "VALUE"
)

// Severity levels
const (
	SEVERITY_CRITICAL byte = iota
	SEVERITY_ERROR
	SEVERITY_WARNING
	SEVERITY_INFO
	SEVERITY_NORMAL
	SEVERITY_ALL
)

// Time units
const (
	UNIT_NANOSECONDS  = "ns"
	UNIT_MICROSECONDS = "us"
	UNIT_MILLISECONDS = "ms"
	UNIT_SECONDS      = "s"
	UNIT_MINUTES      = "m"
	UNIT_HOURS        = "h"
	UNIT_DAYS         = "d"
	UNIT_YEARS        = "y"
)

// transaction rate constants
const (
	unit_transaction_rate_tps = "TPS"
	unit_transaction_rate_tpm = "TPM"
)

// global types and severity for readability of messages
var types = [7]string{"COMPLIANCE", "LOG", "METRIC", "T-RATE", "DEBUG", "TRACE", "ALL"}
var severities = [6]string{"CRITICAL", "ERROR", "WARNING", "INFO", "NORMAL", "ALL"}
