# terraform-aws-flowlogs [![Build Status](https://github.com/barundel/terraform-aws-flowlogs/workflows/build/badge.svg)](https://github.com/barundel/terraform-aws-flowlogs/actions)

> **A Terraform module for creating FlowLogs resources.**

## Table of Contents

- [Maintenance](#maintenance)
- [Getting Started](#getting-started)
- [License](#license)

## Maintenance

This project is maintained [Ben](https://github.com/barundel), anyone is welcome to contribute. 


## Getting Started

These are terraform examples covering some of the basic usages of this sub-module.

### Simple VPC Flow-logs to S3

The folowing example creates a VPC flow log to s3. 
````
module "vpc_flowlogs" {
  source = "../../flow-logs"

  vpc_id = "vpc-vpc_id"

  log_destination = "bucket_arn"
  log_destination_type = "s3"

}
````

### VPC Flow-logs to s3 with log-format

````
module "vpc_flowlogs" {
  source = "../../tf-modules/terraform-aws-logging/flow-logs"

  vpc_id = "vpc-vpc_id"

  log_destination = "bucket_arn"
  log_destination_type = "s3"

  log_format = "$${account-id} $${action} $${bytes} $${dstaddr} $${dstport} $${end} $${instance-id} $${interface-id} $${log-status} $${packets} $${pkt-dstaddr} $${pkt-srcaddr} $${protocol} $${srcaddr} $${srcport} $${start} $${subnet-id} $${tcp-flags} $${type} $${version} $${vpc-id}"

}
````

For more information on the available log format fields see the [documentation](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html) 

> Note that in terraform you need the double $$ format. 

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| aws | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| create\_role | Creates a basic IAM role for flowlogs to do its job | `bool` | `true` | no |
| eni\_id | Elastic Network Interface ID to attach flow logs to. | `string` | `""` | no |
| iam\_role\_arn | The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group | `string` | `""` | no |
| log\_destination | The ARN of the logging destination. | `string` | `""` | no |
| log\_destination\_type | The type of the logging destination. Valid values: cloud-watch-logs, s3. Default: cloud-watch-logs. | `string` | `"cloud-watch-logs"` | no |
| log\_format | The fields to include in the flow log record, in the order in which they should appear. | `string` | `""` | no |
| subnet\_id | Subnet ID to attach the flow logs to. | `string` | `""` | no |
| traffic\_type | The type of traffic to capture. Valid values: ACCEPT,REJECT, ALL, Defaults to ALL. | `string` | `"ALL"` | no |
| vpc\_id | VPC ID to attach flow logs to. | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| flow\_log\_id | The Flow Log ID |
| iam\_role\_name | Name of the IAM role |
<!--- END_TF_DOCS --->

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.