# Getting Started Guide: Cloud Cost Optimizer

Welcome to **Cloud Cost Optimizer**. This guide provides step-by-step instructions for installation, Golang compilation, and running test suites.

## Prerequisites

- Golang `>= 1.22.0`
- OpenTofu `>= 1.6.0` or Terraform `>= 1.6.0`
- Kubernetes `kubectl` CLI

## Local Installation & Binary Build

```bash
# Clone the repository
git clone https://github.com/Devopstrio/cloud-cost-optimizer.git
cd cloud-cost-optimizer

# Build Go server binary
go build -o optimizer-server main.go

# Run server binary
./optimizer-server

# Run native Go test suite
go test -v ./...
```
