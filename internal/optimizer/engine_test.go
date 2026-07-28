package optimizer

import (
	"testing"
)

func TestEvaluateResourceIdle(t *testing.T) {
	eng := NewOptimizerEngine(5.0)
	res := ResourceMetrics{
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

func TestEvaluateResourceActive(t *testing.T) {
	eng := NewOptimizerEngine(5.0)
	res := ResourceMetrics{
		ResourceID:     "i-001924-active",
		CPUUtilPct:     45.0,
		MemoryUtilPct:  60.0,
		MonthlyCostUSD: 1000.0,
	}

	opt := eng.EvaluateResource(res)
	if opt.IsIdle {
		t.Error("expected active resource to not be flagged as idle")
	}
}
