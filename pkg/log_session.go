package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type LogSession struct {
	sync.Mutex
	ID             string
	ServiceID      string
	MaxSeverity    byte
	MaxType        byte
	Anonymized     bool
	timeCreated    time.Time
	rx             chan logEvent // receiver for log events
	metricDuration mapMetricDuration
}

func NewLogSession(serviceID string, maxSeverity byte, maxType byte, anonymous bool) *LogSession {
	ls := new(LogSession)
	ls.ID = uuid.New().String()[:8] // just the first part of UUID is sufficient for uniqueness and avoid hyphens
	ls.ServiceID = serviceID
	ls.timeCreated = time.Now().UTC()
	ls.MaxSeverity = maxSeverity
	ls.MaxType = maxType
	ls.Anonymized = anonymous
	// init
	initConstants()
	ls.metricDuration.init()
	// init receiver channell
	ls.rx = make(chan logEvent, 100)
	go ls.receiver()
	return ls
}

// generic log function for all log events
// allows filtering logic and other initial decicions before sending event to receiver
func (ls *LogSession) log(le logEvent) {
	// ensure only the
	if le.logtype > ls.MaxType {
		return
	}
	if le.severity > ls.MaxSeverity {
		return
	}
	ls.rx <- le
}

// receiver for all log events
func (ls *LogSession) receiver() {
	tailMessage := ""
	for {
		le := <-ls.rx
		switch le.logtype {
		case TYPE_METRIC:
			// TODO: various things todo
			// make sure locking is correct all allong
			// delete entry from map once timing is complete/ended
			// write log entries
			// SLO framework , create other events in case under SLO (critical, etc)
			// define ways to collect average
			if le.subType == SUBTYPE_METRIC_DURATION {
				tailMessage = ls.metricDuration.set(&le)
			}
		}
		// write log entries
		fmt.Printf("%s, %s%s\n", le.HeaderString(), le.MessageString(), tailMessage)
	}
}
