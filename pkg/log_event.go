package log

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"time"
)

// internal structure for all log events
type logEvent struct {
	logtype       byte
	subType       string // used to diferentiate for e.g. TYPE_METRIC could be duration (start/end) or increment, etc
	severity      byte
	tenantID      string
	serviceID     string
	transactionID string
	scopeID       string
	message       string
	metricKey     string
	value         float64
	unit          string
	fact          string // an observation/detail (such as start/end for metric duration)
}

// used to anonymize log events... should replace PPI info by a hash
func (le *logEvent) anonymize() {
	h := fnv.New64a()
	h.Write([]byte(le.tenantID))
	le.tenantID = strconv.FormatUint(h.Sum64(), 32) // octal representation
	h.Reset()
}

// generates the header part of the log message
func (le *logEvent) headerString(format string) string {
	return fmt.Sprintf("%q, %q, %q, %q, %q, %q, %q, %q",
		time.Now().UTC().Format(format),
		types[le.logtype],
		le.subType, severities[le.severity],
		le.tenantID,
		le.serviceID,
		le.transactionID,
		le.scopeID,
	)
}

// generates the body part of the message
func (le *logEvent) messageString() string {
	if le.logtype == TYPE_METRIC {
		return fmt.Sprintf("%q, %q, %f, %q, %q", le.message, le.metricKey, le.value, le.unit, le.fact)
	}
	if le.logtype == TYPE_TRANS_RATE {
		return fmt.Sprintf("%q, %q, %f, %q", le.message, le.metricKey, le.value, le.unit)
	}
	return fmt.Sprintf("%q", le.message)
}

// hash function that uniquely defines a combination of transaction, tenant, service and scope
// used to associate a program flow specific to a transaction/tenant such as for a duration type of log
func (le *logEvent) getHash() uint64 {
	h := fnv.New64a()
	h.Write([]byte(le.transactionID + le.tenantID + le.serviceID + le.scopeID))
	return h.Sum64()
}
