package log

import (
	"sync"
	"time"
)

type transactionRateEngine struct {
	sync.Mutex
	logSession    *LogSession
	mapCounterTPS map[string]uint32
	mapCounterTPM map[string]uint32
	tickerTPS     time.Ticker
	tickerTPM     time.Ticker
}

func NewTransactionRate(logSession *LogSession) *transactionRateEngine {
	var tre = &transactionRateEngine{}
	tre.init(logSession)
	return tre
}

func (tre *transactionRateEngine) init(logSession *LogSession) {
	tre.logSession = logSession
	// initialize maps
	tre.mapCounterTPS = make(map[string]uint32)
	tre.mapCounterTPM = make(map[string]uint32)
	// initialize tickers
	tre.tickerTPS = *time.NewTicker(time.Second)
	tre.tickerTPM = *time.NewTicker(time.Minute)
	// start ticker go routine
	go tre.tickListener()
}

func (tre *transactionRateEngine) tickListener() {
	for {
		select {
		case <-tre.tickerTPS.C:
			tre.Lock()
			for k, v := range tre.mapCounterTPS {
				tre.logSession.LogTransactionRate(k, float64(v), unit_transaction_rate_tps)
			}
			// reset counters
			tre.mapCounterTPS = make(map[string]uint32)
			tre.Unlock()
		case <-tre.tickerTPM.C:
			tre.Lock()
			for k, v := range tre.mapCounterTPM {
				tre.logSession.LogTransactionRate(k, float64(v), unit_transaction_rate_tpm)
			}
			// reset counters
			tre.mapCounterTPM = make(map[string]uint32)
			tre.Unlock()
		}
	}
}

func (ls *LogSession) TPSadd(kpi string) uint32 {
	ls.transactionRateEngine.Lock()
	defer ls.transactionRateEngine.Unlock()
	if val, ok := ls.transactionRateEngine.mapCounterTPS[kpi]; ok {
		val++
		ls.transactionRateEngine.mapCounterTPS[kpi] = val
		return val
	}
	ls.transactionRateEngine.mapCounterTPS[kpi] = 1
	return 1
}

func (ls *LogSession) TPMadd(kpi string) uint32 {
	ls.transactionRateEngine.Lock()
	defer ls.transactionRateEngine.Unlock()
	if val, ok := ls.transactionRateEngine.mapCounterTPM[kpi]; ok {
		val++
		ls.transactionRateEngine.mapCounterTPM[kpi] = val
		return val
	}
	ls.transactionRateEngine.mapCounterTPM[kpi] = 1
	return 1
}

func (ls *LogSession) LogTransactionRate(kpi string, value float64, unit string) {
	var le = logEvent{
		logtype:       TYPE_TRANS_RATE,
		severity:      SEVERITY_INFO,
		transactionID: "",
		tenantID:      "",
		scopeID:       "",
		message:       "",
		fact:          "",
		metricKey:     kpi,
		value:         value,
		unit:          unit,
	}
	ls.log(le)
}
