package optimizer

// ResourceMetrics represents cloud resource utilization data
type ResourceMetrics struct {
	ResourceID   string  `json:"resource_id"`
	CPUUtilPct   float64 `json:"cpu_util_pct"`
	MemoryUtilPct float64 `json:"memory_util_pct"`
	MonthlyCostUSD float64 `json:"monthly_cost_usd"`
}

// OptimizationResult holds rightsizing recommendations
type OptimizationResult struct {
	ResourceID          string  `json:"resource_id"`
	IsIdle              bool    `json:"is_idle"`
	RecommendedAction   string  `json:"recommended_action"`
	EstimatedMonthlySave float64 `json:"estimated_monthly_save"`
}

// OptimizerEngine computes cloud resource rightsizing
type OptimizerEngine struct {
	cpuIdleThreshold float64
}

// NewOptimizerEngine initializes a new OptimizerEngine
func NewOptimizerEngine(cpuIdleThreshold float64) *OptimizerEngine {
	if cpuIdleThreshold <= 0 {
		cpuIdleThreshold = 5.0 // 5% CPU utilization considered idle
	}
	return &OptimizerEngine{cpuIdleThreshold: cpuIdleThreshold}
}

// EvaluateResource evaluates utilization and calculates savings
func (oe *OptimizerEngine) EvaluateResource(res ResourceMetrics) OptimizationResult {
	if res.CPUUtilPct < oe.cpuIdleThreshold {
		return OptimizationResult{
			ResourceID:          res.ResourceID,
			IsIdle:              true,
			RecommendedAction:   "Terminate or Downsize Instance",
			EstimatedMonthlySave: res.MonthlyCostUSD * 0.80,
		}
	}

	return OptimizationResult{
		ResourceID:          res.ResourceID,
		IsIdle:              false,
		RecommendedAction:   "Maintain Current Instance Size",
		EstimatedMonthlySave: 0,
	}
}
