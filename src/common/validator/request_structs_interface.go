package service

import (
	workflow "github.com/sanksons/reflorest/src/core/common/orchestrator"
)

type RequestStructs struct {
	Headers interface{}
	Params  interface{}
	Body    interface{}
}

type ValidateRequest interface {
	GetStructs(data workflow.WorkFlowData) RequestStructs
}
