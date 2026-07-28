# Multi-Cloud Deployment Guide

This document covers production deployment across Kubernetes clusters, AWS, and Azure.

## 1. OpenTofu / Terraform Infrastructure Provisioning

```bash
cd terraform
tofu init
tofu plan
tofu apply -auto-approve
```

## 2. Kubernetes Deployment via Kustomize

```bash
# Apply base namespace and configuration
kubectl apply -f deployment/kubernetes/base/namespace.yaml
kubectl apply -f deployment/kubernetes/base/configmap.yaml

# Deploy production overlay
kubectl apply -k deployment/kubernetes/overlays/prod/
```
