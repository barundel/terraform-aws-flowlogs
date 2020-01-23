package flowlongs
import (
	"fmt"
	session2 "github.com/JoshiiSinfield/go-helpers/aws/service/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Terratest docs - https://godoc.org/github.com/gruntwork-io/terratest

var awsRegion = "eu-west-1"
var awsProfile = "a"
var vpcId = "vpc-a"
var log_destination = "a"
var log_destination_type = "s3"

func ValidateFlowLogAttributes(t *testing.T, flowLogId string) {
	// Creating a session using profile and region var. (giving you permissions)
	mySess := session2.CreateSessionFromProfile(awsProfile, awsRegion)
	// Passing the permissions into an ec2 thing
	ec2Svc := ec2.New(mySess)
	// Calling ec2 describe flow logs input and passing in filter on what to search for?
	params := &ec2.DescribeFlowLogsInput{
		FlowLogIds: aws.StringSlice([]string{flowLogId}),
	}
	// Doing something with the results?
	result, err := ec2Svc.DescribeFlowLogs(params)
	// If there is an error print error?
	// Not sure if this is just after the results or at any point during the code execution?
	if err != nil {
		fmt.Println(err)
	}
	// creating a var with the results of flowlogs
	ben := result.FlowLogs
	// creating a loop through of the results var
	for _, v := range ben {
		//v.FlowLogStatus
		// print the output of the results that equal to the flowlog status
		//fmt.Printf("FlowLog Status: %s", *v.FlowLogStatus)
		assert.Equal(t, "ACTIVE", *v.FlowLogStatus)
	}
}

func SetupTerraformOptions() *terraform.Options {
	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vpc_id":               vpcId,
			"log_destination_type": log_destination_type,
			"log_destination":      log_destination,
		},
		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
			"AWS_PROFILE":        awsProfile,
		},
	}
	return terraformOptions
}

func TestTerraformVPCFlowLogs(t *testing.T) {
	t.Parallel()
	terraformOptions := SetupTerraformOptions()
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
	// Run `terraform output` to get the value of an output variable
	roleName := terraform.Output(t, terraformOptions, "iam_role_name")
	theFlowLogId := terraform.Output(t, terraformOptions, "flow_log_id")
	assert.Equal(t, fmt.Sprintf("%s-FlowLogs-Role", vpcId), roleName)
	ValidateFlowLogAttributes(t, theFlowLogId)
	// At the end of the test, run `terraform d estroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)
}