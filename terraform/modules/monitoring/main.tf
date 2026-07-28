resource "aws_cloudwatch_metric_alarm" "idle_resource_alarm" {
  alarm_name          = "cloud-cost-idle-workload-high"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "IdleResourceCount"
  namespace           = "Devopstrio/FinOps"
  period              = 300
  statistic           = "Maximum"
  threshold           = 5
  alarm_description   = "Alarm when idle workload count exceeds 5"
}
