package log

func (ls *LogSession) Log(transactionID string, tenantID string, scopeID string, message string) {
	var le = logEvent{
		logtype:       TYPE_LOG,
		severity:      SEVERITY_NORMAL,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       message,
	}
	ls.log(le)
}

func (ls *LogSession) LogWarning(transactionID string, tenantID string, scopeID string, message string) {
	var le = logEvent{
		logtype:       TYPE_LOG,
		severity:      SEVERITY_WARNING,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       message,
	}
	ls.log(le)
}

func (ls *LogSession) LogError(transactionID string, tenantID string, scopeID string, message string) {
	var le = logEvent{
		logtype:       TYPE_LOG,
		severity:      SEVERITY_ERROR,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       message,
	}
	ls.log(le)
}

func (ls *LogSession) LogCritical(transactionID string, tenantID string, scopeID string, message string) {
	var le = logEvent{
		logtype:       TYPE_LOG,
		severity:      SEVERITY_CRITICAL,
		transactionID: transactionID,
		tenantID:      tenantID,
		scopeID:       scopeID,
		message:       message,
	}
	ls.log(le)
}
