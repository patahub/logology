package log

import (
	"fmt"
	"hash/fnv"
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
	value         float64
	unit          string
}

func (le *logEvent) HeaderString() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", types[le.logtype], le.subType, severities[le.severity], le.tenantID, le.serviceID, le.transactionID, le.scopeID)
}

func (le *logEvent) MessageString() string {
	if le.logtype != TYPE_METRIC {
		return fmt.Sprintf("%s", le.message)
	}
	// METRIC
	return fmt.Sprintf("%s, %f, %s", le.message, le.value, le.unit)
}

func (le *logEvent) getHash() uint64 {
	h := fnv.New64a()
	h.Write([]byte(le.transactionID + le.tenantID + le.serviceID + le.scopeID))
	return h.Sum64()
}
