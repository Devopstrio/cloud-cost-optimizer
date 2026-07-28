variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "VPC CIDR Block"
}

variable "environment" {
  type        = string
  default     = "production"
  description = "Deployment Environment"
}
