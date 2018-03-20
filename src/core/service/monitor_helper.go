package service

import (
	"github.com/sanksons/reflorest/src/common/constants"
	workflow "github.com/sanksons/reflorest/src/core/common/orchestrator"
)

func getCustomMetricPrefix(data workflow.WorkFlowData) string {
	var monitorMetricPrefix string
	monitorCustomMetricPrefix, mcmpError := data.ExecContext.Get(constants.MonitorCustomMetric)
	if mcmpError == nil {
		monitorMetricPrefix, _ = monitorCustomMetricPrefix.(string)
	}
	return monitorMetricPrefix
}
