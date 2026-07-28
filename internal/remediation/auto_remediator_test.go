package remediation

import (
	"strings"
	"testing"

	"github.com/Devopstrio/cloud-cost-optimizer/internal/optimizer"
)

func TestRemediateDryRun(t *testing.T) {
	remediator := NewAutoRemediator(true)
	opt := optimizer.OptimizationResult{
		ResourceID:        "i-9912",
		IsIdle:            true,
		RecommendedAction: "Terminate",
	}

	executed, msg := remediator.Remediate(opt)
	if !executed {
		t.Error("expected dry-run remediation to execute")
	}

	if !strings.Contains(msg, "[DRY-RUN]") {
		t.Errorf("expected dry-run message tag, got %s", msg)
	}
}
