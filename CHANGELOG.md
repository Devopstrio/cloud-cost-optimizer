# Changelog

All notable changes to the **Cloud Cost Optimizer** platform are documented in this file.

## [1.0.0] - 2026-07-28

### Added
- Enterprise Golang Rightsizing Engine (`internal/optimizer/engine.go`).
- FinOps Savings Recommendation Advisor (`internal/recommendation/advisor.go`).
- Automated Dry-Run & Live Cost Remediator (`internal/remediation/auto_remediator.go`).
- Modular OpenTofu / Terraform IaC blueprints (`terraform/modules/`).
- Kubernetes Kustomize base and environment overlays (`deployment/kubernetes/`).
- Native Go 1.22 test suite (`tests/` and `*_test.go`).
- Complete enterprise documentation suite (`docs/`).
