package log

import (
	"sync"
	"time"
)

type mapMetricDuration struct {
	sync.Mutex
	mapMetricTime map[uint64]timeMetric
}

// manages the mutex and mapMetricDuraction
// returns a tail for the log message (START, END)
func (m *mapMetricDuration) Set(key uint64, le *logEvent) string {
	m.Lock()
	defer m.Unlock()
	hash := le.getHash()
	timeMetric, ok := m.mapMetricTime[hash]
	if ok {
		timeMetric.end(le)
		// calculate duration in milliseconds (float64)
		le.value = float64(timeMetric.measurement) / float64(time.Millisecond)
		le.unit = UNIT_MILLISECONDS
		return ", END"
	}
	// new metric
	NewTimeMetric(le)
	m.mapMetricTime[hash] = *NewTimeMetric(le)
	return ", START"
}
