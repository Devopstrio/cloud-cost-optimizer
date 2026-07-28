package unit

import (
	"testing"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

func TestUnitEvaluateResourceIdle(t *testing.T) {
	eng := optimizer.NewOptimizerEngine(5.0)
	res := optimizer.ResourceMetrics{
		ResourceID:     "i-001923-idle",
		CPUUtilPct:     2.1,
		MemoryUtilPct:  10.0,
		MonthlyCostUSD: 1000.0,
	}

	opt := eng.EvaluateResource(res)
	if !opt.IsIdle {
		t.Error("expected resource to be flagged as idle")
	}

	if opt.EstimatedMonthlySave != 800.0 {
		t.Errorf("expected $800 monthly savings, got $%.2f", opt.EstimatedMonthlySave)
	}
}
