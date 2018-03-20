package service

import (
	"github.com/sanksons/reflorest/src/common/ratelimiter"
	"github.com/sanksons/reflorest/src/core/common/orchestrator"
	"github.com/sanksons/reflorest/src/core/common/utils/healthcheck"
	"github.com/sanksons/reflorest/src/core/common/versionmanager"
)

type APIInterface interface {
	GetVersion() versionmanager.Version

	GetOrchestrator() orchestrator.Orchestrator

	GetHealthCheck() healthcheck.HCInterface

	GetRateLimiter() ratelimiter.RateLimiter

	Init()
}
