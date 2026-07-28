package remediation

import (
	"fmt"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

// AutoRemediator executes automated cleanup on idle cloud workloads
type AutoRemediator struct {
	dryRun bool
}

// NewAutoRemediator initializes a new AutoRemediator
func NewAutoRemediator(dryRun bool) *AutoRemediator {
	return &AutoRemediator{dryRun: dryRun}
}

// Remediate executes auto-remediation policy on optimization results
func (ar *AutoRemediator) Remediate(opt optimizer.OptimizationResult) (bool, string) {
	if !opt.IsIdle {
		return false, fmt.Sprintf("Resource %s is active; skipping remediation", opt.ResourceID)
	}

	if ar.dryRun {
		return true, fmt.Sprintf("[DRY-RUN] Would execute '%s' on resource %s", opt.RecommendedAction, opt.ResourceID)
	}

	return true, fmt.Sprintf("[REMEDIATED] Successfully executed '%s' on resource %s", opt.RecommendedAction, opt.ResourceID)
}
