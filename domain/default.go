package domain

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func CreateDomain(name, ip_address string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS"), os.Getenv("AWS_SECRET"), ""),
	})
	if err != nil {
		return err
	}
	recordSet := &route53.ResourceRecordSet{
		Name:            aws.String(name),
		ResourceRecords: []*route53.ResourceRecord{&route53.ResourceRecord{Value: aws.String(os.Getenv("AWS_IP"))}},
		Type:            aws.String("A"),
		TTL:             aws.Int64(300),
	}
	svc := route53.New(sess)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action:            aws.String("CREATE"),
					ResourceRecordSet: recordSet,
				},
			},
			Comment: aws.String(ip_address),
		},
		HostedZoneId: aws.String(os.Getenv("AWS_HOST")),
	}
	_, err = svc.ChangeResourceRecordSets(input)
	if err != nil {
		return err
	}
	return nil
}
