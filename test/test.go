package main

import (
	"fmt"

	log "github.com/patahub/logology/pkg"
)

func main() {
	fmt.Printf("here1\n")
	var ls = log.NewLogSession("myservice", log.SEVERITY_INFO, log.TYPE_ALL, log.ANONYMOUS_FALSE)
	ls.Log("transaction1", "tenant1", "scope1", "something 1")
	ls.Log("transaction1", "tenant1", "scope1", "something 2")
	ls.Log("transaction1", "tenant1", "scope1", "something 3")
	ls.Log("transaction1", "tenant1", "scope1", "something varied")
	ls.LogWarning("transaction1", "tenant1", "scope1", "something varied")
	ls.LogError("transaction1", "tenant1", "scope1", "something varied")
	ls.LogCritical("transaction1", "tenant1", "scope1", "something varied")
	ls.LogMetricDuration("transaction1", "tenant1", "scope1", "time measurement start A")
	ls.LogMetricDuration("transaction1", "tenant1", "scope2", "time measurement start B")
	ls.LogMetricDuration("transaction2", "tenant1", "scope1", "time measurement start C")

	ls.LogMetricDuration("transaction1", "tenant1", "scope1", "time measurement end A")
	ls.LogMetricDuration("transaction2", "tenant1", "scope1", "time measurement end C")
	ls.LogMetricDuration("transaction1", "tenant1", "scope2", "time measurement end B")
	for {
	}
}
