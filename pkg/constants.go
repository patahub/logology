package log

var timestamp_format = "2006-01-02 15:04:05.000000"

const (
	CONSOLE_ON  = true
	CONSOLE_OFF = false
)

const (
	ANONYMOUS_TRUE  = true
	ANONYMOUS_FALSE = false
)

const (
	TYPE_COMPLIANCE byte = iota
	TYPE_LOG
	TYPE_METRIC
	TYPE_TRANS_RATE
	TYPE_DEBUG
	TYPE_TRACE
	TYPE_ALL
)

const (
	SUBTYPE_METRIC_DURATION  = "DURATION"
	SUBTYPE_METRIC_TIMESTAMP = "TIMESTAMP"
	SUBTYPE_METRIC_COUNTER   = "COUNTER"
	SUBTYPE_METRIC_VALUE     = "VALUE"
)

const (
	SEVERITY_CRITICAL byte = iota
	SEVERITY_ERROR
	SEVERITY_WARNING
	SEVERITY_INFO
	SEVERITY_NORMAL
	SEVERITY_ALL
)

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

const (
	unit_transaction_rate_tps = "TPS"
	unit_transaction_rate_tpm = "TPM"
)

var types [7]string
var severities [6]string

func initConstants() {
	types = [7]string{"COMPLIANCE", "LOG", "METRIC", "T-RATE", "DEBUG", "TRACE", "ALL"}
	severities = [6]string{"CRITICAL", "ERROR", "WARNING", "INFO", "NORMAL", "ALL"}
}
