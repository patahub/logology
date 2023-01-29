package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// LogSession contains all specifications for the log session. Each program should have only one LogSession.
// Initialize it through the NewLogSession function.
type LogSession struct {
	sync.Mutex
	id                    string
	serviceID             string
	maxSeverity           byte
	maxType               byte
	anonymized            bool
	filePath              string // non mandatory
	consoleOutput         bool
	timeCreated           time.Time
	timeStampFormat       string
	rx                    chan logEvent // receiver for log events
	metricDuration        mapMetricDuration
	transactionRateEngine *transactionRateEngine
}

// Initializes a new LogSession. Notice that the log time stamp format defaults to
// "2006-01-02 15:04:05.000000" and it can be overridden by calling SetTimeStampFormat(string)
func NewLogSession(serviceID string, maxSeverity byte, maxType byte, anonymize bool, consoleOutput bool) *LogSession {
	ls := new(LogSession)
	ls.id = uuid.New().String()[:8] // just the first part of UUID is sufficient for uniqueness and avoid hyphens
	ls.serviceID = serviceID
	ls.timeCreated = time.Now().UTC()
	ls.maxSeverity = maxSeverity
	ls.maxType = maxType
	ls.anonymized = anonymize
	ls.consoleOutput = consoleOutput
	ls.timeStampFormat = "2006-01-02 15:04:05.000000"
	// init
	ls.metricDuration.init()
	// init receiver channell
	ls.rx = make(chan logEvent, 100)
	go ls.receiver()
	// init KPI / transaction rate
	ls.transactionRateEngine = newTransactionRate(ls)
	return ls
}

// Overrides the default timestamp format which is "2006-01-02 15:04:05.000000"
func (ls *LogSession) SetTimeStampFormat(format string) {
	ls.Lock()
	defer ls.Unlock()
	ls.timeStampFormat = format
}

// generic log function for all log events
// allows filtering logic and other initial decicions before sending event to receiver
func (ls *LogSession) log(le logEvent) {
	// ensure only the
	if le.logtype > ls.maxType {
		return
	}
	if le.severity > ls.maxSeverity {
		return
	}
	ls.rx <- le
}

// receiver for all log events
func (ls *LogSession) receiver() {
	for {
		le := <-ls.rx
		// anonymize log event if applicable
		if ls.anonymized {
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
		if ls.consoleOutput {
			fmt.Printf("%s, %s\n", le.headerString(ls.timeStampFormat), le.messageString())
		}
	}
}
