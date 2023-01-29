package log

import (
	"sync"
	"time"
)

// structure that hols the duration data the duration is associated with a unique hash that identifies the flow details,
// such as the combination of a tenant, transaction, service and scope
type mapMetricDuration struct {
	sync.Mutex
	mapMetricTime map[uint64]timeMetric
}

func (m *mapMetricDuration) init() {
	m.mapMetricTime = make(map[uint64]timeMetric)
}

// manages the mutex and mapMetricDuraction
// returns a tail for the log message (START, END)
func (m *mapMetricDuration) set(le *logEvent) {
	m.Lock()
	defer m.Unlock()
	hash := le.getHash()
	timeMetric, ok := m.mapMetricTime[hash]
	if ok {
		timeMetric.end(le)
		// calculate duration in milliseconds (float64)
		le.value = float64(timeMetric.measurement) / float64(time.Millisecond)
		le.unit = UNIT_MILLISECONDS
		// free map entry
		delete(m.mapMetricTime, hash)
		le.fact = "START"
		return
	}
	// new metric
	NewTimeMetric(le)
	m.mapMetricTime[hash] = *NewTimeMetric(le)
	le.fact = "BEGIN"
}
