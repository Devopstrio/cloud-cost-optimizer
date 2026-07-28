output "alarm_arn" {
  value       = aws_cloudwatch_metric_alarm.idle_resource_alarm.arn
  description = "CloudWatch Alarm ARN"
}
