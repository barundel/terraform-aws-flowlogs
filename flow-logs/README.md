# Usage
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

No output.
<!--- END_TF_DOCS --->
