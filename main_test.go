package main

import (
	"testing"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/recommendation"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/remediation"
)

func TestFullOptimizationPipeline(t *testing.T) {
	eng := optimizer.NewOptimizerEngine(5.0)
	adv := recommendation.NewAdvisor()
	remediator := remediation.NewAutoRemediator(true)

	res := optimizer.ResourceMetrics{
		ResourceID:     "res-01",
		CPUUtilPct:     1.5,
		MemoryUtilPct:  5.0,
		MonthlyCostUSD: 1000.0,
	}

	opt := eng.EvaluateResource(res)
	if !opt.IsIdle {
		t.Error("expected resource to be idle")
	}

	rep := adv.AggregateSavings([]optimizer.OptimizationResult{opt})
	if rep.TotalIdleResources != 1 || rep.TotalMonthlySaveUSD != 800.0 {
		t.Errorf("savings aggregation failed: %+v", rep)
	}

	executed, _ := remediator.Remediate(opt)
	if !executed {
		t.Error("expected remediation execution to succeed")
	}
}
