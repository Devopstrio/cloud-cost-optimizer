# Golang Engine Implementation Guide

This guide details how to integrate and extend the `cloud-cost-optimizer` Go engine.

## Evaluating Workload Metrics

```go
package main

import (
	"fmt"
	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

func main() {
	eng := optimizer.NewOptimizerEngine(5.0) // 5% CPU threshold
	res := optimizer.ResourceMetrics{
		ResourceID:     "i-0992-idle",
		CPUUtilPct:     1.8,
		MemoryUtilPct:  8.0,
		MonthlyCostUSD: 1500.0,
	}

	opt := eng.EvaluateResource(res)
	if opt.IsIdle {
		fmt.Printf("Idle workload found! Potential monthly savings: $%.2f\n", opt.EstimatedMonthlySave)
	}
}
```

## Executing Auto-Remediation

```go
remediator := remediation.NewAutoRemediator(true) // dry-run mode
executed, msg := remediator.Remediate(opt)
fmt.Println(msg)
```
