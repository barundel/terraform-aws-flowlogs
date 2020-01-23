variable "iam_role_arn" {
  default = ""
  description = "The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group"
}

variable "log_format" {
  default = ""
  description = "The fields to include in the flow log record, in the order in which they should appear."
}

variable "log_destination" {
  default = ""
  description = "The ARN of the logging destination."
}

variable "log_destination_type" {
  default = "cloud-watch-logs"
  description = "The type of the logging destination. Valid values: cloud-watch-logs, s3. Default: cloud-watch-logs."
}

variable "eni_id" {
  default = ""
  description = "Elastic Network Interface ID to attach flow logs to."
}

variable "subnet_id" {
  default = ""
  description = "Subnet ID to attach the flow logs to."
}

variable "vpc_id" {
  description = "VPC ID to attach flow logs to."
  default = ""
}

variable "traffic_type" {
  default = "ALL"
  description = "The type of traffic to capture. Valid values: ACCEPT,REJECT, ALL, Defaults to ALL."
}

variable "create_role" {
  default = true
  description = "Creates a basic IAM role for flowlogs to do its job"
}