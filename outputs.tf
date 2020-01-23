output "flow_log_id" {
  value = concat(aws_flow_log.vpc_flow_log.*.id, [""])[0]
  description = "The Flow Log ID"
}

output "iam_role_name" {
  value = concat(aws_iam_role.flow_logs_role.*.name, [""])[0]
  description = "Name of the IAM role"
}