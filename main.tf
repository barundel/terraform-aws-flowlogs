########## ##########
# Locals
########## ##########
locals {
  //iam_role_arn = coalesce([aws_iam_role.flow_logs_role.*.arn[0]], var.iam_role_arn)
  iam_role_arn = coalescelist(aws_iam_role.flow_logs_role.*.arn, [var.iam_role_arn])
  name = random_pet.name.id
}

resource "random_pet" "name" {
  length = 1
}

########## ##########
# VPC FlowLogs
########## ##########
resource "aws_flow_log" "vpc_flow_log" {
  count = length(var.vpc_id) > 0 ? 1 : 0

  vpc_id = var.vpc_id

  iam_role_arn = local.iam_role_arn[0]

  log_destination = var.log_destination
  log_destination_type = var.log_destination_type
  log_format = var.log_format

  traffic_type = upper(var.traffic_type)

}

########## ##########
# Subnet FlowLogs
########## ##########
resource "aws_flow_log" "subnet_flow_log" {
  count = length(var.subnet_id) > 0 ? length(var.subnet_id) : 0

  subnet_id = var.subnet_id[count.index]

  iam_role_arn = local.iam_role_arn[0]

  log_destination = var.log_destination
  log_destination_type = var.log_destination_type
  log_format = var.log_format

  traffic_type = upper(var.traffic_type)

}

########## ##########
# ENI FlowLogs
########## ##########
resource "aws_flow_log" "eni_flow_log" {
  count = length(var.eni_id) > 0 ? length(var.eni_id) : 0

  eni_id = element(var.eni_id, count.index)

  iam_role_arn = local.iam_role_arn[0]

  log_destination = var.log_destination
  log_destination_type = var.log_destination_type
  log_format = var.log_format

  traffic_type = upper(var.traffic_type)

}

########## ##########
# Basic IAM Role
########## ##########
resource "aws_iam_role" "flow_logs_role" {
  count = var.create_role ? 1 : 0

  name = "${local.name}-FlowLogs-Role"

  assume_role_policy = data.aws_iam_policy_document.flow_log_trust.json
}

resource "aws_iam_role_policy" "flow_logs" {
  count = var.create_role ? 1 : 0

  name = "${local.name}-FlowLogs-Policy"
  role =  aws_iam_role.flow_logs_role.*.id[0]
  policy = data.aws_iam_policy_document.flow_log_permissions.json
}






