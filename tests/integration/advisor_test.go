package integration

import (
	"testing"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/recommendation"
)

func TestIntegrationAdvisorSavings(t *testing.T) {
	adv := recommendation.NewAdvisor()
	results := []optimizer.OptimizationResult{
		{ResourceID: "res-1", IsIdle: true, EstimatedMonthlySave: 500.0},
		{ResourceID: "res-2", IsIdle: true, EstimatedMonthlySave: 300.0},
	}

	rep := adv.AggregateSavings(results)
	if rep.TotalIdleResources != 2 {
		t.Errorf("expected 2 idle resources, got %d", rep.TotalIdleResources)
	}

	if rep.TotalMonthlySaveUSD != 800.0 {
		t.Errorf("expected $800 total savings, got $%.2f", rep.TotalMonthlySaveUSD)
	}
}
