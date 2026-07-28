package recommendation

import (
	"strings"
	"testing"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

func TestAggregateSavings(t *testing.T) {
	adv := NewAdvisor()
	results := []optimizer.OptimizationResult{
		{ResourceID: "res-1", IsIdle: true, EstimatedMonthlySave: 500.0},
		{ResourceID: "res-2", IsIdle: true, EstimatedMonthlySave: 300.0},
		{ResourceID: "res-3", IsIdle: false, EstimatedMonthlySave: 0.0},
	}

	rep := adv.AggregateSavings(results)
	if rep.TotalIdleResources != 2 {
		t.Errorf("expected 2 idle resources, got %d", rep.TotalIdleResources)
	}

	if rep.TotalMonthlySaveUSD != 800.0 {
		t.Errorf("expected $800 total savings, got $%.2f", rep.TotalMonthlySaveUSD)
	}

	summary := adv.FormatSummary(rep)
	if !strings.Contains(summary, "Identified 2 idle resource(s)") {
		t.Errorf("unexpected summary text: %s", summary)
	}
}
