package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

func main() {
	//mySess := session.Must(session.NewSessionWithOptions(session.Options{
	//    Config:  aws.Config{Region: aws.String("eu-west-1"),},
	//    Profile: "profile cdl-kingfisher-non-prod",
	//}))

	//Use this if not wanting to specifty a profile ...
	mySess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})
	if err != nil {
		fmt.Printf("Error getting a session: %o", err)
	}

	s3Svc := s3.New(mySess)

	s3BucketName := "flowlogs-data-pricing-686794321847"

	selectInput := s3.SelectObjectContentInput{
		Bucket:         aws.String(s3BucketName),
		Key:            aws.String("AWSLogs/686794321847/vpcflowlogs/eu-west-1/2019/12/18/686794321847_vpcflowlogs_eu-west-1_fl-0f7c58d2ddbcf20d7_20191218T0000Z_3b25f0ed.log.gz"),
		Expression:     aws.String("select * from s3object s limit 50"),
		ExpressionType: aws.String("SQL"),
		InputSerialization: &s3.InputSerialization{
			CompressionType: aws.String("GZIP"),
			CSV: &s3.CSVInput{
				FileHeaderInfo: aws.String("Use"),
				FieldDelimiter: aws.String(" "),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{},
		},
	}
	results, err := s3Svc.SelectObjectContent(&selectInput)

	if err != nil {
		fmt.Printf("Error when sleecitng from object: %s", err.Error())
	}
	defer results.EventStream.Close()

	var theResults []FlowLogRecord

	for event := range results.EventStream.Events() {
		switch v := event.(type) {
		case *s3.RecordsEvent:
			// s3.RecordsEvent.Records is a byte slice of select records
			//fmt.Println(string(v.Payload))

			splitStrings := strings.Split(string(v.Payload), "\n")

			for _, splitString := range splitStrings {
				//fmt.Printf("index: %d, value: %s\n", i, splitString)
				if splitString == "" {
					continue
				}
				var thisFlowLogRecord FlowLogRecord
				err := json.Unmarshal([]byte(splitString), &thisFlowLogRecord)
				if err != nil {
					fmt.Printf("Error unmarshalling Flow Log Record: %s\n", err.Error())
				}
				theResults = append(theResults, thisFlowLogRecord)
			}

		case *s3.StatsEvent:
			// s3.StatsEvent contains information on the data that’s processed
			//fmt.Println("Processed", *tv.Details.BytesProcessed, "bytes")
			fmt.Println("stats")
		case *s3.EndEvent:
			// s3.EndEvent
			fmt.Println("SelectObjectContent completed")
		}
	}

	for _, theFlowLogRecord := range theResults {
		fmt.Printf("\nFlow Log Record: %s", theFlowLogRecord)
		fmt.Printf("\nRecord Source Addr: %s", theFlowLogRecord.SrcAddr)
	}

	//results.EventStream.Reader

}

type FlowLogRecord struct {
	Version     string `json:"version"`
	AccountId   string `json:"account-id"`
	InterfaceId string `json:"interface-id"`
	SrcAddr     string `json:"srcaddr"`
	DstAddr     string `json:"dstaddr"`
	SrcPort     string `json:"srcport"`
	DstPort     string `json:"dstport"`
	Protocol    string `json:"protocol"`
	Packets     string `json:"packets"`
	Bytes       string `json:"bytes"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Action      string `json:"action"`
	LogStatus   string `json:"log-status"`
}
