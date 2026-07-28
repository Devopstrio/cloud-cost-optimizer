output "log_group_arn" {
  value       = aws_cloudwatch_log_group.optimizer_log_group.arn
  description = "Log Group ARN"
}
