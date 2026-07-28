output "vpc_id" {
  value       = module.network.vpc_id
  description = "VPC ID"
}

output "ecs_cluster_id" {
  value       = module.agent_runtime.cluster_id
  description = "ECS Cluster ID"
}

output "log_group_arn" {
  value       = module.logging.log_group_arn
  description = "CloudWatch Log Group ARN"
}
