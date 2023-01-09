package log

import (
	"time"
)

type timeMetric struct {
	transactionID string
	tenantID      string
	serviceID     string
	scopeID       string
	startTime     time.Time
	measurement   time.Duration
}

// returns a uint64 representing the hash of the keys and the timeMetric
func NewTimeMetric(le *logEvent) *timeMetric {
	var tm = &timeMetric{
		transactionID: le.transactionID,
		tenantID:      le.tenantID,
		serviceID:     le.serviceID,
		scopeID:       le.scopeID,
		startTime:     time.Now().UTC(),
		measurement:   time.Duration(0),
	}
	return tm
}

func (tm *timeMetric) end(le *logEvent) {
	tm.measurement = time.Now().UTC().Sub(tm.startTime)
}

func (ls *LogSession) LogMetricDurationStart(transactionID string, tenantID string, scopeID string) {
	var le = logEvent{
		logtype:       TYPE_METRIC,
		subType:       SUBTYPE_METRIC_DURATION,
		severity:      SEVERITY_INFO,
		serviceID:     ls.ServiceID,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       "",
		metricKey:     "",
	}
	ls.log(le)
}

func (ls *LogSession) LogMetricDurationEnd(transactionID string, tenantID string, scopeID string, message string, metricKey string) {
	var le = logEvent{
		logtype:       TYPE_METRIC,
		subType:       SUBTYPE_METRIC_DURATION,
		severity:      SEVERITY_INFO,
		serviceID:     ls.ServiceID,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       message,
		metricKey:     metricKey,
	}
	ls.log(le)
}
