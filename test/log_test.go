package main

import (
	"fmt"
	"testing"
	"time"

	log "github.com/patahub/logology/pkg"
)

func TestLog(t *testing.T) {
	var ls = log.NewLogSession("myservice", log.SEVERITY_INFO, log.TYPE_ALL, log.ANONYMOUS_TRUE, log.CONSOLE_ON)
	ls.Log("transaction1", "tenant1", "scope1", "something 1")
	ls.Log("transaction1", "tenant1", "scope1", "something 2")
	ls.Log("transaction1", "tenant1", "scope1", "something 3")
	ls.Log("transaction1", "tenant1", "scope1", "something varied")
	ls.LogWarning("transaction1", "tenant1", "scope1", "something varied")
	ls.LogError("transaction1", "tenant1", "scope1", "something varied")
	ls.LogCritical("transaction1", "tenant1", "scope1", "something varied")
	// delay for channel consuming events
	time.Sleep(1 * time.Second)
}

func TestMetricDuration(t *testing.T) {
	var ls = log.NewLogSession("myservice", log.SEVERITY_INFO, log.TYPE_ALL, log.ANONYMOUS_TRUE, log.CONSOLE_ON)
	ls.SetTimeStampFormat("2006-01-02 15:04:05")
	ls.LogMetricDurationStart("transaction1", "tenant1", "scope1")
	ls.LogMetricDurationStart("transaction1", "tenant1", "scope2")
	ls.LogMetricDurationStart("transaction2", "tenant2", "scope1")
	ls.LogMetricDurationEnd("transaction1", "tenant1", "scope1", "time measurement end A", "loop duration")
	ls.LogMetricDurationEnd("transaction2", "tenant2", "scope1", "time measurement end C", "blink")
	ls.LogMetricDurationEnd("transaction1", "tenant1", "scope2", "time measurement end B", "time-to-dothis")
	// delay for channel consuming events
	time.Sleep(1 * time.Second)
}

func TestKPI(t *testing.T) {
	var ls = log.NewLogSession("myservice", log.SEVERITY_INFO, log.TYPE_ALL, log.ANONYMOUS_TRUE, log.CONSOLE_ON)
	ls.SetTimeStampFormat("2006-01-02 15:04:05")
	// delay for channel consuming events
	time.Sleep(2 * time.Second)
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT PERA", ls.TPMadd("PERA"))
	fmt.Println("COUNT PERA", ls.TPMadd("PERA"))
	time.Sleep(5 * time.Second)
	fmt.Println("COUNT PERA", ls.TPMadd("PERA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT BANANA", ls.TPSadd("BANANA"))
	fmt.Println("COUNT PERA", ls.TPMadd("PERA"))
	time.Sleep(61 * time.Second)
}
