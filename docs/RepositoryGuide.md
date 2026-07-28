# Repository Structure Guide

This guide explains the directory layout and module responsibilities.

```
cloud-cost-optimizer/
├── .github/              # CI/CD workflows, issue & PR templates
├── architecture/         # Mermaid sequence flow diagrams
├── deployment/           # Kubernetes manifests & Kustomize environment overlays
├── docs/                 # Enterprise architectural, deployment, & operational guides
├── examples/             # Real-world request/response JSON payloads
├── images/               # High-resolution architecture & workflow diagrams
├── internal/             # Go source packages (optimizer, recommendation, remediation)
├── terraform/            # Multi-cloud OpenTofu / Terraform IaC modules
├── tests/                # Unit, integration, and API test suites
├── Dockerfile            # Container build specification
├── docker-compose.yml    # Multi-container local orchestration
├── main.go               # Go CLI server entrypoint
├── go.mod                # Go module manifest
└── README.md             # Accelerator documentation manual
```
