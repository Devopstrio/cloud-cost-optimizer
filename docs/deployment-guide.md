# Developer & Deployment Guide: Cloud Cost Optimizer

This guide outlines installation, Golang compilation, execution, and testing procedures.

## 1. Installation & Binary Build

```bash
git clone https://github.com/Devopstrio/cloud-cost-optimizer.git
cd cloud-cost-optimizer

# Build Go server binary
go build -o optimizer-server main.go
```

## 2. Execute Cost Optimizer CLI

```bash
./optimizer-server
```

## 3. Run Native Go Test Suite

```bash
go test -v ./...
```
