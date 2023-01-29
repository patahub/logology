package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type LogSession struct {
	sync.Mutex
	ID                    string
	ServiceID             string
	MaxSeverity           byte
	MaxType               byte
	Anonymized            bool
	FilePath              string // non mandatory
	ConsoleOutput         bool
	timeCreated           time.Time
	rx                    chan logEvent // receiver for log events
	metricDuration        mapMetricDuration
	transactionRateEngine *transactionRateEngine
}

func NewLogSession(serviceID string, maxSeverity byte, maxType byte, anonymize bool, consoleOutput bool) *LogSession {
	ls := new(LogSession)
	ls.ID = uuid.New().String()[:8] // just the first part of UUID is sufficient for uniqueness and avoid hyphens
	ls.ServiceID = serviceID
	ls.timeCreated = time.Now().UTC()
	ls.MaxSeverity = maxSeverity
	ls.MaxType = maxType
	ls.Anonymized = anonymize
	ls.ConsoleOutput = consoleOutput
	// init
	initConstants()
	ls.metricDuration.init()
	// init receiver channell
	ls.rx = make(chan logEvent, 100)
	go ls.receiver()
	// init KPI / transaction rate
	ls.transactionRateEngine = NewTransactionRate(ls)
	return ls
}

func (ls *LogSession) SetTimeStampFormat(format string) {
	timestamp_format = format
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
	for {
		le := <-ls.rx
		// anonymize log event if applicable
		if ls.Anonymized {
			le.anonymize()
		}
		switch le.logtype {
		case TYPE_METRIC:
			if le.subType == SUBTYPE_METRIC_DURATION {
				// this handles both duration start and end events
				ls.metricDuration.set(&le)
				break
			}
		}
		// write log entries
		if ls.ConsoleOutput {
			fmt.Printf("%s, %s\n", le.headerString(), le.messageString())
		}
	}
}
