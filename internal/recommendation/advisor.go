package recommendation

import (
	"fmt"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

// SavingsReport aggregates cost optimization savings across cloud workloads
type SavingsReport struct {
	TotalIdleResources int     `json:"total_idle_resources"`
	TotalMonthlySaveUSD float64 `json:"total_monthly_save_usd"`
}

// Advisor generates FinOps optimization recommendations
type Advisor struct{}

// NewAdvisor initializes a new Advisor
func NewAdvisor() *Advisor {
	return &Advisor{}
}

// AggregateSavings compiles total savings from optimization results
func (a *Advisor) AggregateSavings(results []optimizer.OptimizationResult) SavingsReport {
	var totalSave float64
	var idleCount int

	for _, r := range results {
		if r.IsIdle {
			idleCount++
			totalSave += r.EstimatedMonthlySave
		}
	}

	return SavingsReport{
		TotalIdleResources:  idleCount,
		TotalMonthlySaveUSD: totalSave,
	}
}

// FormatSummary returns a human-readable recommendation summary
func (a *Advisor) FormatSummary(report SavingsReport) string {
	return fmt.Sprintf("FinOps Savings Advisor: Identified %d idle resource(s) with potential savings of $%.2f/month",
		report.TotalIdleResources, report.TotalMonthlySaveUSD)
}
