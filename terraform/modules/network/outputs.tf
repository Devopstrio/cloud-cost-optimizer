output "vpc_id" {
  value       = aws_vpc.optimizer_vpc.id
  description = "Provisioned VPC ID"
}
