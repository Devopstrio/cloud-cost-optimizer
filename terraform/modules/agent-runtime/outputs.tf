output "cluster_id" {
  value       = aws_ecs_cluster.optimizer_cluster.id
  description = "ECS Cluster ID"
}
