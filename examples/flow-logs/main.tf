module "vpc_flowlogs" {
  source = "../../flow-logs"

  vpc_id = "vpc-0f7ad8a4a5b2fc575"

  log_destination = "arn:aws:s3:::cloud-team-test-bucket"
  log_destination_type = "s3"

}