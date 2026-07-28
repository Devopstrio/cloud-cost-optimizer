resource "aws_cloudwatch_log_group" "optimizer_log_group" {
  name              = "/aws/finops/cloud-cost-optimizer"
  retention_in_days = 30
}
