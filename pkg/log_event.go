package log

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"time"
)

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

func (le *logEvent) anonymize() {
	h := fnv.New64a()
	h.Write([]byte(le.tenantID))
	le.tenantID = strconv.FormatUint(h.Sum64(), 32) // octal representation
	h.Reset()
}

func (le *logEvent) headerString() string {
	return fmt.Sprintf("%q, %q, %q, %q, %q, %q, %q, %q",
		time.Now().UTC().Format(timestamp_format),
		types[le.logtype],
		le.subType, severities[le.severity],
		le.tenantID,
		le.serviceID,
		le.transactionID,
		le.scopeID,
	)
}

func (le *logEvent) messageString() string {
	if le.logtype != TYPE_METRIC {
		return fmt.Sprintf("%q", le.message)
	}
	// METRIC
	return fmt.Sprintf("%q, %q, %f, %q, %q", le.message, le.metricKey, le.value, le.unit, le.fact)
}

func (le *logEvent) getHash() uint64 {
	h := fnv.New64a()
	h.Write([]byte(le.transactionID + le.tenantID + le.serviceID + le.scopeID))
	return h.Sum64()
}
