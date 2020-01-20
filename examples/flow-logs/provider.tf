provider "aws" {
    region = "eu-west-1"

    assume_role {
        role_arn = "arn:aws:iam::116147290797:role/cdl/admin/CloudTeam"
    }
}