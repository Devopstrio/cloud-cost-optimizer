terraform {
  required_version = ">= 1.6.0"
}

resource "aws_vpc" "optimizer_vpc" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "cloud-cost-optimizer-vpc"
    Environment = var.environment
    ManagedBy   = "Devopstrio"
  }
}
