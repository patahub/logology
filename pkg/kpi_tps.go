package log

import (
	"sync"
	"time"
)

type TransactionRate struct {
	sync.Mutex
	mapCounterTPS map[uint64]uint32
	mapCounterTPM map[uint64]uint32
	tickerTPS     time.Ticker
}

func (tr *TransactionRate) init() {
	// .... TODO init maps
	tr.tickerTPS = time.NewTicker(tr.ticker(1 * time.Second)
	
}

func (tr *TransactionRate) incTPS() {
	TransactionRate.mTPS.Lock()
	tr.mapCounter
}
