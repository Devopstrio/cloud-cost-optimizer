package main

import (
	"fmt"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/recommendation"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/remediation"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("  Devopstrio Cloud Cost Optimizer Engine (v1.0.0)")
	fmt.Println("==================================================")

	eng := optimizer.NewOptimizerEngine(5.0)
	adv := recommendation.NewAdvisor()
	remediator := remediation.NewAutoRemediator(true)

	sampleMetrics := []optimizer.ResourceMetrics{
		{ResourceID: "i-aws-idle-01", CPUUtilPct: 2.0, MemoryUtilPct: 8.0, MonthlyCostUSD: 1200.0},
		{ResourceID: "i-aws-active-02", CPUUtilPct: 65.0, MemoryUtilPct: 70.0, MonthlyCostUSD: 800.0},
	}

	results := make([]optimizer.OptimizationResult, 0)
	for _, res := range sampleMetrics {
		opt := eng.EvaluateResource(res)
		results = append(results, opt)
		fmt.Printf("[1/3] Evaluated %s -> Idle: %t, Action: %s\n", opt.ResourceID, opt.IsIdle, opt.RecommendedAction)
	}

	rep := adv.AggregateSavings(results)
	fmt.Printf("[2/3] %s\n", adv.FormatSummary(rep))

	for _, opt := range results {
		if opt.IsIdle {
			_, msg := remediator.Remediate(opt)
			fmt.Printf("[3/3] Remediation: %s\n", msg)
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("SUCCESS: Cloud Cost Optimization Evaluation Complete!")
}
